package db

import (
	"gorm.io/gorm/logger"
	"majinyao.cn/my-app/backend/pkg/cflog"
	"majinyao.cn/my-app/backend/pkg/crypto/cipher"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher"
	"majinyao.cn/my-app/backend/pkg/crypto/keygen"
	"majinyao.cn/my-app/backend/pkg/snowflake"
)

const DrvSqlite string = "sqlite"

type Options struct {
	Cflog       cflog.Options     `json:"cflog"`
	LogLevel    logger.LogLevel   `json:"logLevel"`
	Driver      string            `json:"driver"`
	Dsn         string            `json:"dsn"`
	Snowflake   snowflake.Options `json:"snowflake"`
	Keygen      keygen.Options    `json:"keygen"`
	Hasher      hasher.Options    `json:"hasher"`
	Cipher      cipher.Options    `json:"cipher"`
	AutoMigrate bool              `json:"autoMigrate"`
}
