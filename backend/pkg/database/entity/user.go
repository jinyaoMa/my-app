package entity

import (
	"crypto/sha256"
	"fmt"
)

type User struct {
	EntitySafe
	Account      string `xorm:"varchar(64) notnull unique"`
	Password     string `xorm:"-"`
	PasswordHash string `xorm:"varchar(64) notnull"`
}

func (u *User) BeforeInsert() {
	u.EntitySafe.BeforeInsert()
	u.hashPassword()
}

func (u *User) BeforeUpdate() {
	u.EntitySafe.BeforeUpdate()
	u.hashPassword()
}

func (u *User) hashPassword() {
	if u.Password != "" {
		passwordSum := sha256.Sum256([]byte(u.Password))
		u.PasswordHash = fmt.Sprintf("%x", passwordSum)
	}
}
