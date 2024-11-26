// データアクセスのインターフェースを定義
package repository

import (
	"context"

	"github.com/motty93/clean-architecture/internal/domain/model"
)

type UserRepository interface {
	GetAllUsers(ctx context.Context) ([]*model.User, error)
	GetUserByID(ctx context.Context, id int) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
}
