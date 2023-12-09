package crud

import (
	"my-app/internal/entity"
	"my-app/pkg/db"

	"gorm.io/gorm"
)

type UserPassword struct {
	db.ICRUD[*entity.UserPassword]
	session *gorm.DB
}

func NewUserPassword(session *gorm.DB) (userPassword *UserPassword, iUserPassword IUserPassword) {
	_, crud := db.NewCRUD[*entity.UserPassword](session)
	userPassword = &UserPassword{
		ICRUD:   crud,
		session: session,
	}
	return userPassword, userPassword
}
