package database

import (
	"my-app/backend/pkg/database/entity"

	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

func migrate(db *gorm.DB, dst ...any) error {
	dst = append(dst, []interface{}{
		new(entity.File),
		new(entity.FileCategory),
		new(entity.FileExtension),
		new(entity.Log),
		new(entity.Node),
		new(entity.Option),
		new(entity.User),
		new(entity.UserFile),
		new(entity.UserPassword),
	}...)
	db.Clauses(dbresolver.Use("logs")).AutoMigrate(new(entity.Log))
	return db.AutoMigrate(dst...)
}
