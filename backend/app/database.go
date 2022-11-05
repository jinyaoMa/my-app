package app

import (
	"my-app/backend/model"
	"my-app/backend/utils"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// ConnectDatabase get database connection
func ConnectDatabase() *gorm.DB {
	// connect database
	db, err := gorm.Open(
		sqlite.Open(utils.Utils().GetExecutablePath(utils.Utils().GetExecutableFileName("db"))),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: "my_",
			},
			FullSaveAssociations: false,
			PrepareStmt:          true,
		},
	)
	if err != nil {
		utils.Utils().Panic("failed to connect database: " + err.Error())
	}

	// enable foreign_keys for SQLite
	if res := db.Exec("PRAGMA foreign_keys = ON"); res.Error != nil {
		utils.Utils().Panic("failed to enable foreign_keys: " + res.Error.Error())
	}

	// migrate tables
	if err := db.AutoMigrate(
		&model.Option{},
		&model.User{},
		&model.Keyring{},
	); err != nil {
		utils.Utils().Panic("failed to auto migrate: " + err.Error())
	}

	return db
}
