package entity

import (
	"my-app/backend/pkg/snowflake"
	"time"
)

type IEntity interface {
	SetSnowflake(*snowflake.Snowflake)
}

type Entity struct {
	EntityBase `xorm:"extends"`
	DeletedAt  time.Time `xorm:"deleted"`
}

type EntityBase struct {
	_snowflake *snowflake.Snowflake `xorm:"-"`

	Id         int64
	CreatedAt  time.Time `xorm:"created"`
	ModifiedAt time.Time `xorm:"updated"`
	Version    int64     `xorm:"version default(1)"`
}

// SetSnowflake implements IEntity
func (e *EntityBase) SetSnowflake(snowflake *snowflake.Snowflake) {
	e._snowflake = snowflake
}

func (e *EntityBase) BeforeInsert() {
	if e.Id == 0 && e._snowflake != nil {
		e.Id = e._snowflake.Generate()
	}
}

func NewEntityBase(entityBase *EntityBase) IEntity {
	return entityBase
}
