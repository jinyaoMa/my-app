package entity

import (
	"database/sql"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/db/model"
	"majinyao.cn/my-app/backend/pkg/flag"
)

type Permission struct {
	model.Model
	model.ModelReserved
	Code        string         `gorm:"<-:create;uniqueIndex;not null;check:code <> '';size:254;comment:Permission Code;"`
	Name        string         `gorm:"index;not null;check:name <> '';size:254;comment:Permission Name;"`
	Description sql.NullString `gorm:"size:254;comment:Permission Description;"`
	Flag        sql.RawBytes   `gorm:"comment:Permission Flag;"`

	RolePermissions []RolePermission
	Roles           []Role `gorm:"many2many:role_permissions;"`
}

// AfterCreate implements model.Setup.
func (p *Permission) AfterCreate(tx *gorm.DB) (err error) {
	return
}

// AfterDelete implements model.Setup.
func (p *Permission) AfterDelete(tx *gorm.DB) (err error) {
	return
}

// AfterFind implements model.Setup.
func (p *Permission) AfterFind(tx *gorm.DB) (err error) {
	return
}

// AfterUpdate implements model.Setup.
func (p *Permission) AfterUpdate(tx *gorm.DB) (err error) {
	return
}

// BeforeCreate implements model.Setup.
func (p *Permission) BeforeCreate(tx *gorm.DB) (err error) {
	if err = p.ModelReserved.BeforeCreate(tx); err != nil {
		return
	}
	return p.Model.BeforeCreate(tx)
}

// BeforeDelete implements model.Setup.
// Subtle: this method shadows the method (Reserved).BeforeDelete of Permission.Reserved.
func (p *Permission) BeforeDelete(tx *gorm.DB) (err error) {
	return p.ModelReserved.BeforeDelete(tx)
}

// BeforeUpdate implements model.Setup.
// Subtle: this method shadows the method (Reserved).BeforeUpdate of Permission.Reserved.
func (p *Permission) BeforeUpdate(tx *gorm.DB) (err error) {
	return p.ModelReserved.BeforeUpdate(tx)
}

// Many2ManySetups implements model.Setup.
func (p *Permission) Many2ManySetups() (setups []model.Many2ManySetup) {
	return []model.Many2ManySetup{
		{
			Field:     "Roles",
			JoinTable: new(RolePermission),
		},
	}
}

func (p *Permission) GetFlag() flag.IFlag {
	return flag.FromBytes(p.Flag)
}

func (p *Permission) SetFlag(f flag.IFlag) {
	p.Flag = f.ToBytes()
}

func (p *Permission) FlagBase64(b64 string) {
	f, _ := flag.FromBase64(b64)
	p.Flag = f.ToBytes()
}
