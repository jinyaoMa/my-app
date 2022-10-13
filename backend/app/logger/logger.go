package logger

import (
	"log"
	"os"

	wailsLogger "github.com/wailsapp/wails/v2/pkg/logger"
	gormLogger "gorm.io/gorm/logger"
)

const (
	LogPrefixWails   = "WLS"
	LogPrefixTray    = "TRY"
	LogPrefixWeb     = "WEB"
	LogPrefixService = "SEV"
	LogPrefixModel   = "MDL"
)

type Logger struct {
	wails   wailsLogger.Logger
	tray    *log.Logger
	web     *log.Logger
	service *log.Logger
	model   gormLogger.Interface
}

func NewConsoleLogger() *Logger {
	return &Logger{
		wails:   wailsConsoleLogger(LogPrefixWails),
		tray:    consoleLogger(LogPrefixTray),
		web:     consoleLogger(LogPrefixWeb),
		service: consoleLogger(LogPrefixService),
		model:   modelConsoleLogger(LogPrefixModel),
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
		wails:   wailsFileLogger(LogPrefixWails, logFile),
		tray:    fileLogger(LogPrefixTray, logFile),
		web:     fileLogger(LogPrefixWeb, logFile),
		service: fileLogger(LogPrefixService, logFile),
		model:   modelFileLogger(LogPrefixModel, logFile),
	}
}

func (l *Logger) Wails() wailsLogger.Logger {
	return l.wails
}

func (l *Logger) Tray() *log.Logger {
	return l.tray
}

func (l *Logger) Web() *log.Logger {
	return l.web
}

func (l *Logger) Service() *log.Logger {
	return l.service
}

func (l *Logger) Model() gormLogger.Interface {
	return l.model
}

func consoleLogger(prefix string) *log.Logger {
	return log.New(os.Stdout, modifyPrefix(prefix), log.Ldate|log.Ltime|log.Lshortfile)
}

func fileLogger(prefix string, f *os.File) *log.Logger {
	return log.New(f, modifyPrefix(prefix), log.Ldate|log.Ltime|log.Llongfile)
}

func modifyPrefix(prefix string) string {
	return "[" + prefix + "] "
}
