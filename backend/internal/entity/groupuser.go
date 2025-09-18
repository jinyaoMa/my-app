package entity

import "majinyao.cn/my-app/backend/pkg/db"

type GroupUser struct {
	db.Entity

	GroupId int64 `gorm:"comment:Group Id;"`
	Group   Group

	UserId int64 `gorm:"comment:User Id;"`
	User   User
}
