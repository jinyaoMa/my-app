package entity

import (
	"my-app/pkg/db"

	"gorm.io/datatypes"
)

type FileCategory struct {
	db.Entity[*FileCategory]
	Name datatypes.JSON `gorm:""` // store serialized json string, eg. {"zh-CN":"分类","en-CA":"Category"}

	/* has many */
	FileExtensions []*FileExtension `gorm:""`
}
