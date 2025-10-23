package entity

import (
	"database/sql"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/db/model"
)

type Group struct {
	model.Model
	model.ModelReserved
	Code        string         `gorm:"<-:create;uniqueIndex;not null;check:code <> '';size:254;comment:Group Code;"`
	Name        string         `gorm:"index;not null;check:name <> '';size:254;comment:Group Name;"`
	Description sql.NullString `gorm:"size:254;comment:Group Description;"`

	GroupUsers []GroupUser
	Users      []User `gorm:"many2many:group_users;"`

	GroupRoles []GroupRole
	Roles      []Role `gorm:"many2many:group_roles;"`

	FileGroups []FileGroup
	Files      []File `gorm:"many2many:file_groups;"`
}

// AfterCreate implements model.Setup.
func (g *Group) AfterCreate(tx *gorm.DB) (err error) {
	return
}

// AfterDelete implements model.Setup.
func (g *Group) AfterDelete(tx *gorm.DB) (err error) {
	return
}

// AfterFind implements model.Setup.
func (g *Group) AfterFind(tx *gorm.DB) (err error) {
	return
}

// AfterUpdate implements model.Setup.
func (g *Group) AfterUpdate(tx *gorm.DB) (err error) {
	return
}

// BeforeCreate implements model.Setup.
func (g *Group) BeforeCreate(tx *gorm.DB) (err error) {
	if err = g.ModelReserved.BeforeCreate(tx); err != nil {
		return
	}
	return g.Model.BeforeCreate(tx)
}

// BeforeDelete implements model.Setup.
// Subtle: this method shadows the method (Reserved).BeforeDelete of Group.Reserved.
func (g *Group) BeforeDelete(tx *gorm.DB) (err error) {
	return g.ModelReserved.BeforeDelete(tx)
}

// BeforeUpdate implements model.Setup.
// Subtle: this method shadows the method (Reserved).BeforeUpdate of Group.Reserved.
func (g *Group) BeforeUpdate(tx *gorm.DB) (err error) {
	return g.ModelReserved.BeforeUpdate(tx)
}

// Many2ManySetups implements model.Setup.
func (g *Group) Many2ManySetups() (setups []model.Many2ManySetup) {
	return []model.Many2ManySetup{
		{
			Field:     "Users",
			JoinTable: new(GroupUser),
		},
		{
			Field:     "Roles",
			JoinTable: new(GroupRole),
		},
		{
			Field:     "Files",
			JoinTable: new(FileGroup),
		},
	}
}
