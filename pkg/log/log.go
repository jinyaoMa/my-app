package log

import (
	"log"
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

type Log struct {
	*Config
	*log.Logger
}

func (l *Log) SetOutput(out ITreeWriter) {
	l.Logger.SetOutput(out)
}

func New(cfg *Config) *Log {
	cfg = NewConfig(cfg)
	return &Log{
		Config: cfg,
		Logger: log.New(cfg.Out, cfg.Prefix, cfg.Flag),
	}
}

func Default() *Log {
	return New(DefaultConfig())
}
