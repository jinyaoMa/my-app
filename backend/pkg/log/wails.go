package log

import (
	"os"
)

type WailsLogger struct {
	log *Logger
}

func NewWailsConsoleLogger(prefix string) WailsLogger {
	return WailsLogger{
		log: ConsoleLogger(prefix),
	}
}

func NewWailsFileLogger(prefix string, file *os.File) WailsLogger {
	return WailsLogger{
		log: FileLogger(prefix, file),
	}
}

func (wl *WailsLogger) Print(message string) {
	wl.log.Println(message)
}

func (wl *WailsLogger) Trace(message string) {
	wl.log.Println("TRA | " + message)
}

func (wl *WailsLogger) Debug(message string) {
	wl.log.Println("DEB | " + message)
}

func (wl *WailsLogger) Info(message string) {
	wl.log.Println("INF | " + message)
}

func (wl *WailsLogger) Warning(message string) {
	wl.log.Println("WAR | " + message)
}

func (wl *WailsLogger) Error(message string) {
	wl.log.Println("ERR | " + message)
}

func (wl *WailsLogger) Fatal(message string) {
	wl.log.Fatalln("FAT | " + message)
}
