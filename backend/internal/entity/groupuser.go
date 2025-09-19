package entity

import (
	"majinyao.cn/my-app/backend/pkg/db"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
)

type GroupUser struct {
	db.Entity

	GroupId datatype.Id `gorm:"comment:Group Id;"`
	Group   Group

	UserId datatype.Id `gorm:"comment:User Id;"`
	User   User
}
