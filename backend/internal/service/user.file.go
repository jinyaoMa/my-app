package service

import (
	"my-app/backend/internal/interfaces"
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/database/entity"
)

type UserFileService struct {
	interfaces.IUserFileService
}

func NewUserFileService(db *database.Database) interfaces.IUserFileService {
	return &UserFileService{
		IUserFileService: database.NewCrudService[*entity.UserFile](db),
	}
}
