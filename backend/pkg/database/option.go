package database

import (
	"log"
	"my-app/backend/pkg/logger"
	"os"
	"time"

	"dario.cat/mergo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type Option struct {
	Dialector     gorm.Dialector
	OnInitialized func(db *gorm.DB)
	Options       []gorm.Option
	Migrate       []any
	Join          []OptionJoin
	Logger        OptionLogger
}

type OptionJoin struct {
	Model     any
	Field     string
	JoinTable any
}

type OptionLogger struct {
	logger.Option
	Config gormLogger.Config
}

func DefaultOption() *Option {
	return &Option{
		Dialector: sqlite.Open("./sqlite.db"),
		Options: []gorm.Option{
			&gorm.Config{},
		},
		Logger: OptionLogger{
			Option: logger.Option{
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

func NewOption(dst *Option) *Option {
	src := DefaultOption()

	err := mergo.Merge(dst, *src)
	if err != nil {
		return src
	}

	return dst
}
