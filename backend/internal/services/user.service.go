package services

import "my-app/backend/internal/interfaces"

type UserService struct {
}

func NewUserService() interfaces.IUserService {
	return &UserService{}
}
