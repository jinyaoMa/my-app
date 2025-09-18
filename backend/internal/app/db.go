package app

import (
	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/cflog"
	"majinyao.cn/my-app/backend/pkg/db"
)

func initDB(log *cflog.Cflog, options db.Options) *gorm.DB {
	var err error
	DB, err = db.Open([]any{
		new(entity.File),
		new(entity.FileCategory),
		new(entity.FileExtension),
		new(entity.FileGroup),
		new(entity.FileUser),
		new(entity.Group),
		new(entity.GroupRole),
		new(entity.GroupUser),
		new(entity.OperationIdEnumPair),
		new(entity.Option),
		new(entity.Permission),
		new(entity.Role),
		new(entity.RolePermission),
		new(entity.User),
		new(entity.UserPassword),
		new(entity.UserRole),
	}, options)
	if err != nil {
		log.Panicln("init db failed", err)
	}
	return DB
}
