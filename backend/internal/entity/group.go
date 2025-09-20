package entity

import "majinyao.cn/my-app/backend/pkg/db/model"

type Group struct {
	model.Model
	model.Reserved
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

func (g *Group) GetM2MSetups() []model.M2MSetup {
	return []model.M2MSetup{
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
