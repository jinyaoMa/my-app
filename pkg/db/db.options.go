package db

import (
	"my-app/pkg/base"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBOptions struct {
	base.Options
	GormDialector     gorm.Dialector
	GormOptions       []gorm.Option
	OnInitialized func(db *DB) (err error)
}

func DefaultDBOptions() *DBOptions {
	return &DBOptions{
		GormDialector: sqlite.Open("./sqlite.db"),
		GormOptions: []gorm.Option{
			&gorm.Config{},
		},
	}
}

func NewDBOptions(dst *DBOptions) (*DBOptions, error) {
	return base.SimpleMerge(DefaultDBOptions(), dst)
}
