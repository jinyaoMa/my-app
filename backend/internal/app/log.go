package app

import (
	"my-app/backend/configs"
	"my-app/backend/internal/crud"
	"my-app/backend/internal/entity"
	"my-app/backend/internal/interfaces"
	"my-app/backend/pkg/db"
	"my-app/backend/pkg/log"
)

func initLog(cfg *configs.Configs, dbs *db.DB) (l *log.Log, err error) {
	l = log.New(&log.Config{
		Out:    newDbLogWriter(dbs, log.NewConsoleLogWriter()),
		Prefix: "[APP] ",
		Flag:   log.DefaultFlag,
	})
	return
}

type DbLogWriter struct {
	*log.LogWriter
	crudLog interfaces.ICRUDLog
}

// Write implements io.Writer.
func (w *DbLogWriter) Write(p []byte) (n int, err error) {
	bound := len(p)
	if bound > 4096 {
		bound = 4096
	}
	_, err = w.crudLog.Save(&entity.Log{
		Message: string(p[:bound]),
	})
	return len(p), err
}

func newDbLogWriter(dbs *db.DB, children ...log.ITreeWriter) log.ITreeWriter {
	return &DbLogWriter{
		LogWriter: log.NewLogWriter(children...),
		crudLog:   crud.NewLog(dbs),
	}
}
