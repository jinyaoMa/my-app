package entity

import iSnowflake "my-app/backend/pkg/snowflake/interfaces"

type Log struct {
	Entity
	Tag     string `gorm:"size:3"`
	Code    int64  `gorm:""`
	Message string `gorm:"size:2048"`
}

func NewLog(snowflake iSnowflake.ISnowflake, log *Log) *Log {
	log.snowflake = snowflake
	return log
}
