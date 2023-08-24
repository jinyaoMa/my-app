package crud

import (
	"my-app/backend/internal/entity"
	"my-app/backend/internal/interfaces"
	"my-app/backend/pkg/db"
)

type CRUDUser struct {
	*db.CRUD[*entity.User]
}

func NewCRUDUser(dbs *db.DB) interfaces.ICRUDUser {
	return &CRUDUser{
		CRUD: db.NewCRUD[*entity.User](dbs),
	}
}
