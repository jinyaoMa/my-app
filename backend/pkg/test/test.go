package test

import (
	"testing"
	"time"
)

func LogTestingTime(t *testing.T, now time.Time) {
	t.Logf("timing: %v μs", time.Since(now).Microseconds())
}
