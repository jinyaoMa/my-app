package entity

type FileCategory struct {
	Entity

	/* internal fields */
	Name string `gorm:"size:512"`

	/* relational fields */
	FileExtensions []*FileExtension `gorm:""`
}
