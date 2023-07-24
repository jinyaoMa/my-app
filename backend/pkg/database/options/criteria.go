package options

import (
	"dario.cat/mergo"
)

const (
	OrdAscending OCriteriaSortOrder = iota
	OrdDescending
)

type OCriteria struct {
	Page   int
	Size   int
	Fields []string
	Sorts  []OCriteriaSort
}

type OCriteriaSort struct {
	Column string
	Order  OCriteriaSortOrder
}

type OCriteriaSortOrder int

func (c *OCriteria) Offset() int {
	return c.Size * (c.Page - 1)
}

func DefaultOCriteria() *OCriteria {
	return &OCriteria{
		Page:   1,
		Size:   10,
		Fields: []string{},
		Sorts:  []OCriteriaSort{},
	}
}

func NewOCriteria(dst *OCriteria) *OCriteria {
	src := DefaultOCriteria()

	err := mergo.Merge(dst, *src, mergo.WithAppendSlice)
	if err != nil {
		return src
	}

	return dst
}
