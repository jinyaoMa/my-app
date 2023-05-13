package interfaces

import (
	iSnowflake "my-app/backend/pkg/snowflake/interfaces"

	"xorm.io/xorm"
)

type IEntity interface {
	// Set Id generator for entity
	SetSnowflake(iSnowflake.ISnowflake)

	BeforeInsert()
	BeforeUpdate()
	BeforeDelete()
	BeforeSet(name string, cell xorm.Cell)
	AfterSet(name string, cell xorm.Cell)
	AfterLoad(*xorm.Session)
	AfterInsert()
	AfterUpdate()
	AfterDelete()
}
