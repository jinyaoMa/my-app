package database

import (
	"my-app/backend/pkg/database/options"
	"my-app/backend/pkg/logger"

	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type Database struct {
	gorm.DB
	options *options.ODatabase
}

func New(opts *options.ODatabase) (*Database, error) {
	opts = options.NewODatabase(opts)

	db, err := gorm.Open(opts.Dialector, opts.Options...)
	if err != nil {
		return nil, err
	}

	db.Logger = gormLogger.New(logger.New(&opts.Logger.Option), opts.Logger.Config)

	opts.OnInitialized(db)

	err = migrate(db, opts.Migrate...)
	if err != nil {
		return nil, err
	}

	err = join(db, opts.Join...)
	if err != nil {
		return nil, err
	}

	return &Database{
		DB:      *db,
		options: opts,
	}, nil
}
