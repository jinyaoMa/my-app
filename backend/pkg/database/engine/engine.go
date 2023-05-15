package engine

import (
	"my-app/backend/pkg/database/interfaces"
	"my-app/backend/pkg/database/options"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Engine[TEntity interfaces.IEntity] struct {
	*gorm.DB
	*options.OEngine
}

func NewEngine(opts *options.OEngine) (*Engine[interfaces.IEntity], error) {
	opts = options.NewOEngine(opts)

	db, err := gorm.Open(opts.Dialector, opts.Options...)
	if err != nil {
		return nil, err
	}

	db.Logger = logger.New(opts.Logger.Writer, opts.Logger.Config)

	err = migrate(db, opts.Migrate...)
	if err != nil {
		return nil, err
	}

	return &Engine[interfaces.IEntity]{
		DB:      db,
		OEngine: opts,
	}, nil
}

func (e *Engine[TEntity]) NewEntity(entity TEntity) TEntity {
	entity.SetSnowflake(e.Snowflake)
	return entity
}
