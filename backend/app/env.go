package app

import (
	"os"
)

type Env struct {
	air string
	log string
}

func DefaultEnv() *Env {
	return &Env{
		air: "0",
		log: "1",
	}
}

func LoadEnv() *Env {
	f := DefaultEnv()

	air := os.Getenv("MY_APP_AIR") // set MY_APP_AIR=1 to enable Web.Air function
	log := os.Getenv("MY_APP_LOG") // set MY_APP_LOG=0 to only log to console
	if air != "" {
		f.air = air
	}
	if log != "" {
		f.log = log
	}

	return f
}

func (f *Env) UseAir() bool {
	return f.air == "1"
}

func (f *Env) Log2File() bool {
	return f.log == "1"
}
