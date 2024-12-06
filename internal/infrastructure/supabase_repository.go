// インフラ層にSupabaseへの具体的なアクセスロジックを定義する
// ex1. ページング処理
// ex2. クエリビルダー
// ex3. ログ記録
// ex4. DB操作のヘルパー(汎用的なinsertメソッド)
// ex5. エラーハンドリングの共通化
// ex6. メタデータ操作

package infrastructure

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
)

type SupabaseRepository struct {
	Conn *pgx.Conn
}

func NewSupabaseRepository(conn *pgx.Conn) *SupabaseRepository {
	return &SupabaseRepository{Conn: conn}
}

// transaction logic
func (r *SupabaseRepository) WithTransaction(ctx context.Context, fn func(tx pgx.Tx) error) error {
	tx, err := r.Conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback(ctx)
			panic(p)
		} else if err != nil {
			_ = tx.Rollback(ctx)
		} else {
			err = tx.Commit(ctx)
		}
	}()

	err = fn(tx)
	return err
}

// close db connection
func (r *SupabaseRepository) Close(ctx context.Context) error {
	return r.Conn.Close(ctx)
}

// db helth check
func (r *SupabaseRepository) HealthCheck(ctx context.Context) error {
	return r.Conn.Ping(ctx)
}

// ex4
func (s *SupabaseRepository) Insert(ctx context.Context, table string, data map[string]interface{}) error {
	columns := []string{}
	values := []interface{}{}
	placeholders := []string{}

	for col, val := range data {
		columns = append(columns, col)
		values = append(values, val)
		placeholders = append(placeholders, fmt.Sprintf("$%d", len(values)))
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, strings.Join(columns, ", "), strings.Join(placeholders, ", "))
	_, err := s.Conn.Exec(ctx, query, values...)
	return err
}
