package db

import (
	"my-app/backend/pkg/code"
	"my-app/backend/pkg/enc"
	"my-app/backend/pkg/id"
	"time"

	"gorm.io/gorm"
)

type Entity struct {
	EntityBase
	IdGenerator   id.IID      `gorm:"-"`
	CodeGenerator code.ICode  `gorm:"-"`
	DataCipher    enc.ICipher `gorm:"-"`
	ID            int64       `gorm:"primaryKey; autoIncrement"`
}

// SetCodeGenerator implements IEntity.
func (e *Entity) SetCodeGenerator(codeGenerator code.ICode) {
	e.CodeGenerator = codeGenerator
}

// SetDataCipher implements IEntity.
func (e *Entity) SetDataCipher(dataCipher enc.ICipher) {
	e.DataCipher = dataCipher
}

// SetIdGenerator implements IEntity.
func (e *Entity) SetIdGenerator(idGenerator id.IID) {
	e.IdGenerator = idGenerator
}

func (e *Entity) BeforeCreate(tx *gorm.DB) (err error) {
	if err = e.EntityBase.BeforeCreate(tx); err != nil {
		return
	}

	if e != nil && e.ID == 0 && e.IdGenerator != nil {
		e.ID = e.IdGenerator.Generate()
	}
	return
}

func NewIEntity(entity *Entity) IEntity {
	return entity
}

type EntityBase struct {
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Version   int64          `gorm:"default:1; <-:update"`
}

// AfterCreate implements IEntity
func (e *EntityBase) AfterCreate(tx *gorm.DB) (err error) {
	return
}

// AfterDelete implements IEntity
func (e *EntityBase) AfterDelete(tx *gorm.DB) (err error) {
	return
}

// AfterFind implements IEntity
func (e *EntityBase) AfterFind(tx *gorm.DB) (err error) {
	return
}

// AfterSave implements IEntity
func (e *EntityBase) AfterSave(tx *gorm.DB) (err error) {
	return
}

// AfterUpdate implements IEntity
func (e *EntityBase) AfterUpdate(tx *gorm.DB) (err error) {
	return
}

// BeforeCreate implements IEntity
func (e *EntityBase) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

// BeforeDelete implements IEntity
func (e *EntityBase) BeforeDelete(tx *gorm.DB) (err error) {
	return
}

// BeforeSave implements IEntity
func (e *EntityBase) BeforeSave(tx *gorm.DB) (err error) {
	return
}

// BeforeUpdate implements IEntity
func (e *EntityBase) BeforeUpdate(tx *gorm.DB) (err error) {
	if e != nil {
		e.Version += 1
	}
	return
}

func NewIEntityBase(entity *EntityBase) IEntityBase {
	return entity
}
