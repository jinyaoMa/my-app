package engine

import (
	"my-app/backend/pkg/database/entity"

	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

type Engine[T entity.IEntity] struct {
	*xorm.Engine
	*Options
}

func NewEngine(opts *Options) (*Engine[entity.IEntity], error) {
	opts = NewOptions(opts)

	engine, err := xorm.NewEngine(opts.Driver, opts.DataSource)
	if err != nil {
		return nil, err
	}

	logger := log.NewSimpleLogger3(
		opts.Logger.Writer,
		opts.Logger.PrefixTemplate(opts.Logger.Tag),
		opts.Logger.Flags,
		opts.Logger.LogLevel,
	)
	logger.ShowSQL(opts.Logger.ShowSQL)
	engine.SetLogger(logger)

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
