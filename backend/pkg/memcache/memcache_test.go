package memcache_test

import (
	"testing"
	"time"

	"majinyao.cn/my-app/backend/pkg/memcache"
	"majinyao.cn/my-app/backend/pkg/test"
)

func TestMemcache(t *testing.T) {
	now := time.Now()
	defer test.LogTestingTime(t, now)

	m := memcache.New(1000)
	m.Set("key", "value", now.Add(time.Second))
	value, err := m.Get("key")
	if err != nil {
		t.Fatal(err)
	}
	if value != "value" {
		t.Fatal("value not equal")
	}

	time.Sleep(time.Second)
	value, err = m.Get("key")
	if err == nil {
		t.Fatal("key not expired")
	}
	if value != nil {
		t.Fatal("value not nil")
	}

	m.Set("key", "value", now.Add(time.Second*10))
	m.Remove("key")
	value, err = m.Get("key")
	if err == nil {
		t.Fatal("key not removed")
	}
	if value != nil {
		t.Fatal("value not nil")
	}
}
