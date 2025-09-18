package bus_test

import (
	"testing"
	"time"

	"majinyao.cn/my-app/backend/pkg/bus"
	"majinyao.cn/my-app/backend/pkg/test"
)

func TestBus(t *testing.T) {
	now := time.Now()
	defer test.LogTestingTime(t, now)

	test1 := "test1"
	test2 := "test2"
	test3 := "test3"

	bus := bus.New()
	wait := make(chan struct{})

	test1Id := bus.On(test1, func(state any) {
		if num, ok := state.(int); ok && num == 1 {
			wait <- struct{}{}
			return
		}
	})

	var test2Id, test3Id string
	go func() {
		test2Id = bus.On(test2, func(state any) {
			if num, ok := state.(string); ok && num == "2" {
				wait <- struct{}{}
				return
			}
		})

		bus.Emit(test1, 1)
		bus.Emit(test2, "2")

		test3Id = bus.On(test3, func(state any) {
			if num, ok := state.(float64); ok && num == 3.0 {
				wait <- struct{}{}
				return
			}
		})

		go func() {
			bus.Emit(test3, 3.0)
		}()
	}()

	for i := 0; i < 3; i++ {
		if _, ok := <-wait; !ok {
			t.Fatal("all listeners should be handled correctly")
		}
	}

	if ok := bus.Off(test1Id); !ok {
		t.Fatal("error turn off test1")
	}

	if ok := bus.Off(test2Id); !ok {
		t.Fatal("error turn off test2")
	}

	if ok := bus.Off(test3Id); !ok {
		t.Fatal("error turn off test3")
	}
}
