package interfaces

import "my-app/backend/pkg/database/options"

type QueryCondition func(where func(query any, args ...any))

type ICrud[TEntity IEntity] interface {
	Query(criteria *options.OCriteria, condition QueryCondition, includes ...string) (entities []TEntity, err error)
	GetById(id int64) (entity TEntity, err error)
	All() (entities []TEntity, err error)
	FindOne(condition QueryCondition) (entity TEntity, err error)
	Save(entity TEntity) (affected int64, err error)
	Delete(id int64) (affected int64, err error)
}
