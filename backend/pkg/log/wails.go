package log

import (
	"os"
)

type WailsLogger struct {
	*Logger
}

func NewWailsConsoleLogger(prefix string) *WailsLogger {
	return &WailsLogger{
		Logger: ConsoleLogger(prefix),
	}
}

func NewWailsFileLogger(prefix string, file *os.File) *WailsLogger {
	return &WailsLogger{
		Logger: FileLogger(prefix, file),
	}
}

func (wl *WailsLogger) Print(message string) {
	wl.Logger.Println(message)
}

func (wl *WailsLogger) Trace(message string) {
	wl.Logger.Println("TRA | " + message)
}

func (wl *WailsLogger) Debug(message string) {
	wl.Logger.Println("DEB | " + message)
}

func (wl *WailsLogger) Info(message string) {
	wl.Logger.Println("INF | " + message)
}

func (wl *WailsLogger) Warning(message string) {
	wl.Logger.Println("WAR | " + message)
}

func (wl *WailsLogger) Error(message string) {
	wl.Logger.Println("ERR | " + message)
}

func (wl *WailsLogger) Fatal(message string) {
	wl.Logger.Fatalln("FAT | " + message)
}
