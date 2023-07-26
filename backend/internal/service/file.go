package service

import (
	"my-app/backend/internal/interfaces"
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/database/entity"
)

type FileService struct {
	interfaces.IFileService
}

func NewFileService(db *database.Database) interfaces.IFileService {
	return &FileService{
		IFileService: database.NewCrudService[*entity.File](db),
	}
}
