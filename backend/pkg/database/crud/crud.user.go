package crud

import (
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/database/entity"
	"my-app/backend/pkg/database/interfaces/crud"
)

type CrudUser struct {
	crud.ICrud[*entity.User]
}

func NewCrudUser(database *database.Database) crud.ICrudUser {
	return &CrudUser{
		ICrud: NewCrud(database, new(entity.User)),
	}
}
