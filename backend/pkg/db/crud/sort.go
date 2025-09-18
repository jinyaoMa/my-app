package crud

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Sort struct {
	Field string `json:"field" doc:"Sorted Field Name"`
	Desc  bool   `json:"desc" required:"false" doc:"Desc or Asc"`
}

func (a *Sort) Apply(tx *gorm.DB) *gorm.DB {
	tx = tx.Order(clause.OrderByColumn{
		Column: clause.Column{Name: a.Field},
		Desc:   a.Desc,
	})
	return tx
}
