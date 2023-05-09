package crud

import (
	"my-app/backend/pkg/database/engine"
	"my-app/backend/pkg/database/entity"
)

type ICrud[T entity.IEntity] interface {
	GetById(id int64) (entity *T, has bool, err error)
	List() (entity []T, count int64, err error)
	Delete(id int64) (affected int64, err error)
}

type Crud[T entity.IEntity] struct {
	_engine *engine.Engine[T]
	_entity T
}

// Delete implements ICrud
func (c *Crud[T]) Delete(id int64) (affected int64, err error) {
	affected, err = c._engine.Engine.ID(id).Delete(c._entity)
	return
}

// List implements ICrud
func (c *Crud[T]) List() (entity []T, count int64, err error) {
	count, err = c._engine.Engine.FindAndCount(&entity)
	return
}

// GetById implements ICrud
func (c *Crud[T]) GetById(id int64) (entity *T, has bool, err error) {
	has, err = c._engine.Engine.ID(id).Get(entity)
	return
}

func New[T entity.IEntity](engine *engine.Engine[T], entity T) ICrud[T] {
	return &Crud[T]{
		_engine: engine,
		_entity: entity,
	}
}
