package config

type Options[T any] struct {
	Path    string `json:"path"`
	Default T      `json:"default"` // default config
}
