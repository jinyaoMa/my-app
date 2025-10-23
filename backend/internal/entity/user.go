package entity

import (
	"database/sql"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
	"majinyao.cn/my-app/backend/pkg/db/model"
)

type User struct {
	model.Model
	model.ModelReserved
	Account                    datatype.Sensitive  `gorm:"uniqueIndex;not null;check:account <> '';size:64;comment:User Account;"`
	Password                   datatype.Password   `gorm:"not null;check:password <> '';size:32;comment:User Password;"`
	Name                       string              `gorm:"index;not null;check:name <> '';size:16;comment:User Name;"`
	Description                sql.NullString      `gorm:"size:254;comment:User Description;"`
	TemporaryPassword          *datatype.Sensitive `gorm:"<-:update;uniqueIndex;size:254;comment:User Temporary Password;"`
	TemporaryPasswordExpiredAt sql.NullTime        `gorm:"<-:update;comment:User Temporary Password Expired At;"`

	UserPasswords []UserPassword

	UserRoles []UserRole
	Roles     []Role `gorm:"many2many:user_roles;"`

	GroupUsers []GroupUser
	Groups     []Group `gorm:"many2many:group_users;"`

	FileUsers []FileUser
	Files     []File `gorm:"many2many:file_users;"`
}

// AfterCreate implements model.Setup.
func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	return
}

// AfterDelete implements model.Setup.
func (u *User) AfterDelete(tx *gorm.DB) (err error) {
	return
}

// AfterFind implements model.Setup.
func (u *User) AfterFind(tx *gorm.DB) (err error) {
	return
}

// AfterUpdate implements model.Setup.
func (u *User) AfterUpdate(tx *gorm.DB) (err error) {
	return
}

// BeforeCreate implements model.Setup.
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if err = u.ModelReserved.BeforeCreate(tx); err != nil {
		return
	}
	return u.Model.BeforeCreate(tx)
}

// BeforeDelete implements model.Setup.
// Subtle: this method shadows the method (Reserved).BeforeDelete of User.Reserved.
func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	return u.ModelReserved.BeforeDelete(tx)
}

// BeforeUpdate implements model.Setup.
// Subtle: this method shadows the method (Reserved).BeforeUpdate of User.Reserved.
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	return u.ModelReserved.BeforeUpdate(tx)
}

// Many2ManySetups implements model.Setup.
func (u *User) Many2ManySetups() (setups []model.Many2ManySetup) {
	return []model.Many2ManySetup{
		{
			Field:     "Roles",
			JoinTable: new(UserRole),
		},
		{
			Field:     "Groups",
			JoinTable: new(GroupUser),
		},
		{
			Field:     "Files",
			JoinTable: new(FileUser),
		},
	}
}
