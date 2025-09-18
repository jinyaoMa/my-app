package entity

import (
	"majinyao.cn/my-app/backend/pkg/db"
)

type Group struct {
	db.Entity
	db.EntityReserved
	Code        string `gorm:"index;not null;size:254;comment:Group Code;"`
	Name        string `gorm:"index;size:254;comment:Group Name;"`
	Description string `gorm:"size:254;comment:Group Description;"`

	GroupUsers []GroupUser
	Users      []User `gorm:"many2many:group_users;"`

	GroupRoles []GroupRole
	Roles      []Role `gorm:"many2many:group_roles;"`

	FileGroups []FileGroup
	Files      []File `gorm:"many2many:file_groups;"`
}

func (g *Group) GetM2MSetups() []db.EntityM2MSetup {
	return []db.EntityM2MSetup{
		{
			Model:     new(Group),
			Field:     "Users",
			JoinTable: new(GroupUser),
		},
		{
			Model:     new(Group),
			Field:     "Roles",
			JoinTable: new(GroupRole),
		},
		{
			Model:     new(Group),
			Field:     "Files",
			JoinTable: new(FileGroup),
		},
	}
}
