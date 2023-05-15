package options

import (
	"log"
	"my-app/backend/pkg/logger"
	"my-app/backend/pkg/logger/options"
	"my-app/backend/pkg/snowflake"
	iSnowflake "my-app/backend/pkg/snowflake/interfaces"
	"os"
	"time"

	"github.com/imdario/mergo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type OEngine struct {
	Dialector gorm.Dialector
	Options   []gorm.Option
	Snowflake iSnowflake.ISnowflake
	Logger    *OEngineLogger
	Migrate   []any
}

type OEngineLogger struct {
	Writer gormLogger.Writer
	Config gormLogger.Config
}

func DefaultOEngine() *OEngine {
	idGenerator, _ := snowflake.Default()

	return &OEngine{
		Dialector: sqlite.Open("./sqlite.db"),
		Options: []gorm.Option{
			&gorm.Config{},
		},
		Snowflake: idGenerator,
		Logger: &OEngineLogger{
			Writer: logger.NewLogger(&options.OLogger{
				Writer: os.Stderr,
				Tag:    "DBS",
				PrefixTemplate: func(tag string) (prefix string) {
					return "[" + tag + "]"
				},
				Flags: log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile,
			}),
			Config: gormLogger.Config{
				SlowThreshold:             time.Second,
				Colorful:                  true,
				IgnoreRecordNotFoundError: false,
				ParameterizedQueries:      false,
				LogLevel:                  gormLogger.Info,
			},
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
