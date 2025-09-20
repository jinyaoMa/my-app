package entity

import "majinyao.cn/my-app/backend/pkg/db/model"

type FileCategory struct {
	model.Model
	model.Reserved
	Code string `gorm:"index;not null;size:254;comment:File Category Code;"`
	Name string `gorm:"index;size:254;comment:File Category Name;"`

	FileExtensions []FileExtension
}
