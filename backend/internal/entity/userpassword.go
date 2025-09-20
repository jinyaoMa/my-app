package entity

import (
	"majinyao.cn/my-app/backend/pkg/db/datatype"
	"majinyao.cn/my-app/backend/pkg/db/model"
)

type UserPassword struct {
	model.Model
	Password datatype.Encrypted `gorm:"comment:Historical Password;"`

	UserId datatype.Id `gorm:"comment:User Id;"`
	User   User
}
