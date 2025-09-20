package entity

import (
	"majinyao.cn/my-app/backend/pkg/db/datatype"
	"majinyao.cn/my-app/backend/pkg/db/model"
)

type GroupRole struct {
	model.Model
	model.Reserved

	GroupId datatype.Id `gorm:"comment:Group Id;"`
	Group   Group

	RoleId datatype.Id `gorm:"comment:Role Id;"`
	Role   Role
}
