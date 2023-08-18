package log

import "dario.cat/mergo"

type Option struct {
	Out    ITreeWriter
	Prefix string
	Flag   int
}

func DefaultOption() Option {
	return Option{
		Out:    NewConsoleLogWriter(),
		Prefix: "[LOG] ",
		Flag:   Ldate | Ltime | Lmicroseconds | Lshortfile,
	}
}

func NewOption(dst *Option) Option {
	src := DefaultOption()

	err := mergo.Merge(dst, src)
	if err != nil {
		return src
	}

	return *dst
}
