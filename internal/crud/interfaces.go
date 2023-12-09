package crud

import (
	"my-app/internal/entity"
	"my-app/pkg/db"
)

type ILog interface {
	db.ICRUD[*entity.Log]
}

type IOption interface {
	db.ICRUD[*entity.Option]
}

type IUserPassword interface {
	db.ICRUD[*entity.UserPassword]
}

type IUser interface {
	db.ICRUD[*entity.User]
}

type IFileCategory interface {
	db.ICRUD[*entity.FileCategory]
}

type IFileExtension interface {
	db.ICRUD[*entity.FileExtension]
}

type IFile interface {
	db.ICRUD[*entity.File]
}
