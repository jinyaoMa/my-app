package crud

type FilterType string

const (
	FilterTypeEqual              FilterType = "equal"
	FilterTypeNotEqual           FilterType = "notEqual"
	FilterTypeLessThan           FilterType = "lessThan"
	FilterTypeLessThanOrEqual    FilterType = "lessThanOrEqual"
	FilterTypeGreaterThan        FilterType = "greaterThan"
	FilterTypeGreaterThanOrEqual FilterType = "greaterThanOrEqual"
	FilterTypeContains           FilterType = "contains"
	FilterTypeNotContains        FilterType = "notContains"
	FilterTypeStartsWith         FilterType = "startsWith"
	FilterTypeNotStartsWith      FilterType = "notStartsWith"
	FilterTypeEndsWith           FilterType = "endsWith"
	FilterTypeNotEndsWith        FilterType = "notEndsWith"
	FilterTypeNull               FilterType = "null"
	FilterTypeNotNull            FilterType = "notNull"
	FilterTypeBetween            FilterType = "between"
	FilterTypeNotBetween         FilterType = "notBetween"
	FilterTypeIn                 FilterType = "in"
	FilterTypeNotIn              FilterType = "notIn"
)
