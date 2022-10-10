package logger

import (
	"log"
	"os"

	"github.com/wailsapp/wails/v2/pkg/logger"
)

type WailsLogger struct {
	log *log.Logger
}

func wailsConsoleLogger(prefix string) logger.Logger {
	return &WailsLogger{
		log: consoleLogger(prefix),
	}
}

func wailsFileLogger(prefix string, file *os.File) logger.Logger {
	return &WailsLogger{
		log: fileLogger(prefix, file),
	}
}

// Print implements logger.Logger
func (w *WailsLogger) Print(message string) {
	w.log.Println(message)
}

// Trace implements logger.Logger
func (w *WailsLogger) Trace(message string) {
	w.log.Println("TRA | " + message)
}

// Debug implements logger.Logger
func (w *WailsLogger) Debug(message string) {
	w.log.Println("DEB | " + message)
}

// Info implements logger.Logger
func (w *WailsLogger) Info(message string) {
	w.log.Println("INF | " + message)
}

// Warning implements logger.Logger
func (w *WailsLogger) Warning(message string) {
	w.log.Println("WAR | " + message)
}

// Error implements logger.Logger
func (w *WailsLogger) Error(message string) {
	w.log.Println("ERR | " + message)
}

// Fatal implements logger.Logger
func (w *WailsLogger) Fatal(message string) {
	w.log.Fatalln("FAT | " + message)
}
