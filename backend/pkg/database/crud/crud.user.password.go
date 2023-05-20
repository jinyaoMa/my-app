package crud

import (
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/database/crud/interfaces"
	"my-app/backend/pkg/database/entity"
)

type CrudUserPassword struct {
	interfaces.ICrud[*entity.UserPassword]
}

func NewCrudUserPassword(database *database.Database) interfaces.ICrudUserPassword {
	return &CrudUserPassword{
		ICrud: NewCrud[*entity.UserPassword](database),
	}
}
