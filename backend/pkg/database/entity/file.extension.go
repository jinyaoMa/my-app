package entity

type FileExtension struct {
	Entity
	Name           string `gorm:"size:256"`
	DotName        string `gorm:"size:64"`
	FileID         int64
	FileCategoryID int64
}
