package app

import (
	"context"
	"time"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/api"
	"majinyao.cn/my-app/backend/pkg/cflog"
	"majinyao.cn/my-app/backend/pkg/config"
	"majinyao.cn/my-app/backend/pkg/executable"
	"majinyao.cn/my-app/backend/pkg/h3server"
	"majinyao.cn/my-app/backend/pkg/i18n"
	"majinyao.cn/my-app/backend/pkg/reactive"
	"majinyao.cn/my-app/backend/pkg/router"
	"majinyao.cn/my-app/backend/pkg/storage"
)

var (
	LOG   *cflog.Cflog
	EXE   executable.IExecutable
	CFG   config.IConfig[Config]
	STORE storage.IStorage
	I18N  i18n.II18n
	DB    *gorm.DB
	API   router.IRouter
	H3S   h3server.IH3Server
	THEME reactive.IReactive[string]
)

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	exe := initEXE()
	cfg := initCFG(exe)
	log := initLOG(cfg.Get().Cflog)
	store := initSTORE(log, cfg.Get().Storage)
	i19 := initI18N(log, cfg.Get().I18n)
	tx := initDB(log, cfg.Get().Db)
	api := initAPI(ctx, log, tx, cfg.Get().Api)
	h3s := initH3S(api)
	setup(ctx, exe, cfg, log, store, i19, tx, api, h3s)
}

func initAPI(ctx context.Context, log *cflog.Cflog, db *gorm.DB, options api.Options) router.IRouter {
	var err error
	API, err = api.New(ctx, db, options)
	if err != nil {
		log.Panicln("init api failed", err)
	}
	return API
}

func initI18N(log *cflog.Cflog, options i18n.Options) i18n.II18n {
	var err error
	I18N, err = i18n.New(options)
	if err != nil {
		log.Panicln("init i18n failed", err)
	}
	return I18N
}

func initSTORE(log *cflog.Cflog, options storage.Options) storage.IStorage {
	var err error
	STORE, err = storage.New(options)
	if err != nil {
		log.Panicln("init store failed", err)
	}
	return STORE
}

func initLOG(options cflog.Options) *cflog.Cflog {
	var err error
	LOG, err = cflog.New(options)
	if err != nil {
		panic(err)
	}
	return LOG
}

func initEXE() executable.IExecutable {
	var err error
	EXE, err = executable.New()
	if err != nil {
		panic(err)
	}
	return EXE
}
