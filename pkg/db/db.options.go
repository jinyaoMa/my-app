package db

import (
	"my-app/pkg/base"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBOptions struct {
	base.Options
	GormDialector gorm.Dialector
	GormOptions   []gorm.Option
	OnInitialized func(db *DB) (err error)
	LoggerConfig  logger.Config
}

func DefaultDBOptions() *DBOptions {
	return &DBOptions{
		GormDialector: sqlite.Open("./sqlite.db"),
		GormOptions: []gorm.Option{
			&gorm.Config{},
		},
		LoggerConfig: logger.Config{
			SlowThreshold:             time.Second,
			Colorful:                  true,
			IgnoreRecordNotFoundError: false,
			ParameterizedQueries:      false,
			LogLevel:                  logger.Info,
		},
	}
}

func NewDBOptions(dst *DBOptions) (*DBOptions, error) {
	return base.MergeOptions(DefaultDBOptions(), dst)
}
