package crud

import (
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/database/crud/interfaces"
	"my-app/backend/pkg/database/entity"
)

type CrudFile struct {
	interfaces.ICrud[*entity.File]
}

func NewCrudFile(database *database.Database) interfaces.ICrudFile {
	return &CrudFile{
		ICrud: NewCrud[*entity.File](database),
	}
}
