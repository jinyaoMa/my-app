package app

import (
	"my-app/backend.new/model"
	"my-app/backend.new/utils"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// ConnectDatabase get database connection
func ConnectDatabase() *gorm.DB {
	// connect database
	db, err := gorm.Open(
		sqlite.Open(utils.Utils().GetExecutablePath("MyApp.db")),
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

	// enable foreign_keys for SQLite
	if res := db.Exec("PRAGMA foreign_keys = ON"); res.Error != nil {
		panic("failed to enable foreign_keys")
	}

	// migrate tables
	if db.AutoMigrate(
		&model.Option{},
		&model.User{},
		&model.Keyring{},
	) != nil {
		panic("failed to auto migrate")
	}

	return db
}
