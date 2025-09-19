package entity

import (
	"majinyao.cn/my-app/backend/pkg/db"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
)

type RolePermission struct {
	db.Entity
	db.EntityReserved

	RoleId datatype.Id `gorm:"comment:Role Id;"`
	Role   Role

	PermissionId datatype.Id `gorm:"comment:Permission Id;"`
	Permission   Permission
}
