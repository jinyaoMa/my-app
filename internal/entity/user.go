package entity

import (
	"encoding/hex"
	"my-app/pkg/db"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	db.Entity[*User]
	Account  string `gorm:"unique; index; not null; <-:create"` // encrypted
	Password string `gorm:"not null"`                           // hashed
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

	if tx.Statement.Changed("Password") {
		tx.Create(&UserPassword{
			Account:  user.Account,
			Password: user.Password,
		})
	}
	return
}

func (user *User) HashPassword(cost int) (err error) {
	bPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), cost)
	if err != nil {
		return err
	}
	user.Password = hex.EncodeToString(bPassword)
	return
}

func (user *User) VerifyPassword(plain string) (ok bool) {
	bPassword, err := hex.DecodeString(user.Password)
	if err != nil {
		return false
	}
	return bcrypt.CompareHashAndPassword(bPassword, []byte(plain)) == nil
}
