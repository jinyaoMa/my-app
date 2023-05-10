package engine

import (
	"io"
	"log"
	"my-app/backend/pkg/snowflake"
	"os"

	"github.com/imdario/mergo"
	xormLog "xorm.io/xorm/log"
)

const (
	DrvSQLite3 string = "sqlite3"
)

type Options struct {
	Driver     string
	DataSource string
	Snowflake  snowflake.ISnowflake
	Logger     *OptionsLogger
}

type OptionsLogger struct {
	Tag            string
	PrefixTemplate func(tag string) string
	Writer         io.Writer
	Flags          int
	LogLevel       xormLog.LogLevel
	ShowSQL        bool
}

func DefaultOptions() *Options {
	idGenerator, _ := snowflake.Default()

	return &Options{
		Driver:     DrvSQLite3,
		DataSource: "./" + DrvSQLite3 + ".db",
		Snowflake:  idGenerator,
		Logger: &OptionsLogger{
			Tag: "DBS",
			PrefixTemplate: func(tag string) string {
				return "[" + tag + "]"
			},
			Writer:   os.Stderr,
			Flags:    log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile,
			LogLevel: xormLog.DEFAULT_LOG_LEVEL,
			ShowSQL:  false,
		},
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
