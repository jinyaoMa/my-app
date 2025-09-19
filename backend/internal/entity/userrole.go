package entity

import (
	"majinyao.cn/my-app/backend/pkg/db"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
)

type UserRole struct {
	db.Entity

	UserId datatype.Id `gorm:"comment:User Id;"`
	User   User

	RoleId datatype.Id `gorm:"comment:Role Id;"`
	Role   Role
}
