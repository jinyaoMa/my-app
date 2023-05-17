package entity

import (
	iSnowflake "my-app/backend/pkg/snowflake/interfaces"
	iUtility "my-app/backend/pkg/utility/interfaces"
)

var (
	snowflake iSnowflake.ISnowflake
	aes       iUtility.IAes
)

func SetSnowflake(snowflake_ iSnowflake.ISnowflake) {
	snowflake = snowflake_
}

func SetAes(aes_ iUtility.IAes) {
	aes = aes_
}
