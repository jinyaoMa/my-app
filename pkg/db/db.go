package db

import "gorm.io/gorm"

type DB struct {
	*gorm.DB
	options *DBOptions
}
