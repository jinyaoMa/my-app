package ctick

import (
	"my-app/pkg/base"
	"reflect"
	"runtime"
	"time"
)

type CodeTicker struct {
	command   base.ICommand
	options   *CodeTickerOptions
	token     *CodeToken
	ticker    *time.Ticker
	resetChan chan *CodeTickerOptions
	getChan   chan chan *CodeToken
}

// Close implements ICodeTicker.
func (codeTicker *CodeTicker) Close() {
	codeTicker.ticker.Stop()
	codeTicker.command.Close()
}

// Get implements ICodeTicker.
func (codeTicker *CodeTicker) Get() (ticket *CodeToken) {
	tokenChan := make(chan *CodeToken, 1)
	defer close(tokenChan)

	codeTicker.getChan <- tokenChan
	ticket = <-tokenChan
	return
}

// Reset implements ICodeTicker.
func (codeTicker *CodeTicker) Reset(options *CodeTickerOptions) {
	codeTicker.resetChan <- options
}

// Verify implements ICodeTicker.
func (codeTicker *CodeTicker) Verify(code string) bool {
	return codeTicker.token.Code == code
}

func (codeTicker *CodeTicker) refresh() {
	codeTicker.token.Code = base.GenerateCode(codeTicker.options.Size, codeTicker.options.Chars...)
	codeTicker.token.ExpiredTime = time.Now().Add(codeTicker.options.Expiration)
}

func (codeTicker *CodeTicker) reset(options *CodeTickerOptions) {
	codeTicker.refresh()
	codeTicker.ticker.Reset(options.Expiration)
}

func NewCodeTicker(options *CodeTickerOptions) (codeTicker *CodeTicker, iCodeTicker ICodeTicker, err error) {
	options, err = NewCodeTickerOptions(options)
	if err != nil {
		return nil, nil, err
	}

	codeTicker = &CodeTicker{
		options: options,
		token: &CodeToken{
			Code:        base.GenerateCode(options.Size, options.Chars...),
			ExpiredTime: time.Now().Add(options.Expiration),
		},
		ticker:    time.NewTicker(options.Expiration),
		resetChan: make(chan *CodeTickerOptions, 1),
		getChan:   make(chan chan *CodeToken, runtime.NumCPU()),
	}

	_, codeTicker.command = base.NewCommand(&base.CommandCase{
		Chan: reflect.ValueOf(codeTicker.resetChan),
		Callback: func(recv reflect.Value, ok bool) bool {
			if options, ok := recv.Interface().(*CodeTickerOptions); ok {
				codeTicker.reset(options)
			}
			return false
		},
	}, &base.CommandCase{
		Chan: reflect.ValueOf(codeTicker.getChan),
		Callback: func(recv reflect.Value, ok bool) bool {
			if tokenChan, ok := recv.Interface().(chan *CodeToken); ok {
				tokenChan <- codeTicker.token
			}
			return false
		},
	}, &base.CommandCase{
		Chan: reflect.ValueOf(codeTicker.ticker.C),
		Callback: func(recv reflect.Value, ok bool) bool {
			codeTicker.refresh()
			return false
		},
	})
	codeTicker.command.Open()
	return codeTicker, codeTicker, nil
}
