package configs

import (
	"my-app/backend/pkg/snowflake"
)

type Database struct {
	CipherKey string
	Snowflake *snowflake.Option
}
