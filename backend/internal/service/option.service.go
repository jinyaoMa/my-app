package service

import (
	"my-app/backend/internal/interfaces"
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/database/entity"
)

type OptionService struct {
	interfaces.IOptionService
}

func NewOptionService(db *database.Database) interfaces.IOptionService {
	return &OptionService{
		IOptionService: database.NewCrudService[*entity.Option](db),
	}
}
