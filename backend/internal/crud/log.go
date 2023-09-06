package crud

import (
	"my-app/backend/internal/entity"
	"my-app/backend/internal/interfaces"
	"my-app/backend/pkg/db"
)

type Log struct {
	*db.CRUD[*entity.Log]
}

func NewLog(dbs *db.DB) *Log {
	return &Log{
		CRUD: db.NewCRUD[*entity.Log](dbs),
	}
}

func NewILog(dbs *db.DB) interfaces.ICRUDLog {
	return NewLog(dbs)
}
