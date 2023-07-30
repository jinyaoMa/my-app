package configs

import (
	"my-app/backend/pkg/helper"
	"my-app/backend/pkg/snowflake"
)

func Default() (cfg *Configs, err error) {
	var key string
	key, err = helper.GetFilenameSameAsExecutable("option.key")
	if err != nil {
		panic(err)
	}

	return &Configs{
		Database: &Database{
			CipherKey: key,
			Snowflake: snowflake.DefaultOption(),
		},
	}, nil
}
