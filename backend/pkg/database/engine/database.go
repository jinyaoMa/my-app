package engine

import (
	"my-app/backend/pkg/database/entity"

	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

type Engine[T entity.IEntity] struct {
	*xorm.Engine
	*Options
}

func New(opts *Options) (*Engine[entity.IEntity], error) {
	engine, err := xorm.NewEngine(opts.Driver, opts.DataSource)
	if err != nil {
		return nil, err
	}

	err = sync(engine)
	if err != nil {
		return nil, err
	}

	return &Engine[entity.IEntity]{
		Engine:  engine,
		Options: opts,
	}, nil
}

func (e *Engine[T]) NewEntity(entity T) T {
	entity.SetSnowflake(e.Snowflake)
	return entity
}
