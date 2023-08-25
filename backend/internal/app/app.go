package app

import (
	"my-app/backend/configs"
	"my-app/backend/internal/entity"
	"my-app/backend/internal/implements/crud"
	"my-app/backend/internal/interfaces"
	"my-app/backend/internal/vmodel"
	"my-app/backend/pkg/aio"
	"my-app/backend/pkg/api"
	"my-app/backend/pkg/db"
	"my-app/backend/pkg/funcs"
	"my-app/backend/pkg/log"

	"gorm.io/gorm"
)

var (
	cfg    *configs.Configs
	dbs    *db.DB
	logger *log.Log
	assets aio.IAIO
	i18n   aio.II18n[*Translation]
	web    api.IAPI

	crudOption         interfaces.ICRUDOption
	currentLanguage    *entity.Option
	currentTranslation *Translation
)

func init() {
	var err error

	cfg, err = initCfg()
	if err != nil {
		panic(err)
	}

	dbs, err = initDB(cfg)
	if err != nil {
		panic(err)
	}

	logger, err = initLog(cfg, dbs.Session(&gorm.Session{}))
	if err != nil {
		panic(err)
	}

	assets = aio.New(cfg.AssetsPath)

	i18n = aio.NewI18n[*Translation](cfg.LanguagesPath)
	availLangs, translationMap := i18n.LoadI18n()

	crudOption = crud.NewCRUDOption(dbs)

	currentLanguage, err = crudOption.GetByOptionName(vmodel.OptionNameDisplayLanguage)
	if err != nil || !funcs.Any(availLangs, func(e aio.Lang) bool {
		return e.Code == currentLanguage.Value
	}) {
		currentLanguage = &entity.Option{
			Key:   vmodel.OptionNameDisplayLanguage,
			Value: "",
		}
		if funcs.Any(availLangs, func(e aio.Lang) bool {
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

	web = api.New()
}

func CFG() *configs.Configs {
	return cfg
}

func DB() *db.DB {
	return dbs.Session(&gorm.Session{})
}

func LOG() *log.Log {
	return logger
}

func ASSETS() aio.IAIO {
	return assets
}

func I18N() aio.II18n[*Translation] {
	return i18n
}

func LANG(langs ...string) string {
	if len(langs) > 0 {
		var ok bool
		if currentTranslation, ok = i18n.LoadTranslation(langs[0]); ok {
			currentLanguage.Value = langs[0]
			crudOption.Save(currentLanguage)
		}
	}
	return currentLanguage.Value
}

func T() (t *Translation) {
	return currentTranslation
}

func API() api.IAPI {
	return web
}
