package db

import (
	"my-app/backend/pkg/database/entity"

	"gorm.io/gorm"
)

func migrate(db *gorm.DB, dst ...interface{}) error {
	dst = append(dst, []interface{}{
		new(entity.File),
		new(entity.FileCategory),
		new(entity.FileExtension),
		new(entity.Log),
		new(entity.Option),
		new(entity.User),
		new(entity.UserFile),
		new(entity.UserPassword),
	}...)
	return db.AutoMigrate(dst...)
}
