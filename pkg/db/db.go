package db

import (
	"log"

	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func NewDB(logger *log.Logger, options *DBOptions) (db *gorm.DB, err error) {
	options, err = NewDBOptions(options)
	if err != nil {
		return nil, err
	}

	db, err = gorm.Open(options.GormDialector, options.GormOptions...)
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

func SetupJoinTables(db *gorm.DB, joinTables ...JoinTable) (err error) {
	for _, joinTable := range joinTables {
		if err = db.SetupJoinTable(joinTable.From, joinTable.Field, joinTable.To); err != nil {
			return err
		}
	}
	return
}
