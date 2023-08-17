package db

import (
	"my-app/redo/log"
	"time"

	"dario.cat/mergo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Option struct {
	Dialector     gorm.Dialector
	OnInitialized func(db *gorm.DB)
	Options       []gorm.Option
	Migrate       []interface{}
	Join          []OptionJoin
	Logger        OptionLogger
}

type OptionJoin struct {
	Model     any
	Field     string
	JoinTable any
}

type OptionLogger struct {
	log.Option
	Config log.GormConfig
}

func DefaultOption() *Option {
	return &Option{
		Dialector: sqlite.Open("./sqlite.db"),
		Options: []gorm.Option{
			&gorm.Config{},
		},
		Logger: OptionLogger{
			Option: log.Option{
				Out:    log.NewConsoleLogWriter(),
				Prefix: "[DBS] ",
				Flag:   log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile,
			},
			Config: log.GormConfig{
				SlowThreshold:             time.Second,
				Colorful:                  true,
				IgnoreRecordNotFoundError: false,
				ParameterizedQueries:      false,
				LogLevel:                  log.Info,
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
