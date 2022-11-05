package app

import (
	"my-app/backend/utils"
	"os"

	wailsLogger "github.com/wailsapp/wails/v2/pkg/logger"
	gormLogger "gorm.io/gorm/logger"
)

const (
	LogPrefixDatabase = "DBS"
	LogPrefixWails    = "WLS"
	LogPrefixI18n     = "I18"
	LogPrefixTray     = "TRY"
	LogPrefixWeb      = "WEB"
	LogPrefixServices = "SEV"
)

type log struct {
	database gormLogger.Interface
	wails    wailsLogger.Logger
	i18n     *utils.Logger
	tray     *utils.Logger
	web      *utils.Logger
	services *utils.Logger
}

func NewConsoleLogger() *log {
	return &log{
		database: utils.NewGormConsoleLogger(LogPrefixDatabase),
		wails:    utils.NewWailsConsoleLogger(LogPrefixWails),
		i18n:     utils.NewConsoleLogger(LogPrefixI18n),
		tray:     utils.NewConsoleLogger(LogPrefixTray),
		web:      utils.NewConsoleLogger(LogPrefixWeb),
		services: utils.NewConsoleLogger(LogPrefixServices),
	}
}

func NewFileLogger(logPath string) *log {
	logFile, err := os.OpenFile(
		logPath,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		utils.Utils().Panic("failed to open log file: " + err.Error())
	}
	return &log{
		database: utils.NewGormFileLogger(LogPrefixDatabase, logFile),
		wails:    utils.NewWailsFileLogger(LogPrefixWails, logFile),
		i18n:     utils.NewFileLogger(LogPrefixI18n, logFile),
		tray:     utils.NewFileLogger(LogPrefixTray, logFile),
		web:      utils.NewFileLogger(LogPrefixWeb, logFile),
		services: utils.NewFileLogger(LogPrefixServices, logFile),
	}
}

func (l *log) Wails() wailsLogger.Logger {
	return l.wails
}

func (l *log) Tray() *utils.Logger {
	return l.tray
}

func (l *log) Web() *utils.Logger {
	return l.web
}

func (l *log) Services() *utils.Logger {
	return l.services
}
