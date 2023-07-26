package settings

import "gopkg.in/ini.v1"

type Settings struct {
}

func NewSettings(path string) (s *Settings, err error) {
	s = new(Settings)
	err = ini.MapTo(s, path)
	if err != nil {
		return nil, err
	}
	return
}
