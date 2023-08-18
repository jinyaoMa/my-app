package entity

import (
	"my-app/backend/pkg/crypto"
	"my-app/backend/pkg/snowflake"

	"gorm.io/gorm"
)

type Entity struct {
	EntityBase
	ID int64 `gorm:"primaryKey; autoIncrement"`
}

var (
	idGenerator snowflake.Interface
	cipher      crypto.Interface
)

func IdGenerator(abc ...snowflake.Interface) snowflake.Interface {
	if len(abc) == 1 {
		idGenerator = abc[0]
	}
	return idGenerator
}

func Cipher(abc ...crypto.Interface) crypto.Interface {
	if len(abc) == 1 {
		cipher = abc[0]
	}
	return cipher
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
