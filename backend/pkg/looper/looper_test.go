package looper_test

import (
	"testing"
	"time"

	"majinyao.cn/my-app/backend/pkg/looper"
	"majinyao.cn/my-app/backend/pkg/test"
)

func TestLooper(t *testing.T) {
	now := time.Now()
	defer test.LogTestingTime(t, now)

	looper, err := looper.New(1, 2, 3, 4, 5)
	if err != nil {
		t.Fatal(err)
	}

	if looper.Next() != 1 {
		t.Fatal("looper.Next() != 1")
	}
	if looper.Next() != 2 {
		t.Fatal("looper.Next() != 2")
	}
	if looper.Next() != 3 {
		t.Fatal("looper.Next() != 3")
	}
	if looper.Next() != 4 {
		t.Fatal("looper.Next() != 4")
	}
	if looper.Next() != 5 {
		t.Fatal("looper.Next() != 5")
	}
	if looper.Next() != 1 {
		t.Fatal("looper.Next() != 1")
	}

	looper.Reset()
	if looper.Next() != 1 {
		t.Fatal("looper.Next()!= 1")
	}
}
