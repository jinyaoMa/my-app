package db

import (
	"my-app/pkg/log"

	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
	options *Option
}

func New(opts *Option) (*DB, error) {
	opts = NewOption(opts)

	db, err := gorm.Open(opts.Dialector, opts.Options...)
	if err != nil {
		return nil, err
	}

	db.Logger = log.Gorm(opts.Logger.Option, opts.Logger.Config)

	opts.OnInitialized(db)

	err = migrate(db, opts.Migrate...)
	if err != nil {
		return nil, err
	}

	err = join(db, opts.Join...)
	if err != nil {
		return nil, err
	}

	return &DB{
		DB:      db,
		options: opts,
	}, nil
}
