package model

import (
	"errors"

	"gorm.io/gorm"
)

type Reserved struct {
	isSystem bool `gorm:"-"`
	Reserved bool `gorm:"<-:create;index;comment:Is System Reserved;"`
}

func (r *Reserved) SysOp(reserved bool) {
	r.isSystem = true
	r.Reserved = reserved
}

func (r *Reserved) BeforeCreate(tx *gorm.DB) (err error) {
	if r.isSystem {
		return nil
	}
	if r.Reserved {
		return errors.New("entity can not be system reserved when operation is not by system")
	}
	return nil
}

func (r *Reserved) BeforeUpdate(tx *gorm.DB) (err error) {
	if r.isSystem {
		return nil
	}
	if tx.Statement.Changed("Reserved") {
		return errors.New("reserved not allowed to change")
	}
	if r.Reserved {
		return errors.New("system reserved entity can not be modified")
	}
	return nil
}

func (r *Reserved) BeforeDelete(tx *gorm.DB) (err error) {
	if r.isSystem {
		return nil
	}
	if tx.Statement.Changed("Reserved") {
		return errors.New("reserved not allowed to change")
	}
	if r.Reserved {
		return errors.New("system reserved entity can not be deleted")
	}
	return nil
}
