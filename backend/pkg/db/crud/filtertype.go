package crud

type FilterType int

const (
	FilterTypeEqual FilterType = iota + 1
	FilterTypeNotEqual
	FilterTypeLessThan
	FilterTypeLessThanOrEqual
	FilterTypeGreaterThan
	FilterTypeGreaterThanOrEqual
	FilterTypeLike
	FilterTypeNotLike
	FilterTypeNull
	FilterTypeNotNull
	FilterTypeBetween
	FilterTypeNotBetween
	FilterTypeIn
	FilterTypeNotIn
)
