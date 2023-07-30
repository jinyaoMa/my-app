package configs

import "gopkg.in/ini.v1"

type Configs struct {
	AssetsPath    string
	LanguagesPath string
	Database      *Database
}

func NewConfigs(path string) (cfg *Configs, err error) {
	var iniFile *ini.File
	iniFile, err = ini.LooseLoad(path)
	if err != nil {
		return nil, err
	}

	cfg, err = Default()
	if err != nil {
		return nil, err
	}

	err = iniFile.MapTo(cfg)
	if err != nil {
		return nil, err
	}

	err = iniFile.ReflectFrom(cfg)
	if err != nil {
		return nil, err
	}

	return
}
