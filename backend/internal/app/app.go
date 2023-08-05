package app

import (
	"my-app/backend/configs"
	"my-app/backend/internal/interfaces"
	"my-app/backend/internal/service"
	"my-app/backend/internal/vmodel"
	"my-app/backend/pkg/assetsio"
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/database/entity"
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

	optionService      interfaces.IOptionService
	currentLanguage    *entity.Option
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

	optionService = service.NewOptionService(db)

	currentLanguage, err = optionService.GetByOptionName(vmodel.OptionNameDisplayLanguage)
	if err != nil || !helper.Any(availLangs, func(e assetsio.Lang) bool {
		return e.Code == currentLanguage.Value
	}) {
		currentLanguage = &entity.Option{
			Key:   vmodel.OptionNameDisplayLanguage,
			Value: "",
		}
		if helper.Any(availLangs, func(e assetsio.Lang) bool {
			return e.Code == cfg.Language
		}) {
			currentLanguage.Value = cfg.Language
		} else if len(availLangs) > 0 {
			currentLanguage.Value = availLangs[0].Code
		}
	}

	var ok bool
	currentTranslation, ok = translationMap[currentLanguage.Value]
	if !ok {
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
		if currentTranslation, ok = i18n.LoadTranslation(langs[0]); ok {
			currentLanguage.Value = langs[0]
			optionService.Save(currentLanguage)
		}
	}
	return currentLanguage.Value
}

func T() (t *Translation) {
	return currentTranslation
}

func Web() server.Interface {
	return web
}
