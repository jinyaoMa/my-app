package app

import (
	"my-app/backend/configs"
	ientity "my-app/backend/internal/entity"
	"my-app/backend/pkg/crypto"
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/database/entity"
	"my-app/backend/pkg/helper"
	"my-app/backend/pkg/logger"
	"my-app/backend/pkg/snowflake"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

func initDB(cfg *configs.Configs) (db *database.Database, err error) {
	entity.Cipher(crypto.NewAesWithSalt(cfg.Database.CipherKey))

	idGen, err := snowflake.New(cfg.Database.Snowflake)
	if err != nil {
		return
	}
	entity.IdGenerator(idGen)

	var mainDbFilename, logDbFilename, optionDbFilename string
	mainDbFilename, err = helper.GetFilenameSameAsExecutable("db")
	if err != nil {
		return
	}
	logDbFilename, err = helper.GetFilenameSameAsExecutable("logs.db")
	if err != nil {
		return
	}
	optionDbFilename, err = helper.GetFilenameSameAsExecutable("options.db")
	if err != nil {
		return
	}

	var logFile *os.File
	logFile, err = os.OpenFile(cfg.Database.LogFile, os.O_APPEND|os.O_CREATE, os.ModeAppend)
	if err != nil {
		return
	}

	db, err = database.New(&database.Option{
		Dialector: sqlite.Open(mainDbFilename + "?_pragma=foreign_keys(1)"),
		OnInitialized: func(db *gorm.DB) {
			logs := new(entity.Log)
			options := new(entity.Option)
			db.Use(dbresolver.Register(dbresolver.Config{
				Sources: []gorm.Dialector{sqlite.Open(logDbFilename)},
			}, logs).Register(dbresolver.Config{
				Sources: []gorm.Dialector{sqlite.Open(optionDbFilename)},
			}, options))
			db.Clauses(dbresolver.Use("logs")).AutoMigrate(logs)
			db.Clauses(dbresolver.Use("options")).AutoMigrate(options)
		},
		Migrate: []interface{}{
			new(ientity.Node),
		},
		Logger: database.OptionLogger{
			Option: logger.Option{
				Writer: logFile,
				Tag:    "DBS",
			},
		},
	})
	if err != nil {
		return
	}

	return
}
