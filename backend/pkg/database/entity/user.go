package entity

import (
	"crypto/sha256"
	"fmt"
)

type User struct {
	Entity       `xorm:"extends"`
	Account      string          `xorm:"varchar(64) notnull unique"`
	Password     string          `xorm:"-"`
	PasswordHash string          `xorm:"varchar(64) notnull"`
	IsFrozen     bool            `xorm:"notnull"`
	OldPasswords []*UserPassword `xorm:"extends"`
}

func (u *User) BeforeInsert() {
	u.Entity.EntityBase.BeforeInsert()
	u.hashPassword()
}

func (u *User) BeforeUpdate() {
	u.hashPassword()
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
