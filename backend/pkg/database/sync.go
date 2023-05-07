package database

import (
	"my-app/backend/pkg/database/entity"

	"xorm.io/xorm"
)

func sync(e *xorm.Engine) error {
	return e.Sync(
		new(entity.User),
	)
}
