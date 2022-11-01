package app

import (
	"my-app/backend.new/utils"
	"os"

	wailsLogger "github.com/wailsapp/wails/v2/pkg/logger"
	gormLogger "gorm.io/gorm/logger"
)

const (
	LoggerPrefixDatabase = "DBS"
	LoggerPrefixWails    = "WLS"
	LoggerPrefixI18n     = "I18"
	LoggerPrefixTray     = "TRY"
	LoggerPrefixWeb      = "WEB"
	LoggerPrefixService  = "SEV"
)

type Logger struct {
	database gormLogger.Interface
	wails    wailsLogger.Logger
	i18n     *utils.Logger
	tray     *utils.Logger
	web      *utils.Logger
	service  *utils.Logger
}

func NewConsoleLogger() *Logger {
	return &Logger{
		database: utils.NewGormConsoleLogger(LoggerPrefixDatabase),
		wails:    utils.NewWailsConsoleLogger(LoggerPrefixWails),
		i18n:     utils.NewConsoleLogger(LoggerPrefixI18n),
		tray:     utils.NewConsoleLogger(LoggerPrefixTray),
		web:      utils.NewConsoleLogger(LoggerPrefixWeb),
		service:  utils.NewConsoleLogger(LoggerPrefixService),
	}
}

func NewFileLogger(logPath string) *Logger {
	logFile, err := os.OpenFile(
		logPath,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		utils.Utils().PanicLogger().Fatalf("failed to open log file: %+v\n", err)
	}
	return &Logger{
		database: utils.NewGormFileLogger(LoggerPrefixDatabase, logFile),
		wails:    utils.NewWailsFileLogger(LoggerPrefixWails, logFile),
		i18n:     utils.NewFileLogger(LoggerPrefixI18n, logFile),
		tray:     utils.NewFileLogger(LoggerPrefixTray, logFile),
		web:      utils.NewFileLogger(LoggerPrefixWeb, logFile),
		service:  utils.NewFileLogger(LoggerPrefixService, logFile),
	}
}

func (l *Logger) Wails() wailsLogger.Logger {
	return l.wails
}

func (l *Logger) Tray() *utils.Logger {
	return l.tray
}

func (l *Logger) Web() *utils.Logger {
	return l.web
}

func (l *Logger) Service() *utils.Logger {
	return l.service
}
