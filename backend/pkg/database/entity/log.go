package entity

type Log struct {
	Entity

	/* internal fields */
	Message string `gorm:"size:4096; default:''"`

	/* relational fields */
}
