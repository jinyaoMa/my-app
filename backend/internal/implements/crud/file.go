package crud

import (
	"my-app/backend/internal/entity"
	"my-app/backend/internal/interfaces"
	"my-app/backend/pkg/db"
)

type CRUDFile struct {
	*db.CRUD[*entity.File]
}

func NewCRUDFile(dbs *db.DB) interfaces.ICRUDFile {
	return &CRUDFile{
		CRUD: db.NewCRUD[*entity.File](dbs),
	}
}
