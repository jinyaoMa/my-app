package crud

import (
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/database/entity"
	"my-app/backend/pkg/database/interfaces/crud"
)

type CrudLog struct {
	crud.ICrud[*entity.Log]
}

func NewCrudLog(database *database.Database) crud.ICrudLog {
	return &CrudLog{
		ICrud: NewCrud[*entity.Log](database),
	}
}
