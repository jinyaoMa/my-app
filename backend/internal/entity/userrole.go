package entity

import (
	"database/sql"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
	"majinyao.cn/my-app/backend/pkg/db/model"
)

type UserRole struct {
	model.Model

	UserId datatype.Id `gorm:"uniqueIndex:idx_user_roles_user_default;comment:User Id;"`
	User   User

	RoleId datatype.Id `gorm:"comment:Role Id;"`
	Role   Role

	IsDefault sql.NullBool `gorm:"uniqueIndex:idx_user_roles_user_default;comment:Is Default Role;"`
}

// AfterCreate implements model.Setup.
func (r *UserRole) AfterCreate(tx *gorm.DB) (err error) {
	return
}

// AfterDelete implements model.Setup.
func (r *UserRole) AfterDelete(tx *gorm.DB) (err error) {
	return
}

// AfterFind implements model.Setup.
func (r *UserRole) AfterFind(tx *gorm.DB) (err error) {
	return
}

// AfterUpdate implements model.Setup.
func (r *UserRole) AfterUpdate(tx *gorm.DB) (err error) {
	return
}

// BeforeCreate implements model.Setup.
// Subtle: this method shadows the method (Model).BeforeCreate of UserRole.Model.
func (r *UserRole) BeforeCreate(tx *gorm.DB) (err error) {
	return r.Model.BeforeCreate(tx)
}

// BeforeDelete implements model.Setup.
func (r *UserRole) BeforeDelete(tx *gorm.DB) (err error) {
	return
}

// BeforeUpdate implements model.Setup.
func (r *UserRole) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

// Many2ManySetups implements model.Setup.
func (r *UserRole) Many2ManySetups() (setups []model.Many2ManySetup) {
	return
}
