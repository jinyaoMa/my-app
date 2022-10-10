package database

import (
	"my-app/backend/pkg/utils"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	db *gorm.DB
)

func init() {
	var err error
	db, err = gorm.Open(
		sqlite.Open(utils.GetExecutablePath("MyApp.db")),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: "my_",
			},
			FullSaveAssociations: false,
			PrepareStmt:          true,
		},
	)
	if err != nil {
		panic("failed to connect database")
	}
}

func DB() *gorm.DB {
	return db
}

func SetLogger(logger logger.Interface) {
	db.Logger = logger
}
