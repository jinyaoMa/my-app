package database

import (
	"my-app/backend/pkg"
	"my-app/backend/pkg/utils"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	db  *gorm.DB
	aes *utils.AES
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

	aes = utils.NewAES(pkg.Copyright)
}

func DB() *gorm.DB {
	return db
}

func AES() *utils.AES {
	return aes
}

func SetLogger(logger logger.Interface) {
	db.Logger = logger
}

func SetAES(a *utils.AES) {
	aes = a
}
