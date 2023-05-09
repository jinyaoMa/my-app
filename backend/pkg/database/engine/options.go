package engine

import "my-app/backend/pkg/snowflake"

const (
	DrvSQLite3 string = "sqlite3"
)

type Options struct {
	Driver     string
	DataSource string
	Snowflake  *snowflake.Snowflake
}
