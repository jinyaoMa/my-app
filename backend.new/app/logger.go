package app

import (
	"my-app/backend.new/utils"
	"os"

	wailsLogger "github.com/wailsapp/wails/v2/pkg/logger"
	gormLogger "gorm.io/gorm/logger"
)

const (
	LogPrefixDatabase = "DBS"
	LogPrefixWails    = "WLS"
	LogPrefixTray     = "TRY"
	LogPrefixWeb      = "WEB"
	LogPrefixService  = "SEV"
)

type Logger struct {
	database gormLogger.Interface
	wails    wailsLogger.Logger
	tray     *utils.Logger
	web      *utils.Logger
	service  *utils.Logger
}

func (l *Logger) Database() gormLogger.Interface {
	return l.database
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

func NewConsoleLogger() *Logger {
	return &Logger{
		database: utils.NewGormConsoleLogger(LogPrefixDatabase),
		wails:    utils.NewWailsConsoleLogger(LogPrefixWails),
		tray:     utils.NewConsoleLogger(LogPrefixTray),
		web:      utils.NewConsoleLogger(LogPrefixWeb),
		service:  utils.NewConsoleLogger(LogPrefixService),
	}
}

func NewFileLogger(logPath string) *Logger {
	logFile, err := os.OpenFile(
		logPath,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		panic("failed to open log file")
	}
	return &Logger{
		database: utils.NewGormFileLogger(LogPrefixDatabase, logFile),
		wails:    utils.NewWailsFileLogger(LogPrefixWails, logFile),
		tray:     utils.NewFileLogger(LogPrefixTray, logFile),
		web:      utils.NewFileLogger(LogPrefixWeb, logFile),
		service:  utils.NewFileLogger(LogPrefixService, logFile),
	}
}
