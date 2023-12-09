package entity

import "my-app/pkg/db"

type UserPassword struct {
	db.Entity[*UserPassword]
	Account  string `gorm:"index; not null; <-:create"`
	Password string `gorm:"not null; <-:create"`
}
