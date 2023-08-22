package log

import (
	"log"
	"time"

	gormLogger "gorm.io/gorm/logger"
)

// copy from standard library "log"
const (
	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)

// copy from "gorm.io/gorm/logger"
const (
	// Silent silent log level
	Silent LogLevel = iota + 1
	// Error error log level
	Error
	// Warn warn log level
	Warn
	// Info info log level
	Info
)

type LogLevel int

type GormConfig struct {
	SlowThreshold             time.Duration
	Colorful                  bool
	IgnoreRecordNotFoundError bool
	ParameterizedQueries      bool
	LogLevel                  LogLevel
}

type Log struct {
	*log.Logger
}

func (l *Log) SetOutput(out ITreeWriter) {
	l.Logger.SetOutput(out)
}

func New(opt *Option) *Log {
	opt = NewOption(opt)
	return &Log{
		Logger: log.New(opt.Out, opt.Prefix, opt.Flag),
	}
}

func Gorm(opt *Option, config GormConfig) gormLogger.Interface {
	opt = NewOption(opt)
	return gormLogger.New(New(opt), gormLogger.Config{
		SlowThreshold:             config.SlowThreshold,
		Colorful:                  config.Colorful,
		IgnoreRecordNotFoundError: config.IgnoreRecordNotFoundError,
		ParameterizedQueries:      config.ParameterizedQueries,
		LogLevel:                  gormLogger.LogLevel(config.LogLevel),
	})
}

func Default() *Log {
	return New(DefaultOption())
}
