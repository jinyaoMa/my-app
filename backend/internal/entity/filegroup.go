package entity

import (
	"majinyao.cn/my-app/backend/pkg/db"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
)

type FileGroup struct {
	db.Entity

	FileId datatype.Id `gorm:"comment:File Id;"`
	File   File

	GroupId datatype.Id `gorm:"comment:Group Id;"`
	Group   Group

	NoCreate bool `gorm:"comment:No Create;"`
	NoRead   bool `gorm:"comment:No Read;"`
	NoUpdate bool `gorm:"comment:No Update;"`
	NoDelete bool `gorm:"comment:No Delete;"`
}
