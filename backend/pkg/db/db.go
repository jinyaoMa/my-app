package db

import (
	"my-app/backend/pkg/db/param"
	"my-app/backend/pkg/enc"
	"my-app/backend/pkg/id"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	config *Config
	*gorm.DB
}

func (db *DB) SetupJoinTables(joinTables ...param.JoinTable) (err error) {
	for _, joinTable := range joinTables {
		if err = db.SetupJoinTable(joinTable.From, joinTable.Field, joinTable.To); err != nil {
			return err
		}
	}
	return
}

func New(cfg *Config) (db *DB, err error) {
	db = &DB{
		config: NewConfig(cfg),
	}

	db.DB, err = gorm.Open(db.config.Dialector, db.config.Options...)
	if err != nil {
		return
	}

	db.Logger = logger.New(db.config.Logger, db.config.LoggerConfig)

	err = db.config.OnInitialized(db)
	if err != nil {
		return
	}

	if db.config.IdGenerator == nil {
		if db.config.IdGenerator, err = id.Default(); err != nil {
			return
		}
	}

	if db.config.DataCipher == nil {
		db.config.DataCipher = enc.NewAesWithSalt("")
	}

	return
}
