package utils

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"
)

type Logger struct {
	*log.Logger
}

func NewConsoleLogger(prefix string) *Logger {
	l := &Logger{}
	l.Logger = log.New(os.Stdout, l.modifyPrefix(prefix), log.Ldate|log.Ltime|log.Lshortfile)
	return l
}

func NewFileLogger(prefix string, f *os.File) *Logger {
	l := &Logger{}
	l.Logger = log.New(f, l.modifyPrefix(prefix), log.Ldate|log.Ltime|log.Llongfile)
	return l
}

func NewGormConsoleLogger(prefix string) logger.Interface {
	return logger.New(
		NewConsoleLogger(prefix),
		logger.Config{
			SlowThreshold:             time.Second,
			Colorful:                  true,
			IgnoreRecordNotFoundError: false,
			LogLevel:                  logger.Info,
		},
	)
}

func NewGormFileLogger(prefix string, file *os.File) logger.Interface {
	return logger.New(
		NewFileLogger(prefix, file),
		logger.Config{
			SlowThreshold:             time.Second,
			Colorful:                  false,
			IgnoreRecordNotFoundError: false,
			LogLevel:                  logger.Error,
		},
	)
}

func (l *Logger) modifyPrefix(prefix string) string {
	if prefix == "" {
		return prefix
	}
	return "[" + prefix + "] "
}
