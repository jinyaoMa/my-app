package crud

import (
	"slices"

	"gorm.io/gorm"
)

type Choices []Choice

func (c Choices) Apply(tx *gorm.DB) *gorm.DB {
	for _, choice := range c {
		tx = choice.Apply(tx)
	}
	return tx
}

type Choice struct {
	Selects []string `json:"selects" required:"false" doc:"Selected Fields"`
	Omits   []string `json:"omits" required:"false" doc:"Omitted Fields"`
}

func (c *Choice) Apply(tx *gorm.DB) *gorm.DB {
	if len(c.Selects) > 0 {
		tx = tx.Select(slices.DeleteFunc(c.Selects, func(s string) bool {
			return !regField.MatchString(s)
		}))
	}
	if len(c.Omits) > 0 {
		tx = tx.Omit(slices.DeleteFunc(c.Omits, func(s string) bool {
			return !regField.MatchString(s)
		})...)
	}
	return tx
}
