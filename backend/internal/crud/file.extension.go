package crud

import (
	"my-app/backend/internal/entity"
	"my-app/backend/internal/interfaces"
	"my-app/backend/pkg/db"
)

type FileExtension struct {
	*db.CRUD[*entity.FileExtension]
}

func NewFileExtension(dbs *db.DB) *FileExtension {
	return &FileExtension{
		CRUD: db.NewCRUD[*entity.FileExtension](dbs),
	}
}

func NewIFileExtension(dbs *db.DB) interfaces.ICRUDFileExtension {
	return NewFileExtension(dbs)
}
