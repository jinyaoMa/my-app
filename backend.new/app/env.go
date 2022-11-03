package app

import "os"

const (
	EnvMyAppAir = "MY_APP_AIR" // set MY_APP_AIR=1 to indicate using air hot reload tool
	EnvMyAppLog = "MY_APP_LOG" // set MY_APP_LOG=1 to log to file, otherwise log to console
)

type env struct {
	pairs map[string]string
}

// DefaultEnv get default environment variables
func DefaultEnv() *env {
	return &env{
		pairs: map[string]string{
			EnvMyAppAir: "0",
			EnvMyAppLog: "1",
		},
	}
}

// LoadEnv load environment variables into map
func LoadEnv() *env {
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

// IsAir check if use air hot reload tool
func (e *env) IsAir() bool {
	return e.pairs[EnvMyAppAir] == "1"
}

// IsLog2File check if log to file
func (e *env) IsLog2File() bool {
	return e.pairs[EnvMyAppLog] == "1"
}
