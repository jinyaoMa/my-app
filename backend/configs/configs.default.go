package configs

import (
	"my-app/backend/pkg/funcs"
	"my-app/backend/pkg/id"
)

func Default() (cfg *Configs, err error) {
	var assetsPath string
	assetsPath, err = funcs.GetPathStartedFromExecutable("Assets")
	if err != nil {
		return
	}
	var languagesPath string
	languagesPath, err = funcs.GetPathStartedFromExecutable("Languages")
	if err != nil {
		return
	}
	var key string
	key, err = funcs.GetFilenameSameAsExecutable("option.key")
	if err != nil {
		return
	}
	var logFile string
	logFile, err = funcs.GetFilenameSameAsExecutable("db.log")
	if err != nil {
		return
	}

	return &Configs{
		AssetsPath:    assetsPath,
		LanguagesPath: languagesPath,
		Database: &Database{
			LogFile:   logFile,
			CipherKey: key,
			Snowflake: id.DefaultConfig(),
		},
	}, nil
}
