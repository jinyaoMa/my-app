package crud

type FilterSpecial string

const (
	FilterSpecialIdString  FilterSpecial = "idString"
	FilterSpecialSensitive FilterSpecial = "sensitive"
	FilterSpecialSecret    FilterSpecial = "secret"
	FilterSpecialPassword  FilterSpecial = "password"
	FilterSpecialOid       FilterSpecial = "oid"
)
