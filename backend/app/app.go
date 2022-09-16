package app

import (
	"my-app/backend/model"
	"my-app/backend/pkg/log"
	"os"
	"sync"
)

var (
	once     sync.Once
	instance *app
)

type app struct {
	logger *Logger
	flag   *Flag
	config *Config
}

func init() {
	instance = &app{
		logger: DefaultLogger(),
	}
}

func App() *app {
	once.Do(func() {
		flag := LoadFlag()
		cfg := LoadConfig()

		logFile, err := os.OpenFile(
			cfg.LogPath,
			os.O_CREATE|os.O_WRONLY|os.O_APPEND,
			0666,
		)
		if err != nil {
			instance.logger.App.Fatalf("failed to open log file: %+v\n", err)
		}

		logger := LoadLogger(logFile)
		model.SetLogger(logger.Model)

		instance = &app{
			logger: logger,
			flag:   flag,
			config: cfg,
		}
	})
	return instance
}

func (a *app) Config() *Config {
	return a.config
}

func (a *app) WebConfig() *WebConfig {
	return a.config.Web
}

func (a *app) Flag() *Flag {
	return a.flag
}

func (a *app) WebLog() *log.Logger {
	return a.logger.Web
}

func (a *app) TrayLog() *log.Logger {
	return a.logger.Tray
}

func (a *app) WailsLog() *log.Logger {
	return a.logger.Wails
}
