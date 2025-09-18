package crud

import "gorm.io/gorm"

type Choice struct {
	Selects []string `json:"selects" required:"false" doc:"Selected Fields"`
	Omits   []string `json:"omits" required:"false" doc:"Omitted Fields"`
}

func (c *Choice) Apply(tx *gorm.DB) *gorm.DB {
	if len(c.Selects) > 0 {
		tx = tx.Select(c.Selects)
	}
	if len(c.Omits) > 0 {
		tx = tx.Omit(c.Omits...)
	}
	return tx
}
