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
	i18n *i18n.I18n
	log  *Logger
	ctx  context.Context // wails context
}

// App app initialized as flow: load env -> connect db -> load cfg -> load i18n -> setup log
func App() *app {
	once.Do(func() {
		env := LoadEnv()
		db := ConnectDatabase()
		cfg := LoadConfig(db)
		i18n := i18n.NewI18n(cfg.Get(model.OptionDirLanguages))

		var log *Logger
		env.Log2Console(func() {
			log = NewConsoleLogger()
		}).Log2File(func() {
			log = NewFileLogger(cfg.Get(model.OptionFileLog))
		})
		db.Logger = log.Database()

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

func (a *app) UseDB(callback func(db *gorm.DB)) *app {
	callback(a.db)
	return a
}

func (a *app) UseCfg(callback func(cfg *Config)) *app {
	callback(a.cfg)
	return a
}

func (a *app) Log() *Logger {
	return a.log
}

func (a *app) UseI18n(callback func(i18n *i18n.I18n)) *app {
	callback(a.i18n)
	return a
}

func (a *app) CurrentTranslation() *i18n.Translation {
	return a.i18n.Translation(a.cfg.Get(model.OptionDisplayLanguage))
}

func (a *app) SetContext(ctx context.Context) *app {
	a.ctx = ctx
	return a
}
