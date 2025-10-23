package apir

import (
	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/db/crud"
	"majinyao.cn/my-app/backend/pkg/storage"
)

type Resources struct {
	CrudOptions crud.Options
	Db          *gorm.DB
	Store       storage.IStorage
}
