package crud

import (
	"my-app/backend/internal/entity"
	"my-app/backend/internal/interfaces"
	"my-app/backend/pkg/db"
)

type User struct {
	*db.CRUD[*entity.User]
}

func NewUser(dbs *db.DB) *User {
	return &User{
		CRUD: db.NewCRUD[*entity.User](dbs),
	}
}

func NewIUser(dbs *db.DB) interfaces.ICRUDUser {
	return NewUser(dbs)
}
