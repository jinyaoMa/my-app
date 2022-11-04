package model

import (
	"my-app/backend.new/utils"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Account  string `gorm:"unique,index"` // User account
	Password string ``                    // User password
	Keyrings Keyrings

	PasswordHashed bool `gorm:"-:all"` // indicate if password is hashed
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if !u.PasswordHashed {
		u.Password = utils.Utils().SHA1(u.Password)
		u.PasswordHashed = true
	}
	return
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if !u.PasswordHashed {
		u.Password = utils.Utils().SHA1(u.Password)
		u.PasswordHashed = true
	}
	return
}

func (u *User) AfterFind(tx *gorm.DB) (err error) {
	u.PasswordHashed = true
	return
}

func (u *User) FindByAccount(db *gorm.DB) (ok bool) {
	tx := db.Where(User{
		Account: u.Account,
	}).Find(u)
	return tx.RowsAffected > 0
}

func (u *User) Create(db *gorm.DB) (ok bool) {
	tx := db.Create(u)
	return tx.RowsAffected == 1
}

func (u *User) Save(db *gorm.DB) (ok bool) {
	tx := db.Save(u)
	return tx.RowsAffected == 1
}
