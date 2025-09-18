package bus

import (
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
)

type IBus interface {
	On(eventKey string, handle func(state any)) (eventId string)
	Off(eventId string) (ok bool)
	Emit(eventKey string, state any)
}

func New() IBus {
	return new(bus).init()
}

type bus struct {
	counter  atomic.Int64
	eventMap sync.Map // id string -> func(state any)
}

func (b *bus) On(eventKey string, handle func(state any)) (eventId string) {
	newCounter := b.counter.Add(1)
	eventId = fmt.Sprintf("%s_%d", eventKey, newCounter)
	b.eventMap.Store(eventId, handle)
	return
}

func (b *bus) Off(eventId string) (ok bool) {
	_, ok = b.eventMap.LoadAndDelete(eventId)
	return
}

func (b *bus) Emit(eventKey string, state any) {
	b.eventMap.Range(func(key, value any) bool {
		if strings.HasPrefix(key.(string), eventKey+"_") {
			value.(func(state any))(state)
		}
		return true
	})
}

func (b *bus) init() *bus {
	return b
}
