package entity

import "my-app/backend/pkg/db"

type FileCategory struct {
	db.Entity

	/* internal fields */
	Name string `gorm:"size:512"`

	/* relational fields */
	FileExtensions []*FileExtension `gorm:""`
}
