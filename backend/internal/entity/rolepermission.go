package entity

import (
	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
	"majinyao.cn/my-app/backend/pkg/db/model"
)

type RolePermission struct {
	model.Model
	model.ModelReserved

	RoleId datatype.Id `gorm:"comment:Role Id;"`
	Role   Role

	PermissionId datatype.Id `gorm:"comment:Permission Id;"`
	Permission   Permission
}

// AfterCreate implements model.Setup.
func (p *RolePermission) AfterCreate(tx *gorm.DB) (err error) {
	return
}

// AfterDelete implements model.Setup.
func (p *RolePermission) AfterDelete(tx *gorm.DB) (err error) {
	return
}

// AfterFind implements model.Setup.
func (p *RolePermission) AfterFind(tx *gorm.DB) (err error) {
	return
}

// AfterUpdate implements model.Setup.
func (p *RolePermission) AfterUpdate(tx *gorm.DB) (err error) {
	return
}

// BeforeCreate implements model.Setup.
func (p *RolePermission) BeforeCreate(tx *gorm.DB) (err error) {
	if err = p.ModelReserved.BeforeCreate(tx); err != nil {
		return
	}
	return p.Model.BeforeCreate(tx)
}

// BeforeDelete implements model.Setup.
// Subtle: this method shadows the method (Reserved).BeforeDelete of RolePermission.Reserved.
func (p *RolePermission) BeforeDelete(tx *gorm.DB) (err error) {
	return p.ModelReserved.BeforeDelete(tx)
}

// BeforeUpdate implements model.Setup.
// Subtle: this method shadows the method (Reserved).BeforeUpdate of RolePermission.Reserved.
func (p *RolePermission) BeforeUpdate(tx *gorm.DB) (err error) {
	return p.ModelReserved.BeforeUpdate(tx)
}

// Many2ManySetups implements model.Setup.
func (p *RolePermission) Many2ManySetups() (setups []model.Many2ManySetup) {
	return
}
