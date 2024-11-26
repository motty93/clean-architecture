// インフラ層にSupabaseへの具体的なアクセスロジックを定義する
package infrastructure

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/motty93/clean-architecture/internal/domain/model"
)

type SupabaseRepository struct {
	Conn *pgx.Conn
}

func NewSupabaseRepository(conn *pgx.Conn) *SupabaseRepository {
	return &SupabaseRepository{Conn: conn}
}

func (s *SupabaseRepository) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	query := `SELECT * FROM users`
	rows, err := s.Conn.Query(ctx, query)
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

func (s *SupabaseRepository) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	query := `SELECT * FROM users WHERE id = $1`
	var user model.User
	err := s.Conn.QueryRow(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *SupabaseRepository) CreateUser(ctx context.Context, user *model.User) error {
	query := `INSERT INTO users (name, email) VALUES ($1, $2)`
	_, err := s.Conn.Exec(ctx, query, user.Name, user.Email)
	if err != nil {
		return err
	}

	return nil
}
