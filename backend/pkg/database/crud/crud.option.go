package crud

import (
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/database/crud/interfaces"
	"my-app/backend/pkg/database/entity"
)

type CrudOption struct {
	interfaces.ICrud[*entity.Option]
}

func NewCrudOption(database *database.Database) interfaces.ICrudOption {
	return &CrudOption{
		ICrud: NewCrud[*entity.Option](database),
	}
}
