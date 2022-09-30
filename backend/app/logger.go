package app

import (
	"my-app/backend/pkg/log"
	"os"
	"time"

	wailsLogger "github.com/wailsapp/wails/v2/pkg/logger"
	"gorm.io/gorm/logger"
)

const (
	LogPrefixModel   = "MDL"
	LogPrefixWails   = "WLS"
	LogPrefixApp     = "APP"
	LogPrefixTray    = "TRY"
	LogPrefixWeb     = "WEB"
	LogPrefixService = "SEV"
)

type Logger struct {
	Model   logger.Interface
	Wails   wailsLogger.Logger
	App     *log.Logger
	Tray    *log.Logger
	Web     *log.Logger
	Service *log.Logger
}

func LoadConsoleLogger() *Logger {
	return &Logger{
		Model: logger.New(
			log.ConsoleLogger(LogPrefixModel),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: false,
				Colorful:                  true,
			},
		),
		Wails:   log.NewWailsConsoleLogger(LogPrefixWails),
		App:     log.ConsoleLogger(LogPrefixApp),
		Tray:    log.ConsoleLogger(LogPrefixTray),
		Web:     log.ConsoleLogger(LogPrefixWeb),
		Service: log.ConsoleLogger(LogPrefixService),
	}
}

func LoadFileLogger(file *os.File) *Logger {
	return &Logger{
		Model: logger.New(
			log.FileLogger(LogPrefixModel, file),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Error,
				IgnoreRecordNotFoundError: false,
				Colorful:                  false,
			},
		),
		Wails:   log.NewWailsFileLogger(LogPrefixWails, file),
		App:     log.FileLogger(LogPrefixApp, file),
		Tray:    log.FileLogger(LogPrefixTray, file),
		Web:     log.FileLogger(LogPrefixWeb, file),
		Service: log.FileLogger(LogPrefixService, file),
	}
}
