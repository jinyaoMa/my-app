package db

import (
	"errors"

	"gorm.io/gorm"
)

type EntityReserved struct {
	isSystem bool `gorm:"-"`
	Reserved bool `gorm:"<-:create;index;comment:Is System Reserved;"`
}

func (p *EntityReserved) SysOp(reserved bool) {
	p.isSystem = true
	p.Reserved = reserved
}

func (p *EntityReserved) BeforeCreate(tx *gorm.DB) (err error) {
	if p.isSystem {
		return nil
	}
	if p.Reserved {
		return errors.New("entity can not be system reserved when operation is not by system")
	}
	return nil
}

func (p *EntityReserved) BeforeUpdate(tx *gorm.DB) (err error) {
	if p.isSystem {
		return nil
	}
	if tx.Statement.Changed("Reserved") {
		return errors.New("reserved not allowed to change")
	}
	if p.Reserved {
		return errors.New("system reserved entity can not be modified")
	}
	return nil
}

func (p *EntityReserved) BeforeDelete(tx *gorm.DB) (err error) {
	if p.isSystem {
		return nil
	}
	if tx.Statement.Changed("Reserved") {
		return errors.New("reserved not allowed to change")
	}
	if p.Reserved {
		return errors.New("system reserved entity can not be deleted")
	}
	return nil
}
