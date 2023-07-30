package app

import (
	"my-app/backend/configs"
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/helper"
	"my-app/backend/pkg/logger"
)

var (
	iniPath string
	cfg     *configs.Configs
	db      *database.Database
	log     logger.Interface
)

func init() {
	var err error

	iniPath, err = helper.GetFilenameSameAsExecutable("config.ini")
	if err != nil {
		panic(err)
	}

	cfg, err = configs.NewConfigs(iniPath)
	if err != nil {
		panic(err)
	}

	db, err = initDB(cfg)
	if err != nil {
		panic(err)
	}

	log, err = initLog(cfg, db)
	if err != nil {
		panic(err)
	}
}

func Cfg() *configs.Configs {
	return cfg
}

func DB() *database.Database {
	return db
}

func Log() logger.Interface {
	return log
}
