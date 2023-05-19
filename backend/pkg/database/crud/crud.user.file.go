package crud

import (
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/database/entity"
	"my-app/backend/pkg/database/interfaces/crud"
)

type CrudUserFile struct {
	crud.ICrud[*entity.UserFile]
}

func NewCrudUserFile(database *database.Database) crud.ICrudUserFile {
	return &CrudUserFile{
		ICrud: NewCrud[*entity.UserFile](database),
	}
}
