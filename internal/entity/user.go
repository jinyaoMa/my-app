package entity

import (
	"my-app/pkg/db"

	"gorm.io/gorm"
)

type User struct {
	db.Entity[*User]
	Account  string `gorm:"unique; index; not null"` // encrypted
	Password string `gorm:"not null"`                // hashed
	Active   bool   `gorm:"default:true"`

	/* has many */
	Files []*File `gorm:""`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	if err = user.Entity.BeforeCreate(tx); err != nil {
		return
	}

	tx.Create(&UserPassword{
		Account:  user.Account,
		Password: user.Password,
	})
	return
}

func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if err = user.Entity.BeforeUpdate(tx); err != nil {
		return
	}

	tx.Create(&UserPassword{
		Account:  user.Account,
		Password: user.Password,
	})
	return
}
