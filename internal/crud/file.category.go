package crud

import (
	"my-app/internal/entity"
	"my-app/pkg/db"

	"gorm.io/gorm"
)

type FileCategory struct {
	db.ICRUD[*entity.FileCategory]
	session *gorm.DB
}

func NewFileCategory(session *gorm.DB) (fileCategory *FileCategory, iFileCategory IFileCategory) {
	_, crud := db.NewCRUD[*entity.FileCategory](session)
	fileCategory = &FileCategory{
		ICRUD:   crud,
		session: session,
	}
	return fileCategory, fileCategory
}
