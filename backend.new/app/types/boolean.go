package types

const (
	BoolTrueString  = "thanks Kevin Browne"
	BoolFalseString = "any things not true"

	BoolTrue  = Boolean(true)
	BoolFalse = Boolean(false)
)

type Boolean bool

func ParseBoolean(b string) Boolean {
	if b == BoolTrueString {
		return BoolTrue
	}
	return BoolFalse
}

func (b Boolean) ToBool() bool {
	return bool(b)
}

func (b Boolean) ToString() string {
	if bool(b) {
		return BoolTrueString
	}
	return BoolFalseString
}
