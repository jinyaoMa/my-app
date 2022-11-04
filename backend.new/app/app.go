package app

import (
	"context"
	"sync"

	"my-app/backend.new/app/i18n"
	"my-app/backend.new/app/types"

	"gorm.io/gorm"
)

var _app = &app{}

type app struct {
	once sync.Once
	env  *env            // environment variables
	db   *gorm.DB        // database connection
	cfg  *config         // application config
	log  *log            // loggers for whole application
	i18n *i18n.I18n      // languages/translations
	ctx  context.Context // wails context
}

// application global resources and states,
// app initialized as flow: load env -> connect db -> load config -> setup log -> setup i18n,
// application panic instead of logging before log setup
func App() *app {
	_app.once.Do(func() {
		// load env
		_app.env = LoadEnv()

		// connect db
		_app.db = ConnectDatabase()

		// load config
		_app.cfg = LoadConfig(_app.db)

		// setup log
		_app.log = NewConsoleLogger()
		if _app.env.IsLog2File() {
			_app.log = NewFileLogger(_app.cfg.Get(types.ConfigNameLogFile))
		}
		_app.db.Logger = _app.log.database

		// setup i18n
		_app.i18n = i18n.NewI18n(_app.cfg.Get(types.ConfigNameDirLanguages), _app.log.i18n)
	})
	return _app
}

// Env get environment variables
func (a *app) Env() *env {
	return a.env
}

// DB get database connection
func (a *app) DB() *gorm.DB {
	return a.db
}

// Cfg get application config
func (a *app) Cfg() *config {
	return a.cfg
}

// Log get loggers for application
func (a *app) Log() *log {
	return a.log
}

// I18n get i18n
func (a *app) I18n() *i18n.I18n {
	return a.i18n
}

// Ctx get wails context
func (a *app) Ctx() context.Context {
	return a.ctx
}

// SetCtx set wails context
func (a *app) SetCtx(ctx context.Context) *app {
	a.ctx = ctx
	return a
}

// T get current translation
func (a *app) T() *i18n.Translation {
	return a.i18n.Translation(a.i18n.ParseLanguage(a.cfg.Get(types.ConfigNameDisplayLanguage)))
}
