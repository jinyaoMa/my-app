package engine

import (
	"my-app/backend/pkg/snowflake"

	"github.com/imdario/mergo"
)

const (
	DrvSQLite3 string = "sqlite3"
)

type Options struct {
	Driver     string
	DataSource string
	Snowflake  snowflake.ISnowflake
}

func DefaultOptions() *Options {
	return &Options{
		Driver:     DrvSQLite3,
		DataSource: "./" + DrvSQLite3 + ".db",
	}
}

func NewOptions(opts *Options) *Options {
	src := DefaultOptions()

	err := mergo.Merge(opts, *src)
	if err != nil {
		return src
	}

	return opts
}
