// データアクセスのインターフェースを定義
package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/motty93/clean-architecture/internal/domain/model"
	"github.com/motty93/clean-architecture/internal/infrastructure"
)

type UserRepository interface {
	GetAllUsers(ctx context.Context) ([]*model.User, error)
	GetUserByID(ctx context.Context, id int) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
	CreateUserWithLog(ctx context.Context, user *model.User, logMessage string) error
}

type UserRepositoryImpl struct {
	// supabase_repositoryで定義したConnを利用するためコメントアウト
	// Conn *pgx.Conn
	db *infrastructure.SupabaseRepository
}

func NewUserRepository(db *infrastructure.SupabaseRepository) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	query := `SELECT * FROM users`
	rows, err := r.db.Conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

// ユーザーをIDで取得
func (r *UserRepositoryImpl) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	query := `SELECT * FROM users WHERE id = $1`
	var user model.User
	err := r.db.Conn.QueryRow(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil // ユーザーが存在しない場合
		}
		return nil, fmt.Errorf("failed to fetch user by ID: %w", err)
	}
	return &user, nil
}

// ユーザーを作成
func (r *UserRepositoryImpl) CreateUser(ctx context.Context, user *model.User) error {
	// transactionでuser作成
	return r.db.WithTransaction(ctx, func(tx pgx.Tx) error {
		query := `INSERT INTO users (name, email) VALUES ($1, $2)`
		_, err := tx.Exec(ctx, query, user.Name, user.Email)
		if err != nil {
			return fmt.Errorf("failed to create user: %w", err)
		}

		return nil
	})
}

func (r *UserRepositoryImpl) CreateUserWithLog(ctx context.Context, user *model.User, logMessage string) error {
	return nil
}
