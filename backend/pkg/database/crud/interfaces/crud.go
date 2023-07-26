package interfaces

import (
	"my-app/backend/pkg/database/entity/interfaces"
	"my-app/backend/pkg/database/vmodel"
)

type ICrud[TEntity interfaces.IEntity] interface {
	Query(criteria *vmodel.Criteria, condition vmodel.QueryCondition, includes ...string) (entities []TEntity, err error)
	GetById(id int64) (entity TEntity, err error)
	All() (entities []TEntity, err error)
	FindOne(condition vmodel.QueryCondition) (entity TEntity, err error)
	Save(entity TEntity) (affected int64, err error)
	SaveAll(entities []TEntity) (affected int64, err error)
	Delete(id int64) (affected int64, err error)
}
