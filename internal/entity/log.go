package entity

import "my-app/pkg/db"

type Log struct {
	db.Entity[*Log]
	Message string `gorm:"not null; <-:create"`
}
