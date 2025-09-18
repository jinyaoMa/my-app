package reactive_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"majinyao.cn/my-app/backend/pkg/reactive"
	"majinyao.cn/my-app/backend/pkg/test"
)

func TestTransformable(t *testing.T) {
	now := time.Now()
	defer test.LogTestingTime(t, now)

	v, err := reactive.NewTransformable(1, func(value int64) (string, error) {
		return strconv.FormatInt(value, 10), nil
	}, func(fValue string) (err error) {
		if fValue != "1" {
			return fmt.Errorf("value %s was not expected", fValue)
		}
		return
	})
	if err != nil {
		t.Fatal(err)
	}
	expectValue := int64(2)
	expectFValue := "2"

	v1 := ""
	v2 := ""
	v3 := ""

	signal := make(chan struct{})

	go func() {
		v.Watch(func(value string) (err error) {
			v1 = value
			signal <- struct{}{}
			return
		})
		signal <- struct{}{}
	}()

	v2Id := v.Watch(func(value string) (err error) {
		v2 = value
		return
	})

	go func() {
		v.Watch(func(value string) (err error) {
			v3 = value
			signal <- struct{}{}
			return
		})
		signal <- struct{}{}
	}()
	<-signal
	<-signal

	v.Filter(func(value int64) (ok bool) {
		return value == expectValue
	})

	v.Unwatch(v2Id)

	v.Set(3)
	if v.Get() != 1 {
		t.Error(`value was not expected`)
		return
	}

	v.Set(2)
	<-signal
	<-signal
	if v1 != expectFValue || v2 != "" || v3 != expectFValue {
		t.Fatalf("values was not expected: v1=%s, v2=%s, v3=%s", v1, v2, v3)
	}

	transformedValue, err := v.GetTransformed()
	if err != nil {
		t.Fatal(err)
	}
	if transformedValue != expectFValue {
		t.Fatalf("transformed value was not expected: %s", transformedValue)
	}
}
