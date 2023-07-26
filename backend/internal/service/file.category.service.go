package service

import (
	"my-app/backend/internal/interfaces"
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/database/entity"
)

type FileCategoryService struct {
	interfaces.IFileCategoryService
}

func NewFileCategoryService(db *database.Database) interfaces.IFileCategoryService {
	return &FileCategoryService{
		IFileCategoryService: database.NewCrudService[*entity.FileCategory](db),
	}
}
