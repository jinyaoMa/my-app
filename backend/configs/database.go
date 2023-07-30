package configs

import (
	"my-app/backend/pkg/snowflake"
)

type Database struct {
	LogFile   string
	CipherKey string
	Snowflake *snowflake.Option
}
