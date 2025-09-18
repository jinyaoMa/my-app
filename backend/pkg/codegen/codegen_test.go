package codegen_test

import (
	"testing"
	"time"

	"majinyao.cn/my-app/backend/pkg/codegen"
	"majinyao.cn/my-app/backend/pkg/test"
)

func TestCodegen(t *testing.T) {
	now := time.Now()
	defer test.LogTestingTime(t, now)

	chars := codegen.Digits + codegen.Letters
	cg, err := codegen.New(codegen.Options{
		Characters: chars,
	})
	if err != nil {
		t.Fatal(err)
	}

	code := cg.Generate(64)
	if len(code) != 64 {
		t.Fatal("code length must be 64")
	}
}
