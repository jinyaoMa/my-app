package interfaces

import "my-app/backend/pkg/database/options"

type QueryCondition func(where func(query any, args ...any))

type ICrud[TEntity IEntity] interface {
	Query(criteria *options.OCriteria, condition QueryCondition) (entities []TEntity, err error)
	GetById(id int64) (entity TEntity)
}
