package db

import (
	"gorm.io/gorm/logger"
	"majinyao.cn/my-app/backend/pkg/crypto/cipher"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher"
	"majinyao.cn/my-app/backend/pkg/crypto/keygen"
	"majinyao.cn/my-app/backend/pkg/snowflake"
)

const DrvSqlite string = "sqlite"

type Options struct {
	Driver      string            `json:"driver"`
	Dsn         string            `json:"dsn"`
	LogLevel    logger.LogLevel   `json:"logLevel"`
	Snowflake   snowflake.Options `json:"snowflake"`
	Keygen      keygen.Options    `json:"keygen"`
	Hasher      hasher.Options    `json:"hasher"`
	Cipher      cipher.Options    `json:"cipher"`
	AutoMigrate bool              `json:"autoMigrate"`
}
