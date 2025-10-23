package db

import (
	"fmt"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"majinyao.cn/my-app/backend/pkg/cflog"
	"majinyao.cn/my-app/backend/pkg/db/dbcontext"
	"majinyao.cn/my-app/backend/pkg/db/model"
	"majinyao.cn/my-app/backend/pkg/utils"
)

func Open(models []model.Setup, options Options) (db *gorm.DB, err error) {
	dialector, err := openDialector(options)
	if err != nil {
		return
	}

	cfg, err := newGormConfig(options)
	if err != nil {
		return
	}

	db, err = gorm.Open(dialector, cfg)
	if err != nil {
		return
	}

	err = setContext(db, options)
	if err != nil {
		return nil, err
	}

	for _, m := range models {
		for _, s := range m.Many2ManySetups() {
			err = db.SetupJoinTable(m, s.Field, s.JoinTable)
			if err != nil {
				return nil, err
			}
		}
	}

	if options.AutoMigrate {
		err = db.AutoMigrate(utils.SliceMap(models, func(m model.Setup) any { return m })...)
		if err != nil {
			return nil, err
		}
	}
	return
}

func setContext(db *gorm.DB, options Options) (err error) {
	_, err = dbcontext.SetSnowflake(db, options.Snowflake)
	if err != nil {
		return
	}
	_, err = dbcontext.SetKeygen(db, options.Keygen)
	if err != nil {
		return
	}
	_, err = dbcontext.SetHasher(db, options.Hasher)
	if err != nil {
		return
	}
	_, err = dbcontext.SetCipher(db, options.Cipher)
	if err != nil {
		return
	}
	return
}

func newGormConfig(options Options) (cfg *gorm.Config, err error) {
	cflog, err := cflog.New(options.Cflog)
	if err != nil {
		return
	}

	return &gorm.Config{
		PrepareStmt:    true,
		PrepareStmtTTL: time.Hour,
		Logger: logger.New(cflog, logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  options.LogLevel,
			IgnoreRecordNotFoundError: !options.LogNotFound,
			Colorful:                  options.Cflog.LogFile == "",
		}),
	}, nil
}

func openDialector(options Options) (dialector gorm.Dialector, err error) {
	switch options.Driver {
	case DrvSqlite:
		dialector = sqlite.Open(options.Dsn)
	default:
		err = fmt.Errorf("unsupported db driver %s", options.Driver)
	}
	return
}
