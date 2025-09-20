package entity

import (
	"majinyao.cn/my-app/backend/pkg/db/datatype"
	"majinyao.cn/my-app/backend/pkg/db/model"
)

type RolePermission struct {
	model.Model
	model.Reserved

	RoleId datatype.Id `gorm:"comment:Role Id;"`
	Role   Role

	PermissionId datatype.Id `gorm:"comment:Permission Id;"`
	Permission   Permission
}
