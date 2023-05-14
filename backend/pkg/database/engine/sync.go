package engine

import (
	"my-app/backend/pkg/database/entity"

	"xorm.io/xorm"
)

func sync(e *xorm.Engine, beans []interface{}) error {
	beans = append(beans, []interface{}{
		new(entity.Option),
		new(entity.Log),
		new(entity.User),
		new(entity.UserPassword),
	})
	return e.Sync(
		new(entity.Option),
		new(entity.Log),
		new(entity.User),
		new(entity.UserPassword),
	)
}
