package crud

import (
	"my-app/internal/entity"
	"my-app/pkg/db"

	"gorm.io/gorm"
)

type FileExtension struct {
	db.ICRUD[*entity.FileExtension]
	session *gorm.DB
}

func NewFileExtension(session *gorm.DB) (fileExtension *FileExtension, iFileExtension IFileExtension) {
	_, crud := db.NewCRUD[*entity.FileExtension](session)
	fileExtension = &FileExtension{
		ICRUD:   crud,
		session: session,
	}
	return fileExtension, fileExtension
}
