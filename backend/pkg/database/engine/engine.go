package engine

import (
	"my-app/backend/pkg/database/interfaces"
	"my-app/backend/pkg/database/options"

	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

type Engine[TEntity interfaces.IEntity] struct {
	*xorm.Engine
	*options.OEngine
}

func NewEngine(opts *options.OEngine) (*Engine[interfaces.IEntity], error) {
	opts = options.NewOEngine(opts)

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

	err = sync(engine, opts.Sync)
	if err != nil {
		return nil, err
	}

	return &Engine[interfaces.IEntity]{
		Engine:  engine,
		OEngine: opts,
	}, nil
}

func (e *Engine[TEntity]) NewEntity(entity TEntity) TEntity {
	entity.SetSnowflake(e.Snowflake)
	return entity
}

type SessionCallback func(session *xorm.Session) error

func (e *Engine[TEntity]) HandleSession(callback SessionCallback) (err error) {
	session := e.NewSession()
	defer session.Close()

	if err = session.Begin(); err != nil {
		return
	}

	if err = callback(session); err != nil {
		return
	}

	err = session.Commit()
	return
}
