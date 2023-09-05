package app

import (
	"my-app/backend/configs"
	"my-app/backend/pkg/funcs"
)

func initCfg() (cfg *configs.Configs, err error) {
	var iniPath string
	iniPath, err = funcs.GetFilenameSameAsExecutable("ini")
	if err != nil {
		return
	}

	cfg, err = configs.NewConfigs(iniPath)
	return
}
