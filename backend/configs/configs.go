package configs

import "gopkg.in/ini.v1"

type Configs struct {
}

func NewConfigs(path string) (cfg *Configs, err error) {
	cfg = new(Configs)
	err = ini.MapTo(cfg, path)
	if err != nil {
		return nil, err
	}
	return
}
