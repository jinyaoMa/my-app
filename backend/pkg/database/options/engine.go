package options

import (
	"io"
	"log"
	"my-app/backend/pkg/snowflake"
	iSnowflake "my-app/backend/pkg/snowflake/interfaces"
	"os"

	"github.com/imdario/mergo"
	xormLog "xorm.io/xorm/log"
)

const (
	DrvSQLite3 string = "sqlite3"
)

type OEngine struct {
	Driver     string
	DataSource string
	Snowflake  iSnowflake.ISnowflake
	Logger     *OEngineLogger
	Sync       []interface{}
}

type OEngineLogger struct {
	Tag            string
	PrefixTemplate func(tag string) string
	Writer         io.Writer
	Flags          int
	LogLevel       xormLog.LogLevel
	ShowSQL        bool
}

func DefaultOEngine() *OEngine {
	idGenerator, _ := snowflake.Default()

	return &OEngine{
		Driver:     DrvSQLite3,
		DataSource: "./" + DrvSQLite3 + ".db",
		Snowflake:  idGenerator,
		Logger: &OEngineLogger{
			Tag: "DBS",
			PrefixTemplate: func(tag string) string {
				return "[" + tag + "] "
			},
			Writer:   os.Stderr,
			Flags:    log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile,
			LogLevel: xormLog.DEFAULT_LOG_LEVEL,
			ShowSQL:  false,
		},
	}
}

func NewOEngine(dst *OEngine) *OEngine {
	src := DefaultOEngine()

	err := mergo.Merge(dst, *src)
	if err != nil {
		return src
	}

	return dst
}
