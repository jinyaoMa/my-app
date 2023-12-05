package db

import (
	"log"

	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type DB struct {
	*gorm.DB
	options *DBOptions
}

func (db *DB) Session(cfg *gorm.Session) *DB {
	return &DB{
		DB:      db.DB.Session(cfg),
		options: db.options,
	}
}

func (db *DB) Transaction(fc func(tx *DB) error) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		return fc(&DB{
			DB:      tx,
			options: db.options,
		})
	})
}

func (db *DB) SetupJoinTables(joinTables ...JoinTable) (err error) {
	for _, joinTable := range joinTables {
		if err = db.SetupJoinTable(joinTable.From, joinTable.Field, joinTable.To); err != nil {
			return err
		}
	}
	return
}

func NewDB(logger *log.Logger, options *DBOptions) (db *DB, err error) {
	options, err = NewDBOptions(options)
	if err != nil {
		return nil, err
	}

	db = &DB{
		options: options,
	}

	db.DB, err = gorm.Open(options.GormDialector, options.GormOptions...)
	if err != nil {
		return
	}

	db.Logger = gormLogger.New(logger, options.LoggerConfig)

	err = options.OnInitialized(db)
	if err != nil {
		return
	}

	return
}
