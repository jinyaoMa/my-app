package app

import (
	"my-app/backend/configs"
	"my-app/backend/internal/crud"
	"my-app/backend/internal/entity"
	"my-app/backend/internal/interfaces"
	"my-app/backend/internal/vmodel"
	"my-app/backend/pkg/aio"
	"my-app/backend/pkg/api"
	"my-app/backend/pkg/db"
	"my-app/backend/pkg/funcs"
	"my-app/backend/pkg/log"

	"github.com/wailsapp/wails/v2/pkg/options/windows"
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

	currentColorTheme *entity.Option
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

	crudOption = crud.NewOption(dbs)

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
	crudOption.Save(currentLanguage)

	var ok bool
	currentTranslation, ok = translationMap[currentLanguage.Value]
	if !ok {
		currentTranslation = DefaultTranslation()
	}

	currentColorTheme, err = crudOption.GetByOptionName(vmodel.OptionNameColorTheme)
	if err != nil {
		currentColorTheme = &entity.Option{
			Key:   vmodel.OptionNameColorTheme,
			Value: vmodel.OptionValueColorThemeString(windows.SystemDefault, vmodel.OptionValueColorThemeSystem),
		}
	}
	crudOption.Save(currentColorTheme)

	web = api.New()

	var webAutoStart, webSwagger, webVitePress *entity.Option
	webAutoStart, err = crudOption.GetByOptionName(vmodel.OptionNameWebAutoStart)
	if err != nil {
		webAutoStart = &entity.Option{
			Key:   vmodel.OptionNameWebAutoStart,
			Value: vmodel.OptionValueBoolString(true),
		}
	}
	webSwagger, err = crudOption.GetByOptionName(vmodel.OptionNameWebSwagger)
	if err != nil {
		webSwagger = &entity.Option{
			Key:   vmodel.OptionNameWebSwagger,
			Value: "https://localhost:10443/swagger/index.html",
		}
	}
	webVitePress, err = crudOption.GetByOptionName(vmodel.OptionNameWebVitePress)
	if err != nil {
		webVitePress = &entity.Option{
			Key:   vmodel.OptionNameWebVitePress,
			Value: "https://localhost:10443/doc/index.html",
		}
	}
	crudOption.SaveAll([]*entity.Option{
		webAutoStart,
		webSwagger,
		webVitePress,
	})
	if vmodel.OptionValueBool(webAutoStart.Value) {
		StartAPI()
	}
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

func LANG(l ...string) string {
	if len(l) > 0 {
		var ok bool
		if currentTranslation, ok = i18n.LoadTranslation(l[0]); ok {
			currentLanguage.Value = l[0]
			crudOption.Save(currentLanguage)
		}
	}
	return currentLanguage.Value
}

func T() (t *Translation) {
	return currentTranslation
}

func THEME(t ...windows.Theme) windows.Theme {
	if len(t) > 0 {
		currentColorTheme.Value = vmodel.OptionValueColorThemeString(t[0], vmodel.OptionValueColorThemeSystem)
		crudOption.Save(currentColorTheme)
	}
	return vmodel.OptionValueColorTheme(currentColorTheme.Value, windows.SystemDefault)
}

func API() api.IAPI {
	return web
}

func OPTION(key string, def string) string {
	opt, err := crudOption.GetByOptionName(key)
	if err != nil {
		return def
	}
	return opt.Value
}
