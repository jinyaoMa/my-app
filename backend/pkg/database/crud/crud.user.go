package crud

import (
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/database/crud/interfaces"
	"my-app/backend/pkg/database/entity"
)

type CrudUser struct {
	interfaces.ICrud[*entity.User]
}

func NewCrudUser(database *database.Database) interfaces.ICrudUser {
	return &CrudUser{
		ICrud: NewCrud[*entity.User](database),
	}
}
