package interfaces

import (
	"my-app/backend/internal/entity"
	"my-app/backend/pkg/db"
)

type ICRUDFile interface {
	db.ICRUD[*entity.File]
}

type ICRUDFileCategory interface {
	db.ICRUD[*entity.FileCategory]
}

type ICRUDFileExtension interface {
	db.ICRUD[*entity.FileExtension]
}

type ICRUDLog interface {
	db.ICRUD[*entity.Log]
}

type ICRUDNode interface {
	db.ICRUD[*entity.Node]
}

type ICRUDOption interface {
	db.ICRUD[*entity.Option]

	GetByOptionName(name string) (opt *entity.Option, err error)
	GetUint16ByOptionName(name string) (value uint16, opt *entity.Option, err error)
}

type ICRUDUserPassword interface {
	db.ICRUD[*entity.UserPassword]
}

type ICRUDUser interface {
	db.ICRUD[*entity.User]
}
