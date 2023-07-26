package service

import (
	"my-app/backend/internal/interfaces"
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/database/entity"
)

type UserPasswordService struct {
	interfaces.IUserPasswordService
}

func NewUserPasswordService(db *database.Database) interfaces.IUserPasswordService {
	return &UserPasswordService{
		IUserPasswordService: database.NewCrudService[*entity.UserPassword](db),
	}
}
