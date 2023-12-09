package crud

import (
	"my-app/internal/entity"
	"my-app/pkg/db"
)

type IUser interface {
	db.ICRUD[*entity.User]
}
