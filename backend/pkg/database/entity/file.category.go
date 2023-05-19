package entity

type FileCategory struct {
	Entity
	Name       string `gorm:"size:512"`
	Extensions []FileExtension
}
