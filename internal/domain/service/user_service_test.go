package service

import "github.com/motty93/clean-architecture/internal/domain/model"

type MockUserService struct{}

func (u *MockUserService) CanSendNotification(user *model.User) bool {
	return user.SlackID != ""
}
