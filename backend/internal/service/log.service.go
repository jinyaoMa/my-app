package service

import (
	"my-app/backend/internal/interfaces"
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/database/entity"
)

type LogService struct {
	interfaces.ILogService
}

func NewLogService(db *database.Database) interfaces.ILogService {
	return &LogService{
		ILogService: database.NewCrudService[*entity.Log](db),
	}
}
