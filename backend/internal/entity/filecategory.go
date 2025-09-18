package entity

import "majinyao.cn/my-app/backend/pkg/db"

type FileCategory struct {
	db.Entity
	db.EntityReserved
	Code string `gorm:"index;not null;size:254;comment:File Category Code;"`
	Name string `gorm:"index;size:254;comment:File Category Name;"`

	FileExtensions []FileExtension
}
