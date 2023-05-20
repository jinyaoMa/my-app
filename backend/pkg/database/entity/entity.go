package entity

import (
	"gorm.io/gorm"
)

type Entity struct {
	EntityBase
	ID int64 `gorm:"primaryKey; autoIncrement"`
}

func (e *Entity) BeforeCreate(tx *gorm.DB) (err error) {
	if err = e.EntityBase.BeforeCreate(tx); err != nil {
		return
	}

	if e != nil && e.ID == 0 && idGenerator != nil {
		e.ID = idGenerator.Generate()
	}
	return
}
