package crud

import (
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/database/crud/interfaces"
	"my-app/backend/pkg/database/entity"
)

type CrudFileCategory struct {
	interfaces.ICrud[*entity.FileCategory]
}

func NewCrudFileCategory(database *database.Database) interfaces.ICrudFileCategory {
	return &CrudFileCategory{
		ICrud: NewCrud[*entity.FileCategory](database),
	}
}
