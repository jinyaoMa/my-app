package crud

import (
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/database/crud/interfaces"
	"my-app/backend/pkg/database/entity"
)

type CrudFileExtension struct {
	interfaces.ICrud[*entity.FileExtension]
}

func NewCrudFileExtension(database *database.Database) interfaces.ICrudFileExtension {
	return &CrudFileExtension{
		ICrud: NewCrud[*entity.FileExtension](database),
	}
}
