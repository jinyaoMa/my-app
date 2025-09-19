package app

import (
	"path/filepath"
	"runtime"
	"time"

	"gorm.io/gorm/logger"
	"majinyao.cn/my-app/backend/api"
	"majinyao.cn/my-app/backend/pkg/api/middlewares/authfwt"
	"majinyao.cn/my-app/backend/pkg/cflog"
	"majinyao.cn/my-app/backend/pkg/codegen"
	"majinyao.cn/my-app/backend/pkg/config"
	"majinyao.cn/my-app/backend/pkg/crypto/cipher"
	"majinyao.cn/my-app/backend/pkg/crypto/cipher/sm4"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/crc64"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/sha3"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/sm3"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/xxh3"
	"majinyao.cn/my-app/backend/pkg/crypto/keygen"
	"majinyao.cn/my-app/backend/pkg/crypto/keygen/argon2"
	"majinyao.cn/my-app/backend/pkg/db"
	"majinyao.cn/my-app/backend/pkg/executable"
	"majinyao.cn/my-app/backend/pkg/fwt"
	"majinyao.cn/my-app/backend/pkg/i18n"
	"majinyao.cn/my-app/backend/pkg/router"
	"majinyao.cn/my-app/backend/pkg/snowflake"
	"majinyao.cn/my-app/backend/pkg/storage"
)

type Config struct {
	Cflog   cflog.Options   `json:"cflog"`
	Storage storage.Options `json:"storage"`
	I18n    i18n.Options    `json:"i18n"`
	Db      db.Options      `json:"db"`
	Api     api.Options     `json:"api"`
}

func defaultConfig(exe executable.IExecutable) Config {
	now := time.Now()
	threads := runtime.NumCPU() / 2
	if threads == 0 {
		threads = 1
	}

	cg, err := codegen.New(codegen.DefaultOptions())
	if err != nil {
		panic(err)
	}

	kg32, err := keygen.New(keygen.Options{
		Alg:       argon2.Alg,
		Salt:      "ml+wqq",
		Threads:   threads,
		KeyLength: 32,
	})
	if err != nil {
		panic(err)
	}

	kg16, err := keygen.New(keygen.Options{
		Alg:       argon2.Alg,
		Salt:      "lfw+zjj",
		Threads:   threads,
		KeyLength: 16,
	})
	if err != nil {
		panic(err)
	}

	kg12, err := keygen.New(keygen.Options{
		Alg:       argon2.Alg,
		Salt:      "mjy+wxb",
		Threads:   threads,
		KeyLength: 12,
	})
	if err != nil {
		panic(err)
	}

	languagesDir := exe.JoinDir("languages")
	libraryDir := exe.JoinDir("library")
	staticsDir := exe.JoinDir("statics")

	return Config{
		Cflog: cflog.Options{
			EnableConsole: true,
			LogFile:       exe.JoinDir("app.log"),
			LogPrefix:     "[APP] ",
		},
		Storage: storage.Options{
			Libraries: []storage.Library{
				{
					Mountpoint: filepath.VolumeName(libraryDir),
					Directory:  libraryDir,
				},
			},
			Temporary:  "tmp",
			BufferSize: 512 * storage.KB,
			Hashers: []hasher.Options{
				{
					Alg:       sha3.Alg,
					BitLength: 256,
				},
				{
					Alg: crc64.Alg,
				},
				{
					Alg: xxh3.Alg,
				},
			},
			MaxPathLength: 254,
		},
		I18n: i18n.Options{
			Fallback:   "zh-CN",
			Directory:  languagesDir,
			DefineJson: "define.json",
		},
		Db: db.Options{
			Cflog: cflog.Options{
				EnableConsole: true,
				LogFile:       exe.JoinDir("db.log"),
				LogPrefix:     "[DB_] ",
			},
			LogLevel: logger.Error,
			Driver:   db.DrvSqlite,
			Dsn:      exe.GetPathWithExt("db") + "?_pragma=foreign_keys(1)",
			Snowflake: snowflake.Options{
				Epoch:    now,
				NodeBits: 7,
				StepBits: 14,
				NodeId:   1,
			},
			Keygen: keygen.Options{
				Alg:       argon2.Alg,
				Salt:      kg16.DeriveBase64(cg.Generate(24)),
				Threads:   threads,
				KeyLength: 32,
			},
			Hasher: hasher.Options{
				Alg: sm3.Alg,
				Key: kg16.DeriveBase64(cg.Generate(24)),
			},
			Cipher: cipher.Options{
				Alg: sm4.Alg,
				Key: kg16.DeriveBase64(cg.Generate(24)),
				Iv:  kg12.DeriveBase64(cg.Generate(18)),
			},
			AutoMigrate: true,
		},
		Api: api.Options{
			Router: router.Options{
				DocsPath:    "/apid",
				DocsTitle:   "MY APP",
				DocsVersion: "0.0.0",
				Cflog: cflog.Options{
					EnableConsole: true,
					LogFile:       exe.JoinDir("api.log"),
					LogPrefix:     "[API] ",
				},
				StaticsDirectory: staticsDir,
				EnableTimeout:    true,
				Timeout:          30,
				EnableCors:       true,
				AllowedOrigins:   []string{"https://*"},
				AllowedMethods:   []string{"GET", "POST", "DELETE"},
				AllowedHeaders:   []string{"*"},
				EnableHttpRate:   true,
				LimitByIp:        true,
				RateLimit:        100,
			},
			AuthFwt: authfwt.Options{
				Fwt: fwt.Options{
					Snowflake: snowflake.Options{
						Epoch:    now,
						NodeBits: 2,
						StepBits: 18,
						NodeId:   1,
					},
					Hasher: hasher.Options{
						Alg:       sha3.Alg,
						Salt:      "mjy+wqq",
						Key:       kg32.DeriveBase64(cg.Generate(48)),
						BitLength: 512,
					},
					Codegen:       codegen.DefaultOptions(),
					Issuer:        "mjy",
					Subject:       "my-app",
					Epoch:         now,
					ExpiredAge:    7200,
					RefreshAge:    864000,
					RefreshLength: 32,
				},
				CacheLimit: 100,
			},
		},
	}
}

func initCFG(exe executable.IExecutable) config.IConfig[Config] {
	var err error
	CFG, err = config.New(config.Options[Config]{
		Path:    exe.GetPathWithExt("json"),
		Default: defaultConfig(exe),
	})
	if err != nil {
		panic(err)
	}
	return CFG
}
