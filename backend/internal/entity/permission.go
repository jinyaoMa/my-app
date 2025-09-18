package entity

import (
	"majinyao.cn/my-app/backend/pkg/db"
	"majinyao.cn/my-app/backend/pkg/flag"
)

type Permission struct {
	db.Entity
	db.EntityReserved
	Code        string `gorm:"index;not null;size:254;comment:Permission Code;"`
	Name        string `gorm:"index;size:254;comment:Permission Name;"`
	Description string `gorm:"size:254;comment:Permission Description;"`
	Flag        []byte `gorm:"comment:Permission Flag;"`

	RolePermissions []RolePermission
	Roles           []Role `gorm:"many2many:role_permissions;"`
}

func (p *Permission) GetFlag() flag.IFlag {
	return flag.FromBytes(p.Flag)
}

func (p *Permission) SetFlag(f flag.IFlag) {
	p.Flag = f.ToBytes()
}

func (p *Permission) GetM2MSetups() []db.EntityM2MSetup {
	return []db.EntityM2MSetup{
		{
			Model:     new(Permission),
			Field:     "Roles",
			JoinTable: new(RolePermission),
		},
	}
}
