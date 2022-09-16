package model

import (
	"context"
	"my-app/backend/pkg/utils"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db *gorm.DB
)

func init() {
	var err error
	db, err = gorm.Open(
		sqlite.Open(utils.GetExecutablePath("MyApp.db")),
		&gorm.Config{
			FullSaveAssociations: false,
			PrepareStmt:          true,
		},
	)
	if err != nil {
		db.Logger.Error(context.Background(), "failed to connect database: %+v\n", err)
	}

	db.AutoMigrate(
		&MyOption{},
	)
}

func SetLogger(logger logger.Interface) {
	db.Logger = logger
}
