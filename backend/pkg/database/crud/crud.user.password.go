package crud

import (
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/database/entity"
	"my-app/backend/pkg/database/interfaces/crud"
)

type CrudUserPassword struct {
	crud.ICrud[*entity.UserPassword]
}

func NewCrudUserPassword(database *database.Database) crud.ICrudUserPassword {
	return &CrudUserPassword{
		ICrud: NewCrud(database, new(entity.UserPassword)),
	}
}
