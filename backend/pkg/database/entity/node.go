package entity

type Node struct {
	Entity

	/* internal fields */
	Name string `gorm:"size:64; default:''"`
	UUID string `gorm:""`

	/* relational fields */
}
