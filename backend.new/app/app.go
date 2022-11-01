package app

import (
	"context"
	"sync"

	"my-app/backend.new/app/i18n"
	"my-app/backend.new/model"

	"gorm.io/gorm"
)

var (
	instance *app
	once     sync.Once
)

type app struct {
	env  *Env
	db   *gorm.DB
	cfg  *Config
	log  *Logger
	i18n *i18n.I18n
	ctx  context.Context // wails context
}

// App app initialized as flow: load env -> connect db -> load cfg -> setup log -> setup i18n,
// application panic instead of logging before log setup
func App() *app {
	once.Do(func() {
		env := LoadEnv()
		db := ConnectDatabase()
		cfg := LoadConfig(db)

		var log *Logger
		if env.IsLog2File() {
			log = NewFileLogger(cfg.Get(model.OptionNameFileLog))
		} else {
			log = NewConsoleLogger()
		}
		db.Logger = log.database

		i18n := i18n.NewI18n(cfg.Get(model.OptionNameDirLanguages), log.i18n)

		// adjust config: color theme
		switch cfg.Get(model.OptionNameColorTheme) {
		case string(ConfigOptionColorThemeSystem):
		case string(ConfigOptionColorThemeLight):
		case string(ConfigOptionColorThemeDark):
		default:
			cfg.Set(model.OptionNameColorTheme, string(ConfigOptionColorThemeSystem))
		}
		// adjust config: display language
		availableLanguages := i18n.AvailableLanguages()
		if !i18n.HasLanguage(cfg.Get(model.OptionNameDisplayLanguage)) && len(availableLanguages) > 0 {
			cfg.Set(model.OptionNameDisplayLanguage, availableLanguages[0])
		}

		// initialize app
		instance = &app{
			env:  env,
			db:   db,
			cfg:  cfg,
			log:  log,
			i18n: i18n,
		}
	})
	return instance
}

func (a *app) UseEnv(callback func(env *Env)) *app {
	callback(a.env)
	return a
}

func (a *app) UseDatabase(callback func(db *gorm.DB)) *app {
	callback(a.db)
	return a
}

func (a *app) UseConfig(callback func(cfg *Config)) *app {
	callback(a.cfg)
	return a
}

func (a *app) Log() *Logger {
	return a.log
}

// T -> get current translation
func (a *app) UseI18n(callback func(T func() *i18n.Translation, i18n *i18n.I18n)) *app {
	callback(func() *i18n.Translation {
		return a.i18n.Translation(a.cfg.Get(model.OptionNameDisplayLanguage))
	}, a.i18n)
	return a
}

func (a *app) UseConfigAndI18n(callback func(cfg *Config, T func() *i18n.Translation, i18n *i18n.I18n)) *app {
	a.UseConfig(func(cfg *Config) {
		a.UseI18n(func(T func() *i18n.Translation, i18n *i18n.I18n) {
			callback(cfg, T, i18n)
		})
	})
	return a
}

func (a *app) SetContext(ctx context.Context) *app {
	a.ctx = ctx
	return a
}
