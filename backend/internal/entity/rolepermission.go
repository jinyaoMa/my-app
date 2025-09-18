package entity

import "majinyao.cn/my-app/backend/pkg/db"

type RolePermission struct {
	db.Entity
	db.EntityReserved

	RoleId int64 `gorm:"comment:Role Id;"`
	Role   Role

	PermissionId int64 `gorm:"comment:Permission Id;"`
	Permission   Permission
}
