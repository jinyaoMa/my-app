package crud

import (
	"my-app/backend/pkg/database/engine"
	"my-app/backend/pkg/database/entity"
)

type ICrud[T entity.IEntity] interface {
	Delete(id int64) (bool, error)
}

type Crud[T entity.IEntity] struct {
	ICrud[T]
	_engine engine.Engine[T]
	_entity T
}

func New[T entity.IEntity](engine engine.Engine[T], entity T) *Crud[T] {
	return &Crud[T]{
		_engine: engine,
		_entity: entity,
	}
}

func (c *Crud[T]) Delete(id int64) (int64, error) {
	return c._engine.Engine.ID(id).Delete(c._entity)
}
