package config_test

import (
	"os"
	"testing"
	"time"

	"majinyao.cn/my-app/backend/pkg/config"
	"majinyao.cn/my-app/backend/pkg/test"
)

type MyConfig struct {
	Test string `json:"test"`
}

func TestConfig(t *testing.T) {
	now := time.Now()
	defer test.LogTestingTime(t, now)

	defer os.Remove("test.json")
	cfg, err := config.New(config.Options[MyConfig]{
		Path: "test.json",
		Default: MyConfig{
			Test: "test",
		},
	})
	if err != nil {
		t.Fatalf("error when creating config: %v", err)
	}

	cfg.Save(MyConfig{Test: "test2"})
	cfg.Reload()
	cfg.Save(MyConfig{Test: "test3"})
	if cfg.Get().Test != "test3" {
		t.Fatalf("expected test3, got %s", cfg.Get().Test)
	}
}
