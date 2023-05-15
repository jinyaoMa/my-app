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

func NewDatabase(opts *options.ODatabase) (*Database, error) {
	opts = options.NewODatabase(opts)

	db, err := gorm.Open(opts.Dialector, opts.Options...)
	if err != nil {
		return nil, err
	}

	db.Logger = gormLogger.New(logger.NewLogger(&opts.Logger.OLogger), opts.Logger.Config)

	err = migrate(db, opts.Migrate...)
	if err != nil {
		return nil, err
	}

	return &Database{
		DB:      *db,
		options: opts,
	}, nil
}
