package entity

type Log struct {
	Entity

	/* internal fields */
	Tag     string `gorm:"size:8; default:''"`
	Code    int64  `gorm:"default:-1"`
	Message string `gorm:"size:4096; default:''"`

	/* relational fields */
}
