package entity

import (
	"my-app/backend/pkg/database/interfaces"
	iSnowflake "my-app/backend/pkg/snowflake/interfaces"
	"time"

	"xorm.io/xorm"
)

type Entity struct {
	EntityBase `xorm:"extends"`
	DeletedAt  time.Time `xorm:"deleted"`
}

type EntityBase struct {
	snowflake iSnowflake.ISnowflake `xorm:"-"`

	Id         int64
	CreatedAt  time.Time `xorm:"created"`
	ModifiedAt time.Time `xorm:"updated"`
	Version    int64     `xorm:"version default(1)"`
}

// AfterDelete implements interfaces.IEntity
func (*EntityBase) AfterDelete() {}

// AfterInsert implements interfaces.IEntity
func (*EntityBase) AfterInsert() {}

// AfterLoad implements interfaces.IEntity
func (*EntityBase) AfterLoad(*xorm.Session) {}

// AfterSet implements interfaces.IEntity
func (*EntityBase) AfterSet(name string, cell xorm.Cell) {}

// AfterUpdate implements interfaces.IEntity
func (*EntityBase) AfterUpdate() {}

// BeforeDelete implements interfaces.IEntity
func (*EntityBase) BeforeDelete() {}

// BeforeSet implements interfaces.IEntity
func (*EntityBase) BeforeSet(name string, cell xorm.Cell) {}

// BeforeUpdate implements interfaces.IEntity
func (*EntityBase) BeforeUpdate() {}

// BeforeInsert implements interfaces.IEntity
func (e *EntityBase) BeforeInsert() {
	if e.Id == 0 && e.snowflake != nil {
		e.Id = e.snowflake.Generate()
	}
}

// SetSnowflake implements interfaces.IEntity
func (e *EntityBase) SetSnowflake(snowflake iSnowflake.ISnowflake) {
	e.snowflake = snowflake
}

func NewEntityBase(entityBase *EntityBase) interfaces.IEntity {
	return entityBase
}
