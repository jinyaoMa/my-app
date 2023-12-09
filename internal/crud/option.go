package crud

import (
	"my-app/internal/entity"
	"my-app/pkg/db"

	"gorm.io/gorm"
)

type Option struct {
	db.ICRUD[*entity.Option]
	session *gorm.DB
}

func NewOption(session *gorm.DB) (option *Option, iOption IOption) {
	_, crud := db.NewCRUD[*entity.Option](session)
	option = &Option{
		ICRUD:   crud,
		session: session,
	}
	return option, option
}
