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
		utils.Utils().PanicLogger().Fatalf("failed to connect database: %+v\n", err)
	}

	// enable foreign_keys for SQLite
	if res := db.Exec("PRAGMA foreign_keys = ON"); res.Error != nil {
		utils.Utils().PanicLogger().Fatalf("failed to enable foreign_keys: %+v\n", res.Error)
	}

	// migrate tables
	if err := db.AutoMigrate(
		&model.Option{},
		&model.User{},
		&model.Keyring{},
	); err != nil {
		utils.Utils().PanicLogger().Fatalf("failed to auto migrate: %+v\n", err)
	}

	return db
}
