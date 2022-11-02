package menus

import (
	"reflect"
	"sync"
)

type SelectList struct {
	IRefresh
	head     *SingleItem
	once     sync.Once
	options  []*SingleItem
	cases    []reflect.SelectCase
	selected chan string
}

func NewSelectList(head *SingleItem, capacity ...int) *SelectList {
	cap := 2
	if len(capacity) > 0 {
		cap = capacity[0]
	}
	return &SelectList{
		head:     head.Disable(),
		options:  make([]*SingleItem, 0, cap),
		cases:    make([]reflect.SelectCase, 0, cap),
		selected: make(chan string),
	}
}

func (sl *SelectList) AddOption(option *SingleItem) *SelectList {
	sl.options = append(sl.options, option)
	sl.cases = append(sl.cases, reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(option.Clicked())})
	return sl
}

func (sl *SelectList) AddOptions(options ...*SingleItem) *SelectList {
	var cases []reflect.SelectCase
	for _, opt := range options {
		cases = append(cases, reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(opt.Clicked())})
	}
	sl.cases = append(sl.cases, cases...)
	sl.options = append(sl.options, options...)
	return sl
}

func (sl *SelectList) UpdateText() *SelectList {
	sl.head.UpdateText()
	for _, item := range sl.options {
		item.UpdateText()
	}
	return sl
}

func (sl *SelectList) Check(id string) *SelectList {
	for _, opt := range sl.options {
		if opt.id == id {
			opt.Check()
		} else {
			opt.Uncheck()
		}
	}
	return sl
}

func (sl *SelectList) Selected() chan string {
	sl.once.Do(func() {
		go func() {
			for {
				chosen, _, _ := reflect.Select(sl.cases)
				sl.selected <- sl.options[chosen].GetID()
			}
		}()
	})
	return sl.selected
}
