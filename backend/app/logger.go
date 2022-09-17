package app

import (
	"my-app/backend/pkg/log"
	"os"
	"time"

	"gorm.io/gorm/logger"
)

const (
	LogPrefixModel = "MODEL"
	LogPrefixApp   = "APP__"
	LogPrefixWeb   = "WEB__"
	LogPrefixTray  = "TRAY_"
	LogPrefixWails = "WAILS"
)

type Logger struct {
	Model logger.Interface
	App   *log.Logger
	Web   *log.Logger
	Tray  *log.Logger
	Wails *log.Logger
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
		App:   log.ConsoleLogger(LogPrefixApp),
		Web:   log.ConsoleLogger(LogPrefixWeb),
		Tray:  log.ConsoleLogger(LogPrefixTray),
		Wails: log.ConsoleLogger(LogPrefixWails),
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
		App:   log.FileLogger(LogPrefixApp, file),
		Web:   log.FileLogger(LogPrefixWeb, file),
		Tray:  log.FileLogger(LogPrefixTray, file),
		Wails: log.FileLogger(LogPrefixWails, file),
	}
}
