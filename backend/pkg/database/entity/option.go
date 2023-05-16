package entity

import iSnowflake "my-app/backend/pkg/snowflake/interfaces"

type Option struct {
	Entity
	Key   string `xorm:"size:256; unique; index"`
	Value string `xorm:"size:256"`
}

func NewOption(snowflake iSnowflake.ISnowflake, option *Option) *Option {
	option.snowflake = snowflake
	return option
}
