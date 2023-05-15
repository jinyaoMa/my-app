package entity

import (
	"my-app/backend/pkg/database/interfaces"
	iSnowflake "my-app/backend/pkg/snowflake/interfaces"
	"time"

	"gorm.io/gorm"
)

type Entity struct {
	EntityBase
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type EntityBase struct {
	snowflake iSnowflake.ISnowflake

	ID        int64     `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:""`
	UpdatedAt time.Time `gorm:""`
	Version   int64     `gorm:"default:1"`
}

// SetSnowflake implements interfaces.IEntity
func (e *EntityBase) SetSnowflake(snowflake iSnowflake.ISnowflake) {
	e.snowflake = snowflake
}

// AfterCreate implements interfaces.IEntity
func (e *EntityBase) AfterCreate(tx *gorm.DB) (err error) {
	return
}

// AfterDelete implements interfaces.IEntity
func (e *EntityBase) AfterDelete(tx *gorm.DB) (err error) {
	return
}

// AfterFind implements interfaces.IEntity
func (e *EntityBase) AfterFind(tx *gorm.DB) (err error) {
	return
}

// AfterSave implements interfaces.IEntity
func (e *EntityBase) AfterSave(tx *gorm.DB) (err error) {
	return
}

// AfterUpdate implements interfaces.IEntity
func (e *EntityBase) AfterUpdate(tx *gorm.DB) (err error) {
	return
}

// BeforeCreate implements interfaces.IEntity
func (e *EntityBase) BeforeCreate(tx *gorm.DB) (err error) {
	if e.ID == 0 && e.snowflake != nil {
		e.ID = e.snowflake.Generate()
	}
	return
}

// BeforeDelete implements interfaces.IEntity
func (e *EntityBase) BeforeDelete(tx *gorm.DB) (err error) {
	return
}

// BeforeSave implements interfaces.IEntity
func (e *EntityBase) BeforeSave(tx *gorm.DB) (err error) {
	return
}

// BeforeUpdate implements interfaces.IEntity
func (e *EntityBase) BeforeUpdate(tx *gorm.DB) (err error) {
	e.Version += 1
	return
}

func NewEntityBase(entityBase *EntityBase) interfaces.IEntity {
	return entityBase
}
