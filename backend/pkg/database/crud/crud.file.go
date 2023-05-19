package crud

import (
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/database/entity"
	"my-app/backend/pkg/database/interfaces/crud"
)

type CrudFile struct {
	crud.ICrud[*entity.File]
}

func NewCrudFile(database *database.Database) crud.ICrudFile {
	return &CrudFile{
		ICrud: NewCrud[*entity.File](database),
	}
}
