package entity

import (
	"my-app/backend/pkg/snowflake"
	"time"
)

type IEntity interface {
	// Set Id generator for entity
	SetSnowflake(snowflake.ISnowflake)
}

type Entity struct {
	EntityBase `xorm:"extends"`
	DeletedAt  time.Time `xorm:"deleted"`
}

type EntityBase struct {
	snowflake snowflake.ISnowflake `xorm:"-"`

	Id         int64
	CreatedAt  time.Time `xorm:"created"`
	ModifiedAt time.Time `xorm:"updated"`
	Version    int64     `xorm:"version default(1)"`
}

// SetSnowflake implements IEntity
func (e *EntityBase) SetSnowflake(snowflake snowflake.ISnowflake) {
	e.snowflake = snowflake
}

func (e *EntityBase) BeforeInsert() {
	if e.Id == 0 && e.snowflake != nil {
		e.Id = e.snowflake.Generate()
	}
}

func NewEntityBase(entityBase *EntityBase) IEntity {
	return entityBase
}
