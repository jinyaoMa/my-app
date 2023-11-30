package db

import (
	"time"

	"gorm.io/gorm"
)

type Entity struct {
	ID        int64          `gorm:"primaryKey; autoIncrement" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Version   int64          `gorm:"default:1; <-:update" json:"version"`
}

// AfterCreate implements IEntity
func (entity *Entity) AfterCreate(tx *gorm.DB) (err error) {
	return
}

// AfterDelete implements IEntity
func (entity *Entity) AfterDelete(tx *gorm.DB) (err error) {
	return
}

// AfterFind implements IEntity
func (entity *Entity) AfterFind(tx *gorm.DB) (err error) {
	return
}

// AfterSave implements IEntity
func (entity *Entity) AfterSave(tx *gorm.DB) (err error) {
	return
}

// AfterUpdate implements IEntity
func (entity *Entity) AfterUpdate(tx *gorm.DB) (err error) {
	return
}

// BeforeCreate implements IEntity
func (entity *Entity) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

// BeforeDelete implements IEntity
func (entity *Entity) BeforeDelete(tx *gorm.DB) (err error) {
	return
}

// BeforeSave implements IEntity
func (entity *Entity) BeforeSave(tx *gorm.DB) (err error) {
	return
}

// BeforeUpdate implements IEntity
func (entity *Entity) BeforeUpdate(tx *gorm.DB) (err error) {
	if entity != nil {
		entity.Version += 1
	}
	return
}

func NewEntity(entity *Entity) IEntity {
	return entity
}
