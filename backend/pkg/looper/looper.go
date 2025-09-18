package looper

import (
	"errors"
	"sync/atomic"
)

type ILooper[T any] interface {
	Reset()
	Next() T
}

func New[T any](items ...T) (ILooper[T], error) {
	return new(looper[T]).init(items...)
}

type looperItem[T any] struct {
	value T
	next  *looperItem[T]
}

type looper[T any] struct {
	head    *looperItem[T]
	tail    *looperItem[T]
	current atomic.Pointer[looperItem[T]]
}

func (l *looper[T]) Reset() {
	l.current.Store(nil)
}

func (l *looper[T]) Next() T {
	if l.current.CompareAndSwap(nil, l.head) {
		return l.head.value
	}
	return l.next()
}

func (l *looper[T]) next() T {
	current := l.current.Load()
	if l.current.CompareAndSwap(current, current.next) {
		return current.next.value
	}
	return l.next()
}

func (l *looper[T]) newLooperItem(items ...T) (next *looperItem[T], tail *looperItem[T]) {
	if len(items) == 1 {
		next = &looperItem[T]{
			value: items[0],
		}
		tail = next
		return
	}
	next, tail = l.newLooperItem(items[1:]...)
	return &looperItem[T]{
		value: items[0],
		next:  next,
	}, tail
}

func (l *looper[T]) init(items ...T) (*looper[T], error) {
	if len(items) == 0 {
		return nil, errors.New("items is empty")
	}

	l.head, l.tail = l.newLooperItem(items...)
	l.tail.next = l.head
	l.Reset()
	return l, nil
}
