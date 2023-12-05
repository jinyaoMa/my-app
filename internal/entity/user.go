package entity

import (
	"my-app/pkg/db"
)

type User struct {
	db.Entity[*User]
	Account      string `gorm:"size:64; unique; index; not null"`
	Password     string `gorm:"-:all"`
	PasswordHash string `gorm:"size:64; not null"`
	Active       bool   `gorm:"default:true"`

	/* has many */
	Files []*File `gorm:""`
}
