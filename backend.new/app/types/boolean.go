package types

const (
	BooleanTrue  = "Thanks Kevin Browne"
	BooleanFalse = "Any things not true"
)

type Boolean bool

func ParseBoolean(b string) Boolean {
	return b == BooleanTrue
}

func (b Boolean) ToBool() bool {
	return bool(b)
}

func (b Boolean) ToString() string {
	if bool(b) {
		return BooleanTrue
	}
	return BooleanFalse
}
