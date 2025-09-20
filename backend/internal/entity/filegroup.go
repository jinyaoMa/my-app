package entity

import (
	"majinyao.cn/my-app/backend/pkg/db/datatype"
	"majinyao.cn/my-app/backend/pkg/db/model"
)

type FileGroup struct {
	model.Model

	FileId datatype.Id `gorm:"comment:File Id;"`
	File   File

	GroupId datatype.Id `gorm:"comment:Group Id;"`
	Group   Group

	NoCreate bool `gorm:"comment:No Create;"`
	NoRead   bool `gorm:"comment:No Read;"`
	NoUpdate bool `gorm:"comment:No Update;"`
	NoDelete bool `gorm:"comment:No Delete;"`
}
