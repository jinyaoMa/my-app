package logger

import (
	"io"
	"log"
)

type ILogger interface {
	Fatal(v ...any)
	Fatalf(format string, v ...any)
	Fatalln(v ...any)
	Flags() int
	Output(calldepth int, s string) error
	Panic(v ...any)
	Panicf(format string, v ...any)
	Panicln(v ...any)
	Prefix() string
	Print(v ...any)
	Printf(format string, v ...any)
	Println(v ...any)
	SetFlags(flag int)
	SetOutput(w io.Writer)
	SetPrefix(prefix string)
	Writer() io.Writer
}

type Logger struct {
	*log.Logger
	*Options
}

// Flags implements ILogger
func (l *Logger) Flags() int {
	return l.Options.Flags
}

// Writer implements ILogger
func (l *Logger) Writer() io.Writer {
	return l.Options.Writer
}

func NewLogger(opts *Options) ILogger {
	opts = NewOptions(opts)

	logger := log.Default()
	logger.SetOutput(opts.Writer)
	logger.SetPrefix(opts.PrefixTemplate(opts.Tag))
	logger.SetFlags(opts.Flags)

	return &Logger{
		Logger:  logger,
		Options: opts,
	}
}
