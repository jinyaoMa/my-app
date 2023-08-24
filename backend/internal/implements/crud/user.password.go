package crud

import (
	"my-app/backend/internal/entity"
	"my-app/backend/internal/interfaces"
	"my-app/backend/pkg/db"
)

type CRUDUserPassword struct {
	*db.CRUD[*entity.UserPassword]
}

func NewCRUDUserPassword(dbs *db.DB) interfaces.ICRUDUserPassword {
	return &CRUDUserPassword{
		CRUD: db.NewCRUD[*entity.UserPassword](dbs),
	}
}
