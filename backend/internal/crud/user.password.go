package crud

import (
	"my-app/backend/internal/entity"
	"my-app/backend/internal/interfaces"
	"my-app/backend/pkg/db"
)

type UserPassword struct {
	*db.CRUD[*entity.UserPassword]
}

func NewUserPassword(dbs *db.DB) *UserPassword {
	return &UserPassword{
		CRUD: db.NewCRUD[*entity.UserPassword](dbs),
	}
}

func NewIUserPassword(dbs *db.DB) interfaces.ICRUDUserPassword {
	return NewUserPassword(dbs)
}
