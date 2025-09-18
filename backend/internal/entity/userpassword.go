package entity

import (
	"majinyao.cn/my-app/backend/pkg/db"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
)

type UserPassword struct {
	db.Entity
	Password datatype.Encrypted `gorm:"comment:Historical Password;"`

	UserId int64 `gorm:"comment:User Id;"`
	User   User
}
