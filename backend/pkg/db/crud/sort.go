package crud

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"majinyao.cn/my-app/backend/pkg/utils"
)

type Sorts []Sort

func (s Sorts) Apply(tx *gorm.DB, maxLength int) *gorm.DB {
	length := len(s)
	if length == 0 {
		return tx
	}
	if length > 0 && length > maxLength {
		return tx
	}
	return tx.Order(clause.OrderBy{
		Columns: utils.SliceMap(s, func(sort Sort) clause.OrderByColumn {
			return clause.OrderByColumn{Column: clause.Column{Name: sort.Field}, Desc: sort.Desc}
		}),
	})
}

type Sort struct {
	Field string `json:"field" required:"true" doc:"Sorted Field Name"`
	Desc  bool   `json:"desc" required:"false" doc:"Desc or Asc"`
}

func (s *Sort) Apply(tx *gorm.DB) *gorm.DB {
	return tx.Order(clause.OrderByColumn{Column: clause.Column{Name: s.Field}, Desc: s.Desc})
}
