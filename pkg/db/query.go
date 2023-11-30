package db

import "my-app/pkg/base"

type QueryCondition func(where func(query any, args ...any))

type QueryCriteria struct {
	base.Options
	Page    int
	Size    int
	Fields  []string
	Sorts   []QueryCriteriaSort
	Filters []QueryCriteriaFilter
}

type QueryCriteriaSort struct {
	Column string
	Desc   bool
}

type QueryCriteriaFilter struct {
	Condition string
	Params    []any
}

func (queryCriteria *QueryCriteria) Offset() int {
	return queryCriteria.Size * (queryCriteria.Page - 1)
}

func DefaultQueryCriteria() *QueryCriteria {
	return &QueryCriteria{
		Page:    1,
		Size:    10,
		Fields:  []string{},
		Sorts:   []QueryCriteriaSort{},
		Filters: []QueryCriteriaFilter{},
	}
}

func NewQueryCriteria(dst *QueryCriteria) (*QueryCriteria, error) {
	return base.MergeOptions(DefaultQueryCriteria(), dst)
}
