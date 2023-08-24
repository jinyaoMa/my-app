package crud

import (
	"my-app/backend/internal/entity"
	"my-app/backend/internal/interfaces"
	"my-app/backend/pkg/db"
)

type CRUDFileExtension struct {
	*db.CRUD[*entity.FileExtension]
}

func NewCRUDFileExtension(dbs *db.DB) interfaces.ICRUDFileExtension {
	return &CRUDFileExtension{
		CRUD: db.NewCRUD[*entity.FileExtension](dbs),
	}
}
