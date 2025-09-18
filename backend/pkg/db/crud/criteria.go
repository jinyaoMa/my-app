package crud

import (
	"gorm.io/gorm"
)

type Criteria struct {
	Choice
	Size     int      `json:"size" doc:"Page Size"`
	Page     int      `json:"page" doc:"Page Number"`
	Includes []string `json:"includes" required:"false" doc:"Included Associations"`
	Joins    []string `json:"joins" required:"false" doc:"Join 121/belong2 Associations"`
	Filters  []Filter `json:"filters" required:"false" doc:"List Filter Conditions"`
	Sorts    []Sort   `json:"sorts" required:"false" doc:"List Sort Conditions"`
}

func (c *Criteria) Apply(tx *gorm.DB) *gorm.DB {
	for _, include := range c.Includes {
		tx = tx.Preload(include)
	}
	for _, join := range c.Joins {
		tx = tx.Joins(join)
	}
	for _, filter := range c.Filters {
		tx = filter.Apply(tx)
	}
	for _, sort := range c.Sorts {
		tx = sort.Apply(tx)
	}
	tx = tx.Limit(c.limit()).Offset(c.offset())
	tx = c.Choice.Apply(tx)
	return tx
}

func (c *Criteria) limit() int {
	if c.Size < 0 {
		return -1
	}
	return c.Size
}

func (c *Criteria) offset() int {
	if c.Size < 0 || c.Page < 0 {
		return -1
	}
	return c.Size * (c.Page - 1)
}
