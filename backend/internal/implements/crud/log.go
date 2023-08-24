package crud

import (
	"my-app/backend/internal/entity"
	"my-app/backend/internal/interfaces"
	"my-app/backend/pkg/db"
)

type CRUDLog struct {
	*db.CRUD[*entity.Log]
}

func NewCRUDLog(dbs *db.DB) interfaces.ICRUDLog {
	return &CRUDLog{
		CRUD: db.NewCRUD[*entity.Log](dbs),
	}
}
