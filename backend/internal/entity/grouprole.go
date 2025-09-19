package entity

import (
	"majinyao.cn/my-app/backend/pkg/db"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
)

type GroupRole struct {
	db.Entity
	db.EntityReserved

	GroupId datatype.Id `gorm:"comment:Group Id;"`
	Group   Group

	RoleId datatype.Id `gorm:"comment:Role Id;"`
	Role   Role
}
