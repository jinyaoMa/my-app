package app

import (
	"my-app/backend/configs"
	"my-app/backend/internal/entity"
	"my-app/backend/pkg/db"
	"my-app/backend/pkg/db/param"
	"my-app/backend/pkg/enc"
	"my-app/backend/pkg/funcs"
	"my-app/backend/pkg/id"
	"my-app/backend/pkg/log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

func initDB(cfg *configs.Configs) (dbs *db.DB, err error) {
	var mainDbFilename string
	mainDbFilename, err = funcs.GetFilenameSameAsExecutable("db")
	if err != nil {
		return
	}

	var out log.ITreeWriter
	out, err = log.NewFileLogWriter(cfg.Database.LogFile, log.NewConsoleLogWriter())
	if err != nil {
		return
	}

	idGenerator, err := id.NewIID(cfg.Database.IdGenerator)
	if err != nil {
		return
	}

	dataCipher := enc.NewAesWithSalt(cfg.Database.CipherKey)

	dbs, err = db.New(&db.Config{
		Dialector: sqlite.Open(mainDbFilename + "?_pragma=foreign_keys(1)"),
		Options:   []gorm.Option{},
		OnInitialized: func(dbs *db.DB) (err error) {
			err = resolve(dbs)
			if err != nil {
				return
			}
			err = dbs.AutoMigrate(entities()...)
			if err != nil {
				return
			}
			err = dbs.SetupJoinTables(joinTables()...)
			if err != nil {
				return
			}
			return
		},
		Logger: log.New(&log.Config{
			Out:    out,
			Prefix: "[DBS] ",
			Flag:   log.DefaultFlag,
		}),
		LoggerConfig: gormLogger.Config{
			SlowThreshold:             0,
			Colorful:                  false,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      false,
			LogLevel:                  gormLogger.Error,
		},
		IdGenerator:   idGenerator,
		CodeGenerator: nil,
		DataCipher:    dataCipher,
	})
	return
}

func resolve(dbs *db.DB) (err error) {
	var logDbFilename, optionDbFilename string
	logDbFilename, err = funcs.GetFilenameSameAsExecutable("logs.db")
	if err != nil {
		return
	}
	optionDbFilename, err = funcs.GetFilenameSameAsExecutable("options.db")
	if err != nil {
		return
	}

	logs := new(entity.Log)
	options := new(entity.Option)
	dbs.Use(dbresolver.Register(dbresolver.Config{
		Sources: []gorm.Dialector{sqlite.Open(logDbFilename)},
	}, logs).Register(dbresolver.Config{
		Sources: []gorm.Dialector{sqlite.Open(optionDbFilename)},
	}, options))
	dbs.Clauses(dbresolver.Use("logs")).AutoMigrate(logs)
	dbs.Clauses(dbresolver.Use("options")).AutoMigrate(options)
	return
}

func entities() []any {
	return []any{
		&entity.User{},
		&entity.UserPassword{},
		&entity.UserFile{},
		&entity.File{},
		&entity.FileCategory{},
		&entity.FileExtension{},
		&entity.Node{},
	}
}

func joinTables() []param.JoinTable {
	userFile := new(entity.UserFile)
	return []param.JoinTable{
		{
			From:  &entity.User{},
			Field: "AccessableFiles",
			To:    userFile,
		},
		{
			From:  &entity.File{},
			Field: "AccessableUsers",
			To:    userFile,
		},
	}
}
