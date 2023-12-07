package base

import (
	"reflect"
	"sync"
)

type ICommand interface {
	Open()
	Close()
}

type CommandCase struct {
	Chan     reflect.Value // channel to receive signal
	Callback func(recv reflect.Value, ok bool) (quit bool)
}

type Command struct {
	mutex     sync.Mutex
	cases     []*CommandCase
	closeChan chan bool
}

// Open implements ICommand.
func (command *Command) Open() {
	command.mutex.Lock()
	defer command.mutex.Unlock()
	if command.closeChan == nil {
		command.closeChan = make(chan bool)
		go command.routine()
	}
}

// Close implements ICommand.
func (command *Command) Close() {
	command.mutex.Lock()
	defer command.mutex.Unlock()
	if command.closeChan != nil {
		command.closeChan <- true
		close(command.closeChan)
		command.closeChan = nil
	}
}

func (command *Command) routine() {
	count := len(command.cases)
	cases := make([]reflect.SelectCase, count+1)
	cases[0] = reflect.SelectCase{
		Dir:  reflect.SelectRecv,
		Chan: reflect.ValueOf(command.closeChan),
	}
	for i := 0; i < count; i++ {
		cases[i+1] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: command.cases[i].Chan,
		}
	}
	for {
		chosen, recv, ok := reflect.Select(cases)
		if chosen == 0 {
			return
		} else {
			if command.cases[chosen-1].Callback != nil && command.cases[chosen-1].Callback(recv, ok) {
				return
			}
		}
	}
}

func NewCommand(cases ...*CommandCase) (command *Command, iCommand ICommand) {
	command = &Command{
		cases: cases,
	}
	return command, command
}
