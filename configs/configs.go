package configs

import "gopkg.in/ini.v1"

// use absolute path for all path values
type Configs struct {
	Dev bool
	Crypto
	Paths
	FileStore
}

func DefaultConfigs() (cfg *Configs, err error) {
	return &Configs{}, nil
}

func NewConfigs(path string) (cfg *Configs, err error) {
	var iniFile *ini.File
	iniFile, err = ini.LooseLoad(path)
	if err != nil {
		return nil, err
	}

	cfg, err = DefaultConfigs()
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
	iniFile.SaveTo(path)

	return
}
