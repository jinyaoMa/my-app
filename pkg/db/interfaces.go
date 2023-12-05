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

type ICRUD[TEntity IEntity] interface {
	BuildQuery(criteria *QueryCriteria, condition QueryCondition, includes ...string) (tx *gorm.DB, err error)
	Query(criteria *QueryCriteria, condition QueryCondition, includes ...string) (entities []TEntity, err error)
	GetById(id int64, includes ...string) (entity TEntity, err error)
	All(selected ...string) (entities []TEntity, err error)
	FindOne(condition QueryCondition, includes ...string) (entity TEntity, err error)
	Save(entity TEntity) (affected int64, err error)
	SaveAll(entities []TEntity) (affected int64, err error)
	Update(entity TEntity, selected []string, omitted ...string) (affected int64, err error)
	Delete(id int64) (affected int64, err error)
	DeleteBy(condition QueryCondition) (affected int64, err error)
}
