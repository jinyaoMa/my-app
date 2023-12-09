package crud

import (
	"my-app/internal/entity"
	"my-app/pkg/db"

	"gorm.io/gorm"
)

type User struct {
	db.ICRUD[*entity.User]
	session *gorm.DB
}

func NewUser(session *gorm.DB) (user *User, iUser IUser) {
	_, crud := db.NewCRUD[*entity.User](session)
	user = &User{
		ICRUD:   crud,
		session: session,
	}
	return user, user
}
