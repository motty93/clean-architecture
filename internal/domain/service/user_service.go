// ドメインに特化したビジネスロジックを実装する
package service

import (
	"errors"
	"regexp"

	"github.com/motty93/clean-architecture/internal/domain/model"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) CanSendNotification(user *model.User) bool {
	return user.SlackID != ""
}

func (u *UserService) ValidateUser(user *model.User) error {
	if user.Name == "" {
		return errors.New("name is required")
	}

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if !regexp.MustCompile(emailRegex).MatchString(user.Email) {
		return errors.New("invalid email format")
	}

	return nil
}
