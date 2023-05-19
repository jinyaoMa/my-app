package crud

import (
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/database/entity"
	"my-app/backend/pkg/database/interfaces/crud"
)

type CrudFileExtension struct {
	crud.ICrud[*entity.FileExtension]
}

func NewCrudFileExtension(database *database.Database) crud.ICrudFileExtension {
	return &CrudFileExtension{
		ICrud: NewCrud[*entity.FileExtension](database),
	}
}
