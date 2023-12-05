package entity

import (
	"my-app/pkg/db"

	"gorm.io/datatypes"
)

type FileExtension struct {
	db.Entity[*FileExtension]
	Name    datatypes.JSON `gorm:""` // store serialized json string, eg. {"zh-CN":"MP4","en-CA":"MP4"}
	DotName string         `gorm:""`

	/* belongs to */
	FileCategoryID int64        `gorm:""`
	FileCategory   FileCategory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	/* has many */
	Files []*File `gorm:""`
}
