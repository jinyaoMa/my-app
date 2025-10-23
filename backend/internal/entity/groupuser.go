package entity

import (
	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
	"majinyao.cn/my-app/backend/pkg/db/model"
)

type GroupUser struct {
	model.Model

	GroupId datatype.Id `gorm:"comment:Group Id;"`
	Group   Group

	UserId datatype.Id `gorm:"comment:User Id;"`
	User   User
}

// AfterCreate implements model.Setup.
func (u *GroupUser) AfterCreate(tx *gorm.DB) (err error) {
	return
}

// AfterDelete implements model.Setup.
func (u *GroupUser) AfterDelete(tx *gorm.DB) (err error) {
	return
}

// AfterFind implements model.Setup.
func (u *GroupUser) AfterFind(tx *gorm.DB) (err error) {
	return
}

// AfterUpdate implements model.Setup.
func (u *GroupUser) AfterUpdate(tx *gorm.DB) (err error) {
	return
}

// BeforeCreate implements model.Setup.
// Subtle: this method shadows the method (Model).BeforeCreate of GroupUser.Model.
func (u *GroupUser) BeforeCreate(tx *gorm.DB) (err error) {
	return u.Model.BeforeCreate(tx)
}

// BeforeDelete implements model.Setup.
func (u *GroupUser) BeforeDelete(tx *gorm.DB) (err error) {
	return
}

// BeforeUpdate implements model.Setup.
func (u *GroupUser) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

// Many2ManySetups implements model.Setup.
func (u *GroupUser) Many2ManySetups() (setups []model.Many2ManySetup) {
	return
}
