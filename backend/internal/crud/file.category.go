package crud

import (
	"my-app/backend/internal/entity"
	"my-app/backend/internal/interfaces"
	"my-app/backend/pkg/db"
)

type FileCategory struct {
	*db.CRUD[*entity.FileCategory]
}

func NewFileCategory(dbs *db.DB) *FileCategory {
	return &FileCategory{
		CRUD: db.NewCRUD[*entity.FileCategory](dbs),
	}
}

func NewIFileCategory(dbs *db.DB) interfaces.ICRUDFileCategory {
	return NewFileCategory(dbs)
}
