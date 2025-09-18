package entity

import "majinyao.cn/my-app/backend/pkg/db"

type FileGroup struct {
	db.Entity

	FileId int64 `gorm:"comment:File Id;"`
	File   File

	GroupId int64 `gorm:"comment:Group Id;"`
	Group   Group

	NoCreate bool `gorm:"comment:No Create;"`
	NoRead   bool `gorm:"comment:No Read;"`
	NoUpdate bool `gorm:"comment:No Update;"`
	NoDelete bool `gorm:"comment:No Delete;"`
}
