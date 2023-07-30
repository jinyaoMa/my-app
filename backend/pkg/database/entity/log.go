package entity

type Log struct {
	Entity

	/* internal fields */
	Tag     string `gorm:"size:8; default:''"`
	Message string `gorm:"size:4096; default:''"`

	/* relational fields */
}
