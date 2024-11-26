// アプリケーション操作を定義
package usecase

import (
	"context"
	"errors"

	"github.com/motty93/clean-architecture/internal/domain/model"
	"github.com/motty93/clean-architecture/internal/domain/service"
	"github.com/motty93/clean-architecture/internal/repository"
)

type UserUsecase struct {
	// dependency injection
	// supbase_repository.goで定義したインターフェースを使える
	repo    repository.UserRepository
	service *service.UserService
}

func NewUseUsecase(repo repository.UserRepository, service *service.UserService) *UserUsecase {
	return &UserUsecase{repo: repo, service: service}
}

func (u *UserUsecase) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	return u.repo.GetUserByID(ctx, id)
}

func (u *UserUsecase) CreateUser(ctx context.Context, user *model.User) error {
	// ビジネスルールを適用
	if err := u.service.ValidateUser(user); err != nil {
		return errors.New("名前、またはメールアドレスが不正です")
	}

	// すでにユーザーが存在するかcheck
	existingUser, err := u.repo.GetUserByID(ctx, user.ID)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("user already exists")
	}

	return u.repo.CreateUser(ctx, user)
}
