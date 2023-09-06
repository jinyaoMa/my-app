package crud

import (
	"my-app/backend/internal/entity"
	"my-app/backend/internal/interfaces"
	"my-app/backend/pkg/db"
)

type File struct {
	*db.CRUD[*entity.File]
}

func NewFile(dbs *db.DB) *File {
	return &File{
		CRUD: db.NewCRUD[*entity.File](dbs),
	}
}

func NewIFile(dbs *db.DB) interfaces.ICRUDFile {
	return NewFile(dbs)
}
