package entity

import (
	"majinyao.cn/my-app/backend/pkg/db"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
)

type FileUser struct {
	db.Entity

	FileId datatype.Id `gorm:"comment:File Id;"`
	File   File

	UserId datatype.Id `gorm:"comment:User Id;"`
	User   User

	NoCreate bool `gorm:"comment:No Create;"`
	NoRead   bool `gorm:"comment:No Read;"`
	NoUpdate bool `gorm:"comment:No Update;"`
	NoDelete bool `gorm:"comment:No Delete;"`
	IsAvatar bool `gorm:"index;comment:Is Used as Avatar or Not;"`
}
