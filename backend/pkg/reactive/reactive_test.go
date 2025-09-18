package reactive_test

import (
	"testing"
	"time"

	"majinyao.cn/my-app/backend/pkg/reactive"
	"majinyao.cn/my-app/backend/pkg/test"
)

func TestReactive(t *testing.T) {
	now := time.Now()
	defer test.LogTestingTime(t, now)

	v, _ := reactive.New(1)
	expectValue := 2

	v1 := 0
	v2 := 0
	v3 := 0

	signal := make(chan struct{})

	go func() {
		v.Watch(func(value int) (err error) {
			v1 = value
			signal <- struct{}{}
			return
		})
		signal <- struct{}{}
	}()

	v2Id := v.Watch(func(value int) (err error) {
		v2 = value
		return
	})

	go func() {
		v.Watch(func(value int) (err error) {
			v3 = value
			signal <- struct{}{}
			return
		})
		signal <- struct{}{}
	}()
	<-signal
	<-signal

	v.Unwatch(v2Id)

	v.Set(expectValue)
	<-signal
	<-signal
	if v1 != expectValue || v2 != 0 || v3 != expectValue {
		t.Error(`values was not expected`)
		return
	}

	realValue := v.Get()
	if realValue != expectValue {
		t.Error(`real value was not expected`)
		return
	}
}
