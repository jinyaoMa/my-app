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
	env    *Env
	config *Config
}

func init() {
	instance = &app{
		env:    LoadEnv(),
		logger: LoadConsoleLogger(),
	}
}

func App() *app {
	once.Do(func() {
		instance.config = LoadConfig()

		if instance.env.Log2File() {
			logFile, err := os.OpenFile(
				instance.config.LogPath,
				os.O_CREATE|os.O_WRONLY|os.O_APPEND,
				0666,
			)
			if err != nil {
				instance.logger.App.Fatalf("failed to open log file: %+v\n", err)
			}
			instance.logger = LoadFileLogger(logFile)
		}

		model.SetLogger(instance.logger.Model)
	})
	return instance
}

func (a *app) Config() *Config {
	return a.config
}

func (a *app) WebConfig() *WebConfig {
	return a.config.Web
}

func (a *app) Env() *Env {
	return a.env
}

func (a *app) WebLog() *log.Logger {
	return a.logger.Web
}

func (a *app) TrayLog() *log.Logger {
	return a.logger.Tray
}

func (a *app) WailsLog() *log.WailsLogger {
	return a.logger.Wails
}
