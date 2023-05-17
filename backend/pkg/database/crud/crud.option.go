package crud

import (
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/database/entity"
	"my-app/backend/pkg/database/interfaces/crud"
)

type CrudOption struct {
	crud.ICrud[*entity.Option]
}

func NewCrudOption(database *database.Database) crud.ICrudOption {
	return &CrudOption{
		ICrud: NewCrud[*entity.Option](database),
	}
}
