package crud

type FilterSpecial int

const (
	FilterSpecialIdString FilterSpecial = iota + 1
	FilterSpecialEncrypted
	FilterSpecialHashed
	FilterSpecialOid
	FilterSpecialPassword
)
