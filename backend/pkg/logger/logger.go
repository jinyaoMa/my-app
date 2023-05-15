package logger

import (
	"io"
	"log"
	"my-app/backend/pkg/logger/interfaces"
	"my-app/backend/pkg/logger/options"
)

type Logger struct {
	log.Logger
	options *options.OLogger
}

// Writer implements interfaces.ILogger
func (l *Logger) Writer() io.Writer {
	return l.Logger.Writer()
}

func NewLogger(opts *options.OLogger) interfaces.ILogger {
	opts = options.NewOLogger(opts)

	logger := log.Default()
	logger.SetOutput(opts.Writer)
	logger.SetPrefix(opts.PrefixTemplate(opts.Tag))
	logger.SetFlags(opts.Flags)

	return &Logger{
		Logger:  *logger,
		options: opts,
	}
}
