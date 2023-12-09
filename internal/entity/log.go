package entity

import "time"

type Log struct {
	ID        uint64    `gorm:"primaryKey; autoIncrement" json:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	Message   string    `gorm:"not null; <-:create"`
}
