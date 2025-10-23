package model

import (
	"gorm.io/gorm"
)

type Many2ManySetup struct {
	Field     string
	JoinTable any
}

type Setup interface {
	BeforeCreate(tx *gorm.DB) (err error)
	BeforeUpdate(tx *gorm.DB) (err error)
	BeforeDelete(tx *gorm.DB) (err error)
	AfterDelete(tx *gorm.DB) (err error)
	AfterUpdate(tx *gorm.DB) (err error)
	AfterCreate(tx *gorm.DB) (err error)
	AfterFind(tx *gorm.DB) (err error)
	Many2ManySetups() (setups []Many2ManySetup)
}
