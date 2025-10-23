package entity

import (
	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/db/model"
)

type FileCategory struct {
	model.Model
	model.ModelReserved
	Code string `gorm:"<-:create;uniqueIndex;not null;check:code <> '';size:254;comment:File Category Code;"`
	Name string `gorm:"index;not null;check:name <> '';size:254;comment:File Category Name;"`

	FileExtensions []FileExtension
}

// AfterCreate implements model.Setup.
func (c *FileCategory) AfterCreate(tx *gorm.DB) (err error) {
	return
}

// AfterDelete implements model.Setup.
func (c *FileCategory) AfterDelete(tx *gorm.DB) (err error) {
	return
}

// AfterFind implements model.Setup.
func (c *FileCategory) AfterFind(tx *gorm.DB) (err error) {
	return
}

// AfterUpdate implements model.Setup.
func (c *FileCategory) AfterUpdate(tx *gorm.DB) (err error) {
	return
}

// BeforeCreate implements model.Setup.
func (c *FileCategory) BeforeCreate(tx *gorm.DB) (err error) {
	if err = c.ModelReserved.BeforeCreate(tx); err != nil {
		return
	}
	return c.Model.BeforeCreate(tx)
}

// BeforeDelete implements model.Setup.
// Subtle: this method shadows the method (Reserved).BeforeDelete of FileCategory.Reserved.
func (c *FileCategory) BeforeDelete(tx *gorm.DB) (err error) {
	return c.ModelReserved.BeforeDelete(tx)
}

// BeforeUpdate implements model.Setup.
// Subtle: this method shadows the method (Reserved).BeforeUpdate of FileCategory.Reserved.
func (c *FileCategory) BeforeUpdate(tx *gorm.DB) (err error) {
	return c.ModelReserved.BeforeUpdate(tx)
}

// Many2ManySetups implements model.Setup.
func (c *FileCategory) Many2ManySetups() (setups []model.Many2ManySetup) {
	return
}
