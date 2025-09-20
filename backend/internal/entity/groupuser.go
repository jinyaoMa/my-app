package entity

import (
	"majinyao.cn/my-app/backend/pkg/db/datatype"
	"majinyao.cn/my-app/backend/pkg/db/model"
)

type GroupUser struct {
	model.Model

	GroupId datatype.Id `gorm:"comment:Group Id;"`
	Group   Group

	UserId datatype.Id `gorm:"comment:User Id;"`
	User   User
}
