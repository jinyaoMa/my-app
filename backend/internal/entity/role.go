package entity

import "majinyao.cn/my-app/backend/pkg/db/model"

type Role struct {
	model.Model
	model.Reserved
	Code        string `gorm:"index;not null;size:254;comment:Role Code;"`
	Name        string `gorm:"index;size:254;comment:Role Name;"`
	Description string `gorm:"size:254;comment:Role Description;"`

	RolePermissions []RolePermission
	Permissions     []Permission `gorm:"many2many:role_permissions;"`

	UserRoles []UserRole
	Users     []User `gorm:"many2many:user_roles;"`

	GroupRoles []GroupRole
	Groups     []Group `gorm:"many2many:group_roles;"`
}

func (r *Role) GetM2MSetups() []model.M2MSetup {
	return []model.M2MSetup{
		{
			Model:     new(Role),
			Field:     "Permissions",
			JoinTable: new(RolePermission),
		},
		{
			Model:     new(Role),
			Field:     "Users",
			JoinTable: new(UserRole),
		},
		{
			Model:     new(Role),
			Field:     "Groups",
			JoinTable: new(GroupRole),
		},
	}
}
