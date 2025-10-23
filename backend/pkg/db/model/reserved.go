package model

import (
	"database/sql"
	"errors"

	"gorm.io/gorm"
)

type ModelReserved struct {
	Reserved sql.NullBool `gorm:"<-:create;index;comment:Is System Reserved;"`
	isSystem bool         `gorm:"-"`
}

func (r *ModelReserved) IsReserved() bool {
	return r.Reserved.Valid && r.Reserved.Bool
}

func (r *ModelReserved) SysOp(reserved bool) {
	r.Reserved = sql.NullBool{Bool: reserved, Valid: true}
	r.isSystem = true
}

func (r *ModelReserved) BeforeCreate(tx *gorm.DB) (err error) {
	if r.isSystem {
		return
	}
	if r.IsReserved() {
		return errors.New("entity can not be system reserved when operation is not by system")
	}
	return
}

func (r *ModelReserved) BeforeUpdate(tx *gorm.DB) (err error) {
	if r.isSystem {
		return
	}
	if r.IsReserved() {
		return errors.New("system reserved entity can not be modified when operation is not by system")
	}
	return
}

func (r *ModelReserved) BeforeDelete(tx *gorm.DB) (err error) {
	if r.isSystem {
		return
	}
	if r.IsReserved() {
		return errors.New("system reserved entity can not be deleted when operation is not by system")
	}
	return
}
