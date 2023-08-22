package db

import (
	"my-app/pkg/db/param"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	*gorm.DB
	options *Option
}

func (db *DB) SetupJoinTables(joinTables ...param.JoinTable) (err error) {
	for _, joinTable := range joinTables {
		if err = db.SetupJoinTable(joinTable.From, joinTable.Field, joinTable.To); err != nil {
			return err
		}
	}
	return
}

func New(opts *Option) (db *DB, err error) {
	db = &DB{
		options: NewOption(opts),
	}

	db.DB, err = gorm.Open(db.options.Dialector, db.options.Options...)
	if err != nil {
		return
	}

	db.Logger = logger.New(db.options.Logger, db.options.LoggerConfig)

	err = opts.OnInitialized(db)
	if err != nil {
		return
	}
	return
}
