package log

import "log"

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
	*log.Logger
}

func (l *Log) SetOutput(out ITreeWriter) {
	l.Logger.SetOutput(out)
}

func New(out ITreeWriter, prefix string, flag int) *Log {
	return &Log{
		Logger: log.New(out, prefix, flag),
	}
}

func Default() *Log {
	return New(NewConsoleLogWriter(), "[LOG] ", Ldate|Ltime|Lmicroseconds|Lshortfile)
}
