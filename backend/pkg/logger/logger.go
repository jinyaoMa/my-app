package logger

import (
	"log"
	"my-app/backend/pkg/logger/interfaces"
	"my-app/backend/pkg/logger/options"
)

type Logger struct {
	*log.Logger
	*options.OLogger
}

func NewLogger(opts *options.OLogger) interfaces.ILogger {
	opts = options.NewOLogger(opts)

	logger := log.Default()
	logger.SetOutput(opts.Writer)
	logger.SetPrefix(opts.PrefixTemplate(opts.Tag))
	logger.SetFlags(opts.Flags)

	return &Logger{
		Logger:  logger,
		OLogger: opts,
	}
}
