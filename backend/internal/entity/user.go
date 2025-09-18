package entity

import (
	"majinyao.cn/my-app/backend/pkg/db"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
)

type User struct {
	db.Entity
	db.EntityReserved
	Account     datatype.Encrypted `gorm:"uniqueIndex;not null;size:64;comment:User Account;"`
	Password    datatype.Password  `gorm:"not null;size:32;comment:User Password;"`
	Name        string             `gorm:"index;size:16;comment:User Name;"`
	Description string             `gorm:"size:254;comment:User Description;"`

	UserPasswords []UserPassword

	UserRoles []UserRole
	Roles     []Role `gorm:"many2many:user_roles;"`

	GroupUsers []GroupUser
	Groups     []Group `gorm:"many2many:group_users;"`

	FileUsers []FileUser
	Files     []File `gorm:"many2many:file_users;"`
}

func (u *User) GetEntityM2MSetups() []db.EntityM2MSetup {
	return []db.EntityM2MSetup{
		{
			Model:     new(User),
			Field:     "Roles",
			JoinTable: new(UserRole),
		},
		{
			Model:     new(User),
			Field:     "Groups",
			JoinTable: new(GroupUser),
		},
		{
			Model:     new(User),
			Field:     "Files",
			JoinTable: new(FileUser),
		},
	}
}
