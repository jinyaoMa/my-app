package app

import (
	"my-app/backend/configs"
	"my-app/backend/internal/crud"
	"my-app/backend/internal/entity"
	"my-app/backend/internal/interfaces"
	"my-app/backend/pkg/aio"
	"my-app/backend/pkg/api"
	"my-app/backend/pkg/db"
	"my-app/backend/pkg/log"
	"strconv"

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

	currentColorTheme_ windows.Theme
	currentColorTheme  *entity.Option
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

	_, currentLanguage, err = crudOption.GetOrSaveDisplayLanguageByOptionName(crud.OptionNameDisplayLanguage, availLangs, cfg.Language)
	if err != nil {
		panic(err)
	}

	var ok bool
	currentTranslation, ok = translationMap[currentLanguage.Value]
	if !ok {
		currentTranslation = DefaultTranslation()
	}

	currentColorTheme_, currentColorTheme, err = crudOption.GetOrSaveColorThemeByOptionName(crud.OptionNameColorTheme, windows.SystemDefault)
	if err != nil {
		panic(err)
	}

	web = api.New()
	webAutoStart, _, err := crudOption.GetOrSaveBoolByOptionName(crud.OptionNameWebAutoStart, true)
	if err != nil {
		panic(err)
	}
	if webAutoStart {
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
		if o, err := crudOption.GetById(currentLanguage.ID); err == nil {
			currentLanguage = o
			if t, ok := i18n.LoadTranslation(l[0]); ok {
				currentTranslation = t
				currentLanguage.Value = l[0]
				crudOption.Save(currentLanguage)
			}
		}
	}
	return currentLanguage.Value
}

func T() (t *Translation) {
	return currentTranslation
}

func THEME(t ...windows.Theme) windows.Theme {
	if len(t) > 0 {
		if o, err := crudOption.GetById(currentColorTheme.ID); err == nil {
			currentColorTheme_ = t[0]
			currentColorTheme = o
			currentColorTheme.Value = strconv.FormatInt(int64(currentColorTheme_), 10)
			crudOption.Save(currentColorTheme)
		}
	}
	return currentColorTheme_
}

func API() api.IAPI {
	return web
}

func OPTION() interfaces.ICRUDOption {
	return crudOption
}
