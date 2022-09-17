package log

import (
	"log"
	"os"
)

type Logger struct {
	*log.Logger
}

func ConsoleLogger(prefix string) *Logger {
	label := "[" + prefix + "] "
	return &Logger{
		Logger: log.New(os.Stdout, label, log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func FileLogger(prefix string, file *os.File) *Logger {
	label := "[" + prefix + "] "
	return &Logger{
		Logger: log.New(file, label, log.Ldate|log.Ltime|log.Lshortfile),
	}
}
