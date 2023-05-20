package crud

import (
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/database/crud/interfaces"
	"my-app/backend/pkg/database/entity"
)

type CrudUserFile struct {
	interfaces.ICrud[*entity.UserFile]
}

func NewCrudUserFile(database *database.Database) interfaces.ICrudUserFile {
	return &CrudUserFile{
		ICrud: NewCrud[*entity.UserFile](database),
	}
}
