package entity

import (
	"majinyao.cn/my-app/backend/pkg/db"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
)

type FileExtension struct {
	db.Entity
	db.EntityReserved
	Ext  string `gorm:"index;not null;size:254;check:ext LIKE '.%';comment:File Extension (.ext);"`
	Name string `gorm:"index;size:254;comment:File Extension Display Name;"`
	Mime string `gorm:"index;size:254;comment:File Extension Mime;"`

	FileCategoryId *datatype.Id `gorm:"comment:File Category Id;"`
	FileCategory   *FileCategory

	Files []File
}
