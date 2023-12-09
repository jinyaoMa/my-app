package crud

import (
	"my-app/internal/entity"
	"my-app/pkg/db"

	"gorm.io/gorm"
)

type File struct {
	db.ICRUD[*entity.File]
	session *gorm.DB
}

func NewFile(session *gorm.DB) (file *File, iFile IFile) {
	_, crud := db.NewCRUD[*entity.File](session)
	file = &File{
		ICRUD:   crud,
		session: session,
	}
	return file, file
}
