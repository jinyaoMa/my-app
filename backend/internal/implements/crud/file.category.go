package crud

import (
	"my-app/backend/internal/entity"
	"my-app/backend/internal/interfaces"
	"my-app/backend/pkg/db"
)

type CRUDFileCategory struct {
	*db.CRUD[*entity.FileCategory]
}

func NewCRUDFileCategory(dbs *db.DB) interfaces.ICRUDFileCategory {
	return &CRUDFileCategory{
		CRUD: db.NewCRUD[*entity.FileCategory](dbs),
	}
}
