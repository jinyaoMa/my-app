package options

import (
	"log"
	"my-app/backend/pkg/logger/options"
	"os"
	"time"

	"github.com/imdario/mergo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type ODatabase struct {
	Dialector gorm.Dialector
	Options   []gorm.Option
	Migrate   []any
	Join      []ODatabaseJoin
	Logger    ODatabaseLogger
}

type ODatabaseJoin struct {
	Model     any
	Field     string
	JoinTable any
}

type ODatabaseLogger struct {
	options.OLogger
	Config gormLogger.Config
}

func DefaultODatabase() *ODatabase {
	return &ODatabase{
		Dialector: sqlite.Open("./sqlite.db"),
		Options: []gorm.Option{
			&gorm.Config{},
		},
		Logger: ODatabaseLogger{
			OLogger: options.OLogger{
				Writer: os.Stderr,
				Tag:    "DBS",
				PrefixTemplate: func(tag string) (prefix string) {
					return "[" + tag + "] "
				},
				Flags: log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile,
			},
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

func NewODatabase(dst *ODatabase) *ODatabase {
	src := DefaultODatabase()

	err := mergo.Merge(dst, *src)
	if err != nil {
		return src
	}

	return dst
}
