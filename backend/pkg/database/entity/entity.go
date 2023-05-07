package entity

import (
	"my-app/backend/pkg/snowflake"
	"time"
)

type IEntity interface {
	SetSnowflake(*snowflake.Snowflake)
	BeforeInsert()
	BeforeUpdate()
	BeforeDelete()
	AfterLoad()
	AfterInsert()
	AfterUpdate()
	AfterDelete()
}

type Entity struct {
	IEntity
	_snowflake *snowflake.Snowflake `xorm:"-"`

	Id         int64
	CreatedAt  time.Time `xorm:"created"`
	ModifiedAt time.Time `xorm:"updated"`
	Version    int64     `xorm:"version"`
}

func (e *Entity) SetSnowflake(snowflake *snowflake.Snowflake) {
	e._snowflake = snowflake
}

func (e *Entity) BeforeInsert() {
	if e.Id == 0 && e._snowflake != nil {
		e.Id = e._snowflake.Generate()
	}
}

type EntitySafe struct {
	Entity
	DeletedAt time.Time `xorm:"deleted"`
}
