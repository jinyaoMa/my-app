package logger

import (
	"log"
	"my-app/backend/pkg/logger/interfaces"
)

type Logger struct {
	*log.Logger
	*Options
}

func NewLogger(opts *Options) interfaces.ILogger {
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
