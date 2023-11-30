package db

type JoinTable struct {
	From  any    // the end table/entity of a many-to-many relationship
	Field string // the relevent field of the end entity
	To    any    // the centered table/entity that the end table/entity connected to
}
