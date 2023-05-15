package engine

import (
	"my-app/backend/pkg/database/entity"

	"gorm.io/gorm"
)

func migrate(db *gorm.DB, dst ...any) error {
	dst = append(dst, []interface{}{
		new(entity.Option),
		new(entity.Log),
		new(entity.User),
		new(entity.UserPassword),
	}...)
	return db.AutoMigrate(dst...)
}
