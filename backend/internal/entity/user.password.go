package entity

import "my-app/backend/pkg/db"

type UserPassword struct {
	db.Entity

	/* internal fields */
	PasswordHash string `gorm:"size:64; notnull; <-:create"`

	/* relational fields */
	UserID int64 `gorm:""`
}
