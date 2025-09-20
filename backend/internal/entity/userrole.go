package entity

import (
	"majinyao.cn/my-app/backend/pkg/db/datatype"
	"majinyao.cn/my-app/backend/pkg/db/model"
)

type UserRole struct {
	model.Model

	UserId datatype.Id `gorm:"comment:User Id;"`
	User   User

	RoleId datatype.Id `gorm:"comment:Role Id;"`
	Role   Role
}
