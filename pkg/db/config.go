package db

import (
	"my-app/pkg/enc"
	"my-app/pkg/id"
	"my-app/pkg/log"
	"time"

	"dario.cat/mergo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Dialector     gorm.Dialector
	Options       []gorm.Option
	OnInitialized func(db *DB) (err error)
	Logger        *log.Log
	LoggerConfig  logger.Config
	IdGenerator   id.IID
	DataCipher    enc.ICipher
}

func DefaultConfig() *Config {
	return &Config{
		Dialector: sqlite.Open("./sqlite.db"),
		Options: []gorm.Option{
			&gorm.Config{},
		},
		Logger: log.New(&log.Config{
			Out:    log.NewConsoleLogWriter(),
			Prefix: "[DBS] ",
			Flag:   log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile,
		}),
		LoggerConfig: logger.Config{
			SlowThreshold:             time.Second,
			Colorful:                  true,
			IgnoreRecordNotFoundError: false,
			ParameterizedQueries:      false,
			LogLevel:                  logger.Info,
		},
	}
}

func NewConfig(dst *Config) *Config {
	src := DefaultConfig()
	err := mergo.Merge(dst, *src)
	if err != nil {
		return src
	}
	return dst
}
