package interfaces

import "my-app/backend/pkg/database/options"

type QueryCondition func() (query interface{}, args []interface{})

type ICrud[TEntity IEntity] interface {
	Query(criteria *options.OCriteria, conditions ...QueryCondition) (entities []TEntity, err error)
	GetById(id int64) (entity TEntity)
}
