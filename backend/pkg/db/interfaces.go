package db

import (
	"my-app/backend/pkg/code"
	"my-app/backend/pkg/db/param"
	"my-app/backend/pkg/enc"
	"my-app/backend/pkg/id"

	"gorm.io/gorm"
)

type IEntity interface {
	IEntityBase

	SetIdGenerator(idGenerator id.IID)
	SetCodeGenerator(codeGenerator code.ICode)
	SetDataCipher(dataCipher enc.ICipher)
}

type IEntityBase interface {
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
	BuildQuery(criteria *param.Criteria, condition param.QueryCondition, includes ...string) (db *DB)
	Query(criteria *param.Criteria, condition param.QueryCondition, includes ...string) (entities []TEntity, err error)
	GetById(id int64) (entity TEntity, err error)
	All() (entities []TEntity, err error)
	FindOne(condition param.QueryCondition) (entity TEntity, err error)
	Save(entity TEntity) (affected int64, err error)
	SaveAll(entities []TEntity) (affected int64, err error)
	Delete(id int64) (affected int64, err error)
}
