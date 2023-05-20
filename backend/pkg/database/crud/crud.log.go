package crud

import (
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/database/crud/interfaces"
	"my-app/backend/pkg/database/entity"
)

type CrudLog struct {
	interfaces.ICrud[*entity.Log]
}

func NewCrudLog(database *database.Database) interfaces.ICrudLog {
	return &CrudLog{
		ICrud: NewCrud[*entity.Log](database),
	}
}
