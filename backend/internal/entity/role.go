package entity

import (
	"database/sql"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/db/model"
)

type Role struct {
	model.Model
	model.ModelReserved
	Code        string         `gorm:"<-:create;uniqueIndex;not null;check:code <> '';size:254;comment:Role Code;"`
	Name        string         `gorm:"index;not null;check:name <> '';size:254;comment:Role Name;"`
	Description sql.NullString `gorm:"size:254;comment:Role Description;"`

	RolePermissions []RolePermission
	Permissions     []Permission `gorm:"many2many:role_permissions;"`

	UserRoles []UserRole
	Users     []User `gorm:"many2many:user_roles;"`

	GroupRoles []GroupRole
	Groups     []Group `gorm:"many2many:group_roles;"`
}

// AfterCreate implements model.Setup.
func (r *Role) AfterCreate(tx *gorm.DB) (err error) {
	return
}

// AfterDelete implements model.Setup.
func (r *Role) AfterDelete(tx *gorm.DB) (err error) {
	return
}

// AfterFind implements model.Setup.
func (r *Role) AfterFind(tx *gorm.DB) (err error) {
	return
}

// AfterUpdate implements model.Setup.
func (r *Role) AfterUpdate(tx *gorm.DB) (err error) {
	return
}

// BeforeCreate implements model.Setup.
func (r *Role) BeforeCreate(tx *gorm.DB) (err error) {
	if err = r.ModelReserved.BeforeCreate(tx); err != nil {
		return
	}
	return r.Model.BeforeCreate(tx)
}

// BeforeDelete implements model.Setup.
// Subtle: this method shadows the method (Reserved).BeforeDelete of Role.Reserved.
func (r *Role) BeforeDelete(tx *gorm.DB) (err error) {
	return r.ModelReserved.BeforeDelete(tx)
}

// BeforeUpdate implements model.Setup.
// Subtle: this method shadows the method (Reserved).BeforeUpdate of Role.Reserved.
func (r *Role) BeforeUpdate(tx *gorm.DB) (err error) {
	return r.ModelReserved.BeforeUpdate(tx)
}

// Many2ManySetups implements model.Setup.
func (r *Role) Many2ManySetups() (setups []model.Many2ManySetup) {
	return []model.Many2ManySetup{
		{
			Field:     "Permissions",
			JoinTable: new(RolePermission),
		},
		{
			Field:     "Users",
			JoinTable: new(UserRole),
		},
		{
			Field:     "Groups",
			JoinTable: new(GroupRole),
		},
	}
}
