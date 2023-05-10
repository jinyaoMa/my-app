package logger

import "log"

type ILogger interface {
}

type Logger struct {
	*log.Logger
	*Options
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
