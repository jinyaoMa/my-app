package memcache

import (
	"fmt"
	"slices"
	"sync"
	"sync/atomic"
	"time"
)

type IMemcache interface {
	Get(key string) (value any, err error)
	Set(key string, value any, expiredAt ...time.Time) (oldValue any, exist bool, err error)
	Remove(key string) (value any, err error)
	Clear(expiredOnly ...bool)
}

func New(limit int) IMemcache {
	return new(memcache).init(limit)
}

type memcacheItem struct {
	value     any
	expiredAt *time.Time
}

type memcache struct {
	store    sync.Map // key string -> *memcacheItem
	setCount atomic.Int64
	limit    int
}

// Get implements IMemcache.
func (m *memcache) Get(key string) (value any, err error) {
	v, ok := m.store.Load(key)
	if !ok {
		return nil, fmt.Errorf("key `%s` not found", key)
	}

	item := v.(*memcacheItem)
	if item.expiredAt != nil && item.expiredAt.Before(time.Now()) {
		m.store.CompareAndDelete(key, v)
		return value, fmt.Errorf("key `%s` expired", key)
	}
	return item.value, nil
}

// Remove implements IMemcache.
func (m *memcache) Remove(key string) (value any, err error) {
	v, ok := m.store.LoadAndDelete(key)
	if !ok {
		return value, fmt.Errorf("key `%s` not found", key)
	}

	item := v.(*memcacheItem)
	return item.value, nil
}

// Set implements IMemcache.
func (m *memcache) Set(key string, value any, expiredAt ...time.Time) (oldValue any, exist bool, err error) {
	if m.limit > 0 {
		if curSetCount := m.setCount.Load(); curSetCount > int64(m.limit) {
			if m.setCount.CompareAndSwap(curSetCount, 0) {
				m.Clear()
			}
		}
	}

	var e *time.Time
	if len(expiredAt) > 0 {
		at := slices.MinFunc(expiredAt, func(a, b time.Time) int {
			if a.Before(b) {
				return -1
			}
			return 1 // b before a, so return 1 to indicate b is less than a, which means b should come before a in the sorted slice
		})
		if at.Before(time.Now()) {
			return nil, false, fmt.Errorf("expiredAt `%s` is before now", at)
		}
		e = &at
	}

	v, exist := m.store.Swap(key, &memcacheItem{
		value:     value,
		expiredAt: e,
	})

	if exist {
		oldValue = v.(*memcacheItem).value
		go m.Clear(true)
		return
	}
	m.setCount.Add(1)
	return
}

// Clear implements IMemcache.
func (m *memcache) Clear(expiredOnly ...bool) {
	if slices.Contains(expiredOnly, true) {
		m.store.Range(func(key, value any) bool {
			item := value.(*memcacheItem)
			if item.expiredAt != nil && item.expiredAt.Before(time.Now()) {
				m.store.CompareAndDelete(key, value)
			}
			return true
		})
	} else {
		m.store.Clear()
	}
}

func (m *memcache) init(limit int) *memcache {
	m.limit = limit
	return m
}
