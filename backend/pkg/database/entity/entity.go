package entity

import (
	"my-app/backend/pkg/database/interfaces"
	"time"

	"gorm.io/gorm"
)

type Entity struct {
	EntityBase
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type EntityBase struct {
	ID        int64     `gorm:"primaryKey; autoIncrement"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Version   int64     `gorm:"default:1; <-:update"`
}

func NewEntityBase(entityBase EntityBase) interfaces.IEntity {
	return &entityBase
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
	if e != nil && e.ID == 0 && snowflake != nil {
		e.ID = snowflake.Generate()
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
	if e != nil {
		e.Version += 1
	}
	return
}
