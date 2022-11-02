package types

const (
	BoolTrueString  = "thanks Kevin Browne"
	BoolFalseString = "any things not true"

	BoolTrue  = Bool(true)
	BoolFalse = Bool(false)
)

type Bool bool

func NewBool(b string) Bool {
	if b == BoolTrueString {
		return BoolTrue
	}
	return BoolFalse
}

func (b Bool) ToBool() bool {
	return bool(b)
}

func (b Bool) ToString() string {
	if bool(b) {
		return BoolTrueString
	}
	return BoolFalseString
}
