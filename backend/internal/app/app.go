package app

import (
	"my-app/backend/configs"
	"my-app/backend/internal/service"
	"my-app/backend/internal/vmodel"
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
	assets assetsio.Interface
	i18n   assetsio.II18n[*Translation]
	web    server.Interface

	lang string
)

func init() {
	var err error

	cfg, err = initCfg()
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

	assets = assetsio.New(cfg.AssetsPath)

	i18n = assetsio.NewI18n[*Translation](cfg.LanguagesPath)

	optionService := service.NewOptionService(db)
	lang, err = optionService.GetByOptionName(vmodel.OptionNameDisplayLanguage)
	if err == nil {
		availLangs, _ := i18n.LoadI18n()
		if len(availLangs) > 0 && !helper.Any(availLangs, func(e *assetsio.Lang) bool {
			return e.Code == lang
		}) {
			lang = availLangs[0].Code
		}
	} else {

	}

	web = server.New()
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

func Assets() assetsio.Interface {
	return assets
}

func I18n() assetsio.II18n[*Translation] {
	return i18n
}

func Web() server.Interface {
	return web
}
