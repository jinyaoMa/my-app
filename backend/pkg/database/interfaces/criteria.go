package interfaces

type Criteria interface {
	Offset() int
	Limit() int
	OrderBy() []string
	Select() []string
	GroupBy() []string
}
