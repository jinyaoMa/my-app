package entity

import iSnowflake "my-app/backend/pkg/snowflake/interfaces"

type Option struct {
	Entity
	Key       string `gorm:"size:100; unique; index"`
	Value     string `gorm:"size:255"`
	Encrypted bool   `gorm:""`
}

func NewOption(snowflake iSnowflake.ISnowflake, option *Option) *Option {
	option.snowflake = snowflake
	return option
}
