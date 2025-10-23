package entity

import (
	"errors"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/db/model"
)

type OperationIdEnumPairs []OperationIdEnumPair

func (pairs OperationIdEnumPairs) ToMap() map[string]int {
	m := make(map[string]int)
	for _, pair := range pairs {
		m[pair.OperationId] = pair.Enum
	}
	return m
}

type OperationIdEnumPair struct {
	OperationId string `gorm:"<-:create;uniqueIndex;not null;check:operation_id <> '';size:254;comment:Operation ID;"`
	Enum        int    `gorm:"primaryKey;comment:Enum Value;"`
}

// AfterCreate implements model.Setup.
func (p *OperationIdEnumPair) AfterCreate(tx *gorm.DB) (err error) {
	return
}

// AfterDelete implements model.Setup.
func (p *OperationIdEnumPair) AfterDelete(tx *gorm.DB) (err error) {
	return
}

// AfterFind implements model.Setup.
func (p *OperationIdEnumPair) AfterFind(tx *gorm.DB) (err error) {
	return
}

// AfterUpdate implements model.Setup.
func (p *OperationIdEnumPair) AfterUpdate(tx *gorm.DB) (err error) {
	return
}

// BeforeCreate implements model.Setup.
func (p *OperationIdEnumPair) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

// BeforeDelete implements model.Setup.
func (p *OperationIdEnumPair) BeforeDelete(tx *gorm.DB) (err error) {
	return errors.New("operation_id_enum_pairs not allowed to delete")
}

// BeforeUpdate implements model.Setup.
func (p *OperationIdEnumPair) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed("OperationId") {
		return errors.New("operation_id not allowed to change")
	}
	return
}

// Many2ManySetups implements model.Setup.
func (p *OperationIdEnumPair) Many2ManySetups() (setups []model.Many2ManySetup) {
	return
}
