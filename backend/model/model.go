package model

import (
	"log"
	"my-app/backend/pkg/utils"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	var err error
	db, err = gorm.Open(
		sqlite.Open(utils.GetExecutablePath("my-app.db")),
		&gorm.Config{
			FullSaveAssociations: false,
			PrepareStmt:          true,
		},
	)
	if err != nil {
		log.Fatalf("failed to connect database: %+v\n", err)
	}

	db.AutoMigrate(
		&MyOption{},
	)
}
