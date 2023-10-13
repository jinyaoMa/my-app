package entity

import (
	"crypto/sha256"
	"fmt"
	"my-app/backend/pkg/db"
	"time"

	"gorm.io/gorm"
)

type User struct {
	db.Entity

	/* internal fields */
	Account               string    `gorm:"size:64; unique; index; not null"`
	Password              string    `gorm:"-:all"`
	PasswordHash          string    `gorm:"size:64; not null"`
	Verification          string    `gorm:"size:64; not null"`
	VerificationExpiredAt time.Time `gorm:"not null"`
	IsFrozen              bool      `gorm:"default:false"`

	/* relational fields */
	UserPasswords []*UserPassword `gorm:""`
	OwnedFiles    []*File         `gorm:""`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if err = u.Entity.BeforeCreate(tx); err != nil {
		return
	}

	if u != nil {
		if err = u.hashPassword(tx); err != nil {
			return
		}
	}
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if err = u.Entity.BeforeUpdate(tx); err != nil {
		return
	}

	if u != nil {
		if err = u.hashPassword(tx); err != nil {
			return
		}
		if tx.Statement.Changed("PasswordHash") {
			if err = u.AddUserPassword(tx); err != nil {
				return
			}
		}
	}
	return
}

func (u *User) hashPassword(tx *gorm.DB) (err error) {
	if u.Password != "" {
		passwordSum := sha256.Sum256([]byte(u.Password))
		u.PasswordHash = fmt.Sprintf("%x", passwordSum)
	}
	return
}

func (u *User) AddUserPassword(tx *gorm.DB) (err error) {
	if u.PasswordHash != "" {
		u.UserPasswords = append(u.UserPasswords, &UserPassword{
			PasswordHash: u.PasswordHash,
		})
	}
	return
}

func (u *User) FillVerification(size int, expiredAt time.Time) *User {
	u.Verification = db.CodeGenerator().Generate(size)
	u.VerificationExpiredAt = expiredAt
	return u
}
