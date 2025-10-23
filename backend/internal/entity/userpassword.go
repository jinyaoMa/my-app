package entity

import (
	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
	"majinyao.cn/my-app/backend/pkg/db/model"
)

type UserPassword struct {
	model.Model
	Password datatype.Sensitive `gorm:"not null;check:password <> '';size:254;comment:Historical Password;"`

	UserId datatype.Id `gorm:"comment:User Id;"`
	User   User
}

// AfterCreate implements model.Setup.
func (p *UserPassword) AfterCreate(tx *gorm.DB) (err error) {
	return
}

// AfterDelete implements model.Setup.
func (p *UserPassword) AfterDelete(tx *gorm.DB) (err error) {
	return
}

// AfterFind implements model.Setup.
func (p *UserPassword) AfterFind(tx *gorm.DB) (err error) {
	return
}

// AfterUpdate implements model.Setup.
func (p *UserPassword) AfterUpdate(tx *gorm.DB) (err error) {
	return
}

// BeforeCreate implements model.Setup.
// Subtle: this method shadows the method (Model).BeforeCreate of UserPassword.Model.
func (p *UserPassword) BeforeCreate(tx *gorm.DB) (err error) {
	return p.Model.BeforeCreate(tx)
}

// BeforeDelete implements model.Setup.
func (p *UserPassword) BeforeDelete(tx *gorm.DB) (err error) {
	return
}

// BeforeUpdate implements model.Setup.
func (p *UserPassword) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

// Many2ManySetups implements model.Setup.
func (p *UserPassword) Many2ManySetups() (setups []model.Many2ManySetup) {
	return
}
