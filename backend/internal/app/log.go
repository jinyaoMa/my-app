package app

import (
	"io"
	"my-app/backend/configs"
	"my-app/backend/internal/crud"
	"my-app/backend/internal/interfaces"
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/database/entity"
	"my-app/backend/pkg/logger"
)

func initLog(cfg *configs.Configs, db *database.Database) (log logger.Interface, err error) {
	tag := "APP"
	log = logger.New(&logger.Option{
		Writer: newDbLogWriter(db),
		Tag:    tag,
	})
	return
}

type DbLogWriter struct {
	logService interfaces.ILogService
	tag        string
}

// Write implements io.Writer.
func (w *DbLogWriter) Write(p []byte) (n int, err error) {
	_, err = w.logService.Save(&entity.Log{
		Message: string(p),
	})
	return len(p), err
}

func newDbLogWriter(db *database.Database) io.Writer {
	return &DbLogWriter{
		logService: crud.NewLogService(db),
	}
}
