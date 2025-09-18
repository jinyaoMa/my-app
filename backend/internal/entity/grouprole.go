package entity

import "majinyao.cn/my-app/backend/pkg/db"

type GroupRole struct {
	db.Entity
	db.EntityReserved

	GroupId int64 `gorm:"comment:Group Id;"`
	Group   Group

	RoleId int64 `gorm:"comment:Role Id;"`
	Role   Role
}
