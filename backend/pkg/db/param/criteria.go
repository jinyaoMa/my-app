package param

import (
	"dario.cat/mergo"
)

const (
	OrdAscending CriteriaSortOrder = iota
	OrdDescending
)

type Criteria struct {
	Page   int
	Size   int
	Fields []string
	Sorts  []CriteriaSort
}

type CriteriaSort struct {
	Column string
	Order  CriteriaSortOrder
}

type CriteriaSortOrder int

func (c *Criteria) Offset() int {
	return c.Size * (c.Page - 1)
}

func DefaultCriteria() *Criteria {
	return &Criteria{
		Page:   1,
		Size:   10,
		Fields: []string{},
		Sorts:  []CriteriaSort{},
	}
}

func NewCriteria(dst *Criteria) *Criteria {
	src := DefaultCriteria()

	err := mergo.Merge(dst, *src, mergo.WithAppendSlice)
	if err != nil {
		return src
	}

	return dst
}
