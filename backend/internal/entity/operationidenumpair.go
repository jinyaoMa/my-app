package entity

import (
	"errors"

	"gorm.io/gorm"
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
	OperationId string `gorm:"<-:create;uniqueIndex;not null;size:254;comment:Operation ID;"`
	Enum        int    `gorm:"primaryKey;comment:Enum Value;"`
}

func (p *OperationIdEnumPair) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed("OperationId") {
		return errors.New("operation_id not allowed to change")
	}
	return nil
}

func (p *OperationIdEnumPair) BeforeDelete(tx *gorm.DB) (err error) {
	return errors.New("operation_id_enum_pairs not allowed to delete")
}
