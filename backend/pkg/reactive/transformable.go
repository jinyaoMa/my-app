package reactive

import (
	"slices"
	"sync"
	"sync/atomic"
)

type ITransformable[T any, F any] interface {
	Get() T
	GetTransformed() (fValuef F, err error)
	Set(value T, sync ...bool) (err error)
	Watch(listener func(fValue F) (err error)) (id int64)
	Unwatch(id int64) (listener func(fValue F) (err error), ok bool)
	Transform(transformer func(value T) (fValue F, err error))
	Filter(filter func(value T) (ok bool))
}

type transformable[T any, F any] struct {
	mutex       sync.RWMutex
	listenerMap sync.Map
	currentId   atomic.Int64
	value       T
	transformer func(value T) (fValue F, err error)
	filter      func(value T) (ok bool)
}

func NewTransformable[T any, F any](value T, transformer func(value T) (fValue F, err error), instantWatches ...func(fValue F) (err error)) (ITransformable[T, F], error) {
	return new(transformable[T, F]).init(value, transformer, instantWatches...)
}

func (t *transformable[T, F]) init(value T, transformer func(value T) (fValue F, err error), instantWatches ...func(fValue F) (err error)) (*transformable[T, F], error) {
	t.value = value
	t.transformer = transformer
	if len(instantWatches) > 0 {
		var fValue F
		var err error
		if transformer != nil {
			fValue, err = transformer(value)
			if err != nil {
				return nil, err
			}
		}
		for _, watch := range instantWatches {
			if watch != nil {
				err = watch(fValue)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	return t, nil
}

func (t *transformable[T, F]) Get() T {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return t.value
}

func (t *transformable[T, F]) GetTransformed() (fValue F, err error) {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	if t.transformer != nil {
		fValue, err = t.transformer(t.value)
		if err != nil {
			return
		}
	}
	return
}

func (t *transformable[T, F]) Set(value T, sync ...bool) (err error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	if t.filter != nil && !t.filter(value) {
		return
	}

	t.value = value

	var fValue F
	if t.transformer != nil {
		fValue, err = t.transformer(value)
		if err != nil {
			return
		}
	}
	if slices.Contains(sync, true) {
		return t.updateListeners(fValue)
	}
	go t.updateListeners(fValue)
	return
}

func (t *transformable[T, F]) Watch(listener func(fValue F) (err error)) (id int64) {
	id = t.currentId.Add(1)
	t.listenerMap.Store(id, listener)
	return
}

func (t *transformable[T, F]) Unwatch(id int64) (listener func(fValue F) (err error), ok bool) {
	l, ok := t.listenerMap.LoadAndDelete(id)
	return l.(func(fValue F) (err error)), ok
}

func (t *transformable[T, F]) Transform(transformer func(value T) (fValue F, err error)) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.transformer = transformer
}

func (t *transformable[T, F]) Filter(filter func(value T) (ok bool)) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.filter = filter
}

func (t *transformable[T, F]) updateListeners(fValue F) (err error) {
	t.listenerMap.Range(func(k, v any) bool {
		listener := v.(func(fValue F) (err error))
		if listener != nil {
			err = listener(fValue)
			if err != nil {
				return false
			}
		}
		return true
	})
	return
}
