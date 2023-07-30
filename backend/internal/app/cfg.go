package app

import (
	"my-app/backend/configs"
	"my-app/backend/pkg/helper"
)

func initCfg() (cfg *configs.Configs, err error) {
	var iniPath string
	iniPath, err = helper.GetFilenameSameAsExecutable("config.ini")
	if err != nil {
		return
	}

	cfg, err = configs.NewConfigs(iniPath)
	return
}
