package db

import (
	"time"

	"gorm.io/gorm"
)

type Entity[TEntity IEntity] struct {
	ID        uint64         `gorm:"primaryKey; autoIncrement" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Version   uint64         `gorm:"default:1; <-:update" json:"version"`
}

// AfterCreate implements IEntity
func (entity *Entity[TEntity]) AfterCreate(tx *gorm.DB) (err error) {
	return
}

// AfterDelete implements IEntity
func (entity *Entity[TEntity]) AfterDelete(tx *gorm.DB) (err error) {
	return
}

// AfterFind implements IEntity
func (entity *Entity[TEntity]) AfterFind(tx *gorm.DB) (err error) {
	return
}

// AfterSave implements IEntity
func (entity *Entity[TEntity]) AfterSave(tx *gorm.DB) (err error) {
	return
}

// AfterUpdate implements IEntity
func (entity *Entity[TEntity]) AfterUpdate(tx *gorm.DB) (err error) {
	return
}

// BeforeCreate implements IEntity
func (entity *Entity[TEntity]) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

// BeforeDelete implements IEntity
func (entity *Entity[TEntity]) BeforeDelete(tx *gorm.DB) (err error) {
	return
}

// BeforeSave implements IEntity
func (entity *Entity[TEntity]) BeforeSave(tx *gorm.DB) (err error) {
	return
}

// BeforeUpdate implements IEntity
func (entity *Entity[TEntity]) BeforeUpdate(tx *gorm.DB) (err error) {
	entity.Version += 1
	result := tx.Model(new(TEntity)).Where("id = ?", entity.ID).Update("Version", entity.Version)
	if result.Error != nil {
		return result.Error
	}
	return
}

func NewEntity[TEntity IEntity](entity *Entity[TEntity]) IEntity {
	return entity
}
