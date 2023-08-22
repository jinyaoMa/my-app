package interfaces

import (
	"my-app/pkg/db/param"
)

type ICRUD[TEntity IEntity] interface {
	Query(criteria *param.Criteria, condition param.QueryCondition, includes ...string) (entities []TEntity, err error)
	GetById(id int64) (entity TEntity, err error)
	All() (entities []TEntity, err error)
	FindOne(condition param.QueryCondition) (entity TEntity, err error)
	Save(entity TEntity) (affected int64, err error)
	SaveAll(entities []TEntity) (affected int64, err error)
	Delete(id int64) (affected int64, err error)
}
