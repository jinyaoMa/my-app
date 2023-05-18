package entity

import (
	"crypto/sha256"
	"fmt"
	"my-app/backend/pkg/utility"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Entity
	Account               string    `gorm:"size:64; unique; index; not null"`
	Password              string    `gorm:"-:all"`
	PasswordHash          string    `gorm:"size:64; not null"`
	Verification          string    `gorm:"size:6; not null"`
	VerificationExpiredAt time.Time `gorm:"not null"`
	IsFrozen              bool      `gorm:"default:false"`
	OldPasswords          []*UserPassword
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if err = u.Entity.BeforeCreate(tx); err != nil {
		return
	}

	if u != nil {
		err = u.hashPassword(tx)
	}
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if err = u.Entity.BeforeUpdate(tx); err != nil {
		return
	}

	if u != nil {
		err = u.hashPassword(tx)
	}
	return
}

func (u *User) hashPassword(tx *gorm.DB) (err error) {
	if u.Password != "" {
		passwordSum := sha256.Sum256([]byte(u.Password))
		u.PasswordHash = fmt.Sprintf("%x", passwordSum)
		u.OldPasswords = append(u.OldPasswords, &UserPassword{
			PasswordHash: u.PasswordHash,
		})
	}
	return
}

func (u *User) FillVerification(size int, expiredAt time.Time) *User {
	u.Verification = utility.NewRandom().GenerateCode(size)
	u.VerificationExpiredAt = expiredAt
	return u
}
