package entity

import (
	"crypto/sha256"
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	Entity
	Account      string `gorm:"size:64; unique; ndex"`
	Password     string `gorm:"-:all"`
	PasswordHash string `gorm:"size:64"`
	IsFrozen     bool   `gorm:""`
	OldPasswords []*UserPassword
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if err = u.Entity.EntityBase.BeforeCreate(tx); err != nil {
		return
	}

	if u != nil {
		u.hashPassword()
	}
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if err = u.Entity.EntityBase.BeforeUpdate(tx); err != nil {
		return
	}

	if u != nil {
		u.hashPassword()
	}
	return
}

func (u *User) hashPassword() {
	if u.Password != "" {
		passwordSum := sha256.Sum256([]byte(u.Password))
		u.PasswordHash = fmt.Sprintf("%x", passwordSum)
		u.OldPasswords = append(u.OldPasswords, &UserPassword{
			PasswordHash: u.PasswordHash,
		})
	}
}
