package entity

import (
	"database/sql"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
	"majinyao.cn/my-app/backend/pkg/db/model"
)

type FileGroup struct {
	model.Model

	FileId datatype.Id `gorm:"comment:File Id;"`
	File   File

	GroupId datatype.Id `gorm:"comment:Group Id;"`
	Group   Group

	NoCreate sql.NullBool `gorm:"comment:No Create;"`
	NoRead   sql.NullBool `gorm:"comment:No Read;"`
	NoUpdate sql.NullBool `gorm:"comment:No Update;"`
	NoDelete sql.NullBool `gorm:"comment:No Delete;"`
}

// AfterCreate implements model.Setup.
func (g *FileGroup) AfterCreate(tx *gorm.DB) (err error) {
	return
}

// AfterDelete implements model.Setup.
func (g *FileGroup) AfterDelete(tx *gorm.DB) (err error) {
	return
}

// AfterFind implements model.Setup.
func (g *FileGroup) AfterFind(tx *gorm.DB) (err error) {
	return
}

// AfterUpdate implements model.Setup.
func (g *FileGroup) AfterUpdate(tx *gorm.DB) (err error) {
	return
}

// BeforeCreate implements model.Setup.
// Subtle: this method shadows the method (Model).BeforeCreate of FileGroup.Model.
func (g *FileGroup) BeforeCreate(tx *gorm.DB) (err error) {
	return g.Model.BeforeCreate(tx)
}

// BeforeDelete implements model.Setup.
func (g *FileGroup) BeforeDelete(tx *gorm.DB) (err error) {
	return
}

// BeforeUpdate implements model.Setup.
func (g *FileGroup) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

// Many2ManySetups implements model.Setup.
func (g *FileGroup) Many2ManySetups() (setups []model.Many2ManySetup) {
	return
}
