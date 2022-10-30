package app

import "os"

const (
	EnvMyAppAir = "MY_APP_AIR" // set MY_APP_AIR=1 to indicate using air hot reload tool
	EnvMyAppLog = "MY_APP_LOG" // set MY_APP_LOG=1 to log to file
)

type Env struct {
	pairs map[string]string
}

// DefaultEnv get default environment variables
func DefaultEnv() *Env {
	return &Env{
		pairs: map[string]string{
			EnvMyAppAir: "0",
			EnvMyAppLog: "1",
		},
	}
}

// LoadEnv load environment variables into map
func LoadEnv() *Env {
	e := DefaultEnv()
	for _, key := range []string{
		EnvMyAppAir,
		EnvMyAppLog,
	} {
		if value := os.Getenv(key); value != "" {
			e.pairs[key] = value
		}
	}
	return e
}

// UseAir run callback if use air hot reload tool
func (e *Env) UseAir(callback func()) *Env {
	if e.pairs[EnvMyAppAir] == "1" {
		callback()
	}
	return e
}

// Log2Console run callback if log to console
func (e *Env) Log2Console(callback func()) *Env {
	if e.pairs[EnvMyAppLog] != "1" {
		callback()
	}
	return e
}

// run callback if log to file
func (e *Env) Log2File(callback func()) *Env {
	if e.pairs[EnvMyAppLog] == "1" {
		callback()
	}
	return e
}
