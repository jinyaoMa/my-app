package entity

import "my-app/backend/pkg/db"

type FileExtension struct {
	db.Entity

	/* internal fields */
	Name    string `gorm:"size:256"`
	DotName string `gorm:"size:64"`

	/* relational fields */
	Files          []*File `gorm:""`
	FileCategoryID int64   `gorm:""`
}
