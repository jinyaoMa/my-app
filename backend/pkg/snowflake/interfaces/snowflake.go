package interfaces

import (
	"database/sql"
	"database/sql/driver"
)

type ISnowflake interface {
	sql.Scanner
	driver.Valuer
	// Generate creates and returns a unique snowflake ID
	// To help guarantee uniqueness
	// - Make sure your system is keeping accurate system time
	// - Make sure you never have multiple nodes running with the same node ID
	Generate() int64
}
