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
	"sync"
)

var (
	singleton *app
	once      sync.Once
)

type app struct {
	cfg    *configs.Configs
	db     *database.Database
	log    logger.Interface
	assets assetsio.Interface
	i18n   assetsio.II18n[*Translation]
	web    server.Interface

	optionService      interfaces.IOptionService
	currentLanguage    *entity.Option
	currentTranslation *Translation
}

func App() *app {
	once.Do(func() {
		var err error
		if singleton, err = newApp(); err != nil {
			panic(err)
		}
	})
	return singleton
}

func newApp() (*app, error) {
	cfg, err := initCfg()
	if err != nil {
		return nil, err
	}

	db, err := initDB(cfg)
	if err != nil {
		return nil, err
	}

	log, err := initLog(cfg, db)
	if err != nil {
		return nil, err
	}

	assets := assetsio.New(cfg.AssetsPath)

	i18n := assetsio.NewI18n[*Translation](cfg.LanguagesPath)
	availLangs, translationMap := i18n.LoadI18n()

	optionService := service.NewOptionService(db)

	currentLanguage, err := optionService.GetByOptionName(vmodel.OptionNameDisplayLanguage)
	if err != nil || !helper.Any(availLangs, func(e *assetsio.Lang) bool {
		return e.Code == currentLanguage.Value
	}) {
		currentLanguage = &entity.Option{
			Key:   vmodel.OptionNameDisplayLanguage,
			Value: "",
		}
		if helper.Any(availLangs, func(e *assetsio.Lang) bool {
			return e.Code == cfg.Language
		}) {
			currentLanguage.Value = cfg.Language
		} else if len(availLangs) > 0 {
			currentLanguage.Value = availLangs[0].Code
		}
	}

	currentTranslation, ok := translationMap[currentLanguage.Value]
	if !ok {
		currentTranslation = DefaultTranslation()
	}

	web := server.New()

	return &app{
		cfg:                cfg,
		db:                 db,
		log:                log,
		assets:             assets,
		i18n:               i18n,
		web:                web,
		optionService:      optionService,
		currentLanguage:    currentLanguage,
		currentTranslation: currentTranslation,
	}, nil
}

func (a *app) Cfg() *configs.Configs {
	return a.cfg
}

func (a *app) Db() *database.Database {
	return a.db
}

func (a *app) Log() logger.Interface {
	return a.log
}

func (a *app) Assets() assetsio.Interface {
	return a.assets
}

func (a *app) I18n() assetsio.II18n[*Translation] {
	return a.i18n
}

func (a *app) Lang(langs ...string) string {
	if len(langs) > 0 {
		var ok bool
		if a.currentTranslation, ok = a.i18n.LoadTranslation(langs[0]); ok {
			a.currentLanguage.Value = langs[0]
			a.optionService.Save(a.currentLanguage)
		}
	}
	return a.currentLanguage.Value
}

func (a *app) T() (t *Translation) {
	return a.currentTranslation
}

func (a *app) Web() server.Interface {
	return a.web
}
