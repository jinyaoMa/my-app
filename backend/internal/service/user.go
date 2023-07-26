package service

import (
	"my-app/backend/internal/interfaces"
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/database/entity"
)

type UserService struct {
	interfaces.IUserService
}

func NewUserService(db *database.Database) interfaces.IUserService {
	return &UserService{
		IUserService: database.NewCrudService[*entity.User](db),
	}
}
