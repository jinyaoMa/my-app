package app

import (
	"my-app/backend/app/config"
	"my-app/backend/app/i18n"
	"my-app/backend/app/logger"
	"my-app/backend/database"
)

var (
	instance *app
)

type app struct {
	cfg  *config.Config
	env  *config.Env
	log  *logger.Logger
	i18n *i18n.I18n
}

func init() {
	var log *logger.Logger
	cfg := config.LoadConfig()
	env := config.LoadEnv().
		Log2Console(func() {
			log = logger.NewConsoleLogger()
		}).
		Log2File(func() {
			log = logger.NewFileLogger(cfg.LogPath)
		})

	instance = &app{
		cfg:  cfg,
		env:  env,
		log:  log,
		i18n: i18n.NewI18n(),
	}

	database.SetLogger(instance.log.Database())
}

func App() *app {
	return instance
}

func (a *app) Config() *config.Config {
	return a.cfg
}

func (a *app) Env() *config.Env {
	return a.env
}

func (a *app) Log() *logger.Logger {
	return a.log
}

func (a *app) CurrentTranslation() *i18n.Translation {
	if t := a.i18n.Translation(a.cfg.DisplayLanguage); t != nil {
		return t
	}
	return &i18n.Translation{}
}
