package log

import "dario.cat/mergo"

type Config struct {
	Out    ITreeWriter
	Prefix string
	Flag   int
}

func DefaultConfig() *Config {
	return &Config{
		Out:    NewConsoleLogWriter(),
		Prefix: "[LOG] ",
		Flag:   Ldate | Ltime | Lmicroseconds | Lshortfile,
	}
}

func NewConfig(dst *Config) *Config {
	src := DefaultConfig()

	err := mergo.Merge(dst, *src)
	if err != nil {
		return src
	}

	return dst
}
