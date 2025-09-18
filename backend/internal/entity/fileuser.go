package entity

import "majinyao.cn/my-app/backend/pkg/db"

type FileUser struct {
	db.Entity

	FileId int64 `gorm:"comment:File Id;"`
	File   File

	UserId int64 `gorm:"comment:User Id;"`
	User   User

	NoCreate bool `gorm:"comment:No Create;"`
	NoRead   bool `gorm:"comment:No Read;"`
	NoUpdate bool `gorm:"comment:No Update;"`
	NoDelete bool `gorm:"comment:No Delete;"`
	IsAvatar bool `gorm:"index;comment:Is Used as Avatar or Not;"`
}
