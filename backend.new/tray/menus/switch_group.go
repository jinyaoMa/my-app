package menus

import (
	"reflect"
	"sync"
)

type SwitchGroup struct {
	flag bool

	onOnce    sync.Once
	onGroup   []*SingleItem
	onCases   []reflect.SelectCase
	onClicked chan string

	offOnce    sync.Once
	offGroup   []*SingleItem
	offCases   []reflect.SelectCase
	offClicked chan string
}

// capacity[0] -> On group capacity, capacity[1] -> Off group capacity
func NewSwitchGroup(default_flag bool, capacity ...int) *SwitchGroup {
	capOn := 2
	capOff := 2
	if len(capacity) > 0 {
		capOn = capacity[0]
	}
	if len(capacity) > 1 {
		capOff = capacity[1]
	}
	return (&SwitchGroup{
		flag:       default_flag,
		onGroup:    make([]*SingleItem, 0, capOn),
		onCases:    make([]reflect.SelectCase, 0, capOn),
		onClicked:  make(chan string),
		offGroup:   make([]*SingleItem, 0, capOff),
		offCases:   make([]reflect.SelectCase, 0, capOff),
		offClicked: make(chan string),
	})
}

func (sg *SwitchGroup) AddItems2OnGroup(items ...*SingleItem) *SwitchGroup {
	var cases []reflect.SelectCase
	for _, item := range items {
		cases = append(cases, reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(item.Clicked())})
	}
	sg.onCases = append(sg.onCases, cases...)
	sg.onGroup = append(sg.onGroup, items...)
	return sg.setFlag(sg.flag)
}

func (sg *SwitchGroup) AddItems2OffGroup(items ...*SingleItem) *SwitchGroup {
	var cases []reflect.SelectCase
	for _, item := range items {
		cases = append(cases, reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(item.Clicked())})
	}
	sg.offCases = append(sg.onCases, cases...)
	sg.offGroup = append(sg.offGroup, items...)
	return sg.setFlag(sg.flag)
}

func (sg *SwitchGroup) UpdateText() *SwitchGroup {
	for _, item := range sg.onGroup {
		item.UpdateText()
	}
	for _, item := range sg.offGroup {
		item.UpdateText()
	}
	return sg
}

func (sg *SwitchGroup) Toggle() *SwitchGroup {
	return sg.setFlag(!sg.flag)
}

func (sg *SwitchGroup) OnGroupClicked() chan string {
	sg.onOnce.Do(func() {
		go func() {
			for {
				chosen, _, _ := reflect.Select(sg.onCases)
				sg.onClicked <- sg.onGroup[chosen].GetID()
			}
		}()
	})
	return sg.onClicked
}

func (sg *SwitchGroup) OffGroupClicked() chan string {
	sg.offOnce.Do(func() {
		go func() {
			for {
				chosen, _, _ := reflect.Select(sg.offCases)
				sg.offClicked <- sg.offGroup[chosen].GetID()
			}
		}()
	})
	return sg.offClicked
}

func (sg *SwitchGroup) setFlag(flag bool) *SwitchGroup {
	sg.flag = flag
	if sg.flag {
		for _, item := range sg.onGroup {
			item.Show()
		}
		for _, item := range sg.offGroup {
			item.Hide()
		}
	} else {
		for _, item := range sg.onGroup {
			item.Hide()
		}
		for _, item := range sg.offGroup {
			item.Show()
		}
	}
	return sg
}
