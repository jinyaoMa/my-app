package service

import (
	"my-app/backend/internal/interfaces"
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/database/entity"
)

type FileExtensionService struct {
	interfaces.IFileExtensionService
}

func NewFileExtensionService(db *database.Database) interfaces.IFileExtensionService {
	return &FileExtensionService{
		IFileExtensionService: database.NewCrudService[*entity.FileExtension](db),
	}
}
