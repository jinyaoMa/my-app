package app

import (
	"flag"
)

type Flag struct {
	air uint
}

func DefaultFlag() *Flag {
	return &Flag{
		air: 0,
	}
}

func LoadFlag() *Flag {
	f := DefaultFlag()
	flag.UintVar(&f.air, "air", f.air, "set `-air 1` to enable web.Air function")
	flag.Parse()
	return f
}

func (f *Flag) UseAir() bool {
	return f.air == 1
}
