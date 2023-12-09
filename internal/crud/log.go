package crud

import (
	"my-app/internal/entity"
	"my-app/pkg/db"

	"gorm.io/gorm"
)

type Log struct {
	db.ICRUD[*entity.Log]
	session *gorm.DB
}

func NewLog(session *gorm.DB) (log *Log, iLog ILog) {
	_, crud := db.NewCRUD[*entity.Log](session)
	log = &Log{
		ICRUD:   crud,
		session: session,
	}
	return log, log
}
