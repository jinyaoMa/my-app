package entity

import "my-app/pkg/db"

type Option struct {
	db.Entity[*Option]
	Name      string `gorm:"unique; index; not null" json:"name"`
	Value     string `gorm:"default:''" json:"value"` // may be encrypted
	Encrypted bool   `gorm:"default:false" json:"encrypted"`
}
