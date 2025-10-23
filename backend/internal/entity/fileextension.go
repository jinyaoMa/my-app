package entity

import (
	"database/sql"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
	"majinyao.cn/my-app/backend/pkg/db/model"
)

type FileExtension struct {
	model.Model
	model.ModelReserved
	Ext  string         `gorm:"<-:create;index;not null;check:ext LIKE '.%';size:254;comment:File Extension (.ext);"`
	Name sql.NullString `gorm:"index;size:254;comment:File Extension Display Name;"`
	Mime sql.NullString `gorm:"index;size:254;comment:File Extension Mime;"`

	FileCategoryId *datatype.Id `gorm:"comment:File Category Id;"`
	FileCategory   *FileCategory

	Files []File
}

// AfterCreate implements model.Setup.
func (e *FileExtension) AfterCreate(tx *gorm.DB) (err error) {
	return
}

// AfterDelete implements model.Setup.
func (e *FileExtension) AfterDelete(tx *gorm.DB) (err error) {
	return
}

// AfterFind implements model.Setup.
func (e *FileExtension) AfterFind(tx *gorm.DB) (err error) {
	return
}

// AfterUpdate implements model.Setup.
func (e *FileExtension) AfterUpdate(tx *gorm.DB) (err error) {
	return
}

// BeforeCreate implements model.Setup.
func (e *FileExtension) BeforeCreate(tx *gorm.DB) (err error) {
	if err = e.ModelReserved.BeforeCreate(tx); err != nil {
		return
	}
	return e.Model.BeforeCreate(tx)
}

// BeforeDelete implements model.Setup.
// Subtle: this method shadows the method (Reserved).BeforeDelete of FileExtension.Reserved.
func (e *FileExtension) BeforeDelete(tx *gorm.DB) (err error) {
	return e.ModelReserved.BeforeDelete(tx)
}

// BeforeUpdate implements model.Setup.
// Subtle: this method shadows the method (Reserved).BeforeUpdate of FileExtension.Reserved.
func (e *FileExtension) BeforeUpdate(tx *gorm.DB) (err error) {
	return e.ModelReserved.BeforeUpdate(tx)
}

// Many2ManySetups implements model.Setup.
func (e *FileExtension) Many2ManySetups() (setups []model.Many2ManySetup) {
	return
}
