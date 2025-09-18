package executable_test

import (
	"testing"
	"time"

	"majinyao.cn/my-app/backend/pkg/executable"
	"majinyao.cn/my-app/backend/pkg/test"
)

func TestExecutable(t *testing.T) {
	now := time.Now()
	defer test.LogTestingTime(t, now)

	exe, err := executable.New()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(exe.GetPath())
	t.Log(exe.GetDir())
	t.Log(exe.GetBase())
	t.Log(exe.GetName())
	t.Log(exe.GetExt())
}
