package app

import (
	"my-app/backend/configs"
	"my-app/backend/pkg/assetsio"
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/helper"
	"my-app/backend/pkg/logger"
	"my-app/backend/pkg/server"
)

var (
	cfg    *configs.Configs
	db     *database.Database
	log    logger.Interface
	web    server.Interface
	assets assetsio.Interface
	i18n   assetsio.II18n[*Translation]
)

func init() {
	var err error

	var iniPath string
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

	web = server.New()

	assets = assetsio.New(cfg.AssetsPath)

	i18n = assetsio.NewI18n[*Translation](cfg.LanguagesPath)
}

func Cfg() *configs.Configs {
	return cfg
}

func Db() *database.Database {
	return db
}

func Log() logger.Interface {
	return log
}

func Web() server.Interface {
	return web
}

func Assets() assetsio.Interface {
	return assets
}

func I18n() assetsio.II18n[*Translation] {
	return i18n
}
