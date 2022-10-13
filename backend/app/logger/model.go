package logger

import (
	"os"
	"time"

	"gorm.io/gorm/logger"
)

func modelConsoleLogger(prefix string) logger.Interface {
	return logger.New(
		consoleLogger(prefix),
		logger.Config{
			SlowThreshold:             time.Second,
			Colorful:                  true,
			IgnoreRecordNotFoundError: false,
			LogLevel:                  logger.Info,
		},
	)
}

func modelFileLogger(prefix string, file *os.File) logger.Interface {
	return logger.New(
		fileLogger(prefix, file),
		logger.Config{
			SlowThreshold:             time.Second,
			Colorful:                  false,
			IgnoreRecordNotFoundError: false,
			LogLevel:                  logger.Error,
		},
	)
}
