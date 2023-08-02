package configs

import (
	"my-app/backend/pkg/helper"
	"my-app/backend/pkg/snowflake"
)

func Default() (cfg *Configs, err error) {
	var assetsPath string
	assetsPath, err = helper.GetPathStartedFromExecutable("Assets")
	if err != nil {
		return
	}
	var languagesPath string
	languagesPath, err = helper.GetPathStartedFromExecutable("Languages")
	if err != nil {
		return
	}
	var key string
	key, err = helper.GetFilenameSameAsExecutable("option.key")
	if err != nil {
		return
	}
	var logFile string
	logFile, err = helper.GetFilenameSameAsExecutable("db.log")
	if err != nil {
		return
	}

	return &Configs{
		AssetsPath:    assetsPath,
		LanguagesPath: languagesPath,
		Database: &Database{
			LogFile:   logFile,
			CipherKey: key,
			Snowflake: snowflake.DefaultOption(),
		},
	}, nil
}