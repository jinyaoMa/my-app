package crud

import (
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/database/entity"
	"my-app/backend/pkg/database/interfaces/crud"
)

type CrudFileCategory struct {
	crud.ICrud[*entity.FileCategory]
}

func NewCrudFileCategory(database *database.Database) crud.ICrudFileCategory {
	return &CrudFileCategory{
		ICrud: NewCrud[*entity.FileCategory](database),
	}
}
