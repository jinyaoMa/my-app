package db

import (
	"gorm.io/gorm"
)

type IEntity interface {
	BeforeSave(tx *gorm.DB) (err error)
	BeforeCreate(tx *gorm.DB) (err error)
	BeforeUpdate(tx *gorm.DB) (err error)
	BeforeDelete(tx *gorm.DB) (err error)
	AfterFind(tx *gorm.DB) (err error)
	AfterDelete(tx *gorm.DB) (err error)
	AfterUpdate(tx *gorm.DB) (err error)
	AfterCreate(tx *gorm.DB) (err error)
	AfterSave(tx *gorm.DB) (err error)
}
