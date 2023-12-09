package entity

import "time"

type UserPassword struct {
	ID        uint64    `gorm:"primaryKey; autoIncrement" json:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	Account   string    `gorm:"index; not null; <-:create"`
	Password  string    `gorm:"not null; <-:create"`
}
