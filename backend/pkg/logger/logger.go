package logger

import (
	"io"
	"log"
)

type Logger struct {
	*log.Logger
	options *Option
}

// Tag implements Interface.
func (l *Logger) Tag() string {
	return l.options.Tag
}

// Writer implements Interface
func (l *Logger) Writer() io.Writer {
	return l.Logger.Writer()
}

func New(opts *Option) Interface {
	opts = NewOption(opts)

	logger := log.Default()
	logger.SetOutput(opts.Writer)
	logger.SetPrefix(opts.PrefixTemplate(opts.Tag))
	logger.SetFlags(opts.Flags)

	return &Logger{
		Logger:  logger,
		options: opts,
	}
}
