package reactive

import (
	"slices"
	"sync"
	"sync/atomic"
)

type IReactive[T any] interface {
	Get() T
	Set(value T, sync ...bool) (err error)
	Watch(func(value T) (err error)) (id int64)
	Unwatch(id int64) (listener func(value T) (err error), ok bool)
}

func New[T any](value T, instantWatches ...func(value T) (err error)) (IReactive[T], error) {
	return new(reactive[T]).init(value, instantWatches...)
}

type reactive[T any] struct {
	mutex       sync.RWMutex
	listenerMap sync.Map
	currentId   atomic.Int64
	value       T
}

func (r *reactive[T]) init(value T, instantWatches ...func(value T) (err error)) (*reactive[T], error) {
	r.value = value
	for _, watch := range instantWatches {
		if watch != nil {
			if err := watch(value); err != nil {
				return nil, err
			}
		}
	}
	return r, nil
}

func (r *reactive[T]) Get() T {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.value
}

func (r *reactive[T]) Set(value T, sync ...bool) (err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.value = value
	if slices.Contains(sync, true) {
		return r.updateListeners(value)
	}
	go r.updateListeners(value)
	return
}

func (r *reactive[T]) Watch(listener func(value T) (err error)) (id int64) {
	id = r.currentId.Add(1)
	r.listenerMap.Store(id, listener)
	return
}

func (r *reactive[T]) Unwatch(id int64) (listener func(value T) (err error), ok bool) {
	l, ok := r.listenerMap.LoadAndDelete(id)
	return l.(func(value T) (err error)), ok
}

func (r *reactive[T]) updateListeners(value T) (err error) {
	r.listenerMap.Range(func(k, v any) bool {
		listener := v.(func(value T) (err error))
		if listener != nil {
			if err = listener(value); err != nil {
				return false
			}
		}
		return true
	})
	return
}
