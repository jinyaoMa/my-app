package entity

import (
	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
	"majinyao.cn/my-app/backend/pkg/db/model"
)

type GroupRole struct {
	model.Model
	model.ModelReserved

	GroupId datatype.Id `gorm:"comment:Group Id;"`
	Group   Group

	RoleId datatype.Id `gorm:"comment:Role Id;"`
	Role   Role
}

// AfterCreate implements model.Setup.
func (r *GroupRole) AfterCreate(tx *gorm.DB) (err error) {
	return
}

// AfterDelete implements model.Setup.
func (r *GroupRole) AfterDelete(tx *gorm.DB) (err error) {
	return
}

// AfterFind implements model.Setup.
func (r *GroupRole) AfterFind(tx *gorm.DB) (err error) {
	return
}

// AfterUpdate implements model.Setup.
func (r *GroupRole) AfterUpdate(tx *gorm.DB) (err error) {
	return
}

// BeforeCreate implements model.Setup.
func (r *GroupRole) BeforeCreate(tx *gorm.DB) (err error) {
	if err = r.ModelReserved.BeforeCreate(tx); err != nil {
		return
	}
	return r.Model.BeforeCreate(tx)
}

// BeforeDelete implements model.Setup.
// Subtle: this method shadows the method (Reserved).BeforeDelete of GroupRole.Reserved.
func (r *GroupRole) BeforeDelete(tx *gorm.DB) (err error) {
	return r.ModelReserved.BeforeDelete(tx)
}

// BeforeUpdate implements model.Setup.
// Subtle: this method shadows the method (Reserved).BeforeUpdate of GroupRole.Reserved.
func (r *GroupRole) BeforeUpdate(tx *gorm.DB) (err error) {
	return r.ModelReserved.BeforeUpdate(tx)
}

// Many2ManySetups implements model.Setup.
func (r *GroupRole) Many2ManySetups() (setups []model.Many2ManySetup) {
	return
}
