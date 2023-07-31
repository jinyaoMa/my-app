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

	currentLanguage    string
	currentTranslation *Translation
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
	availLangs, translationMap := i18n.LoadI18n()
	if helper.Any(availLangs, func(e *assetsio.Lang) bool {
		return e.Code == cfg.Language
	}) {
		currentLanguage = cfg.Language
	} else if len(availLangs) > 0 {
		currentLanguage = availLangs[0].Code
	}

	optionService := service.NewOptionService(db)

	var displayLanguage string
	displayLanguage, err = optionService.GetByOptionName(vmodel.OptionNameDisplayLanguage)
	if err == nil && helper.Any(availLangs, func(e *assetsio.Lang) bool {
		return e.Code == displayLanguage
	}) {
		currentLanguage = displayLanguage
	}

	var ok bool
	if currentTranslation, ok = translationMap[currentLanguage]; !ok {
		currentTranslation = DefaultTranslation()
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

func Lang(langs ...string) string {
	if len(langs) > 0 {
		var ok bool
		if currentTranslation, ok = i18n.LoadTranslation(langs[0]); !ok {
			currentTranslation = DefaultTranslation()
		}
		currentLanguage = langs[0]
	}
	return currentLanguage
}

func T() (t *Translation) {
	return currentTranslation
}

func Web() server.Interface {
	return web
}
