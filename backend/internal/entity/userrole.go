package entity

import "majinyao.cn/my-app/backend/pkg/db"

type UserRole struct {
	db.Entity

	UserId int64 `gorm:"comment:User Id;"`
	User   User

	RoleId int64 `gorm:"comment:Role Id;"`
	Role   Role
}
