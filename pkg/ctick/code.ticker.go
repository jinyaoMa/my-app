package ctick

import (
	"my-app/pkg/base"
	"runtime"
	"time"
)

type CodeTicker struct {
	options *CodeTickerOptions
	token   *CodeToken
	ticker  *time.Ticker
	resetCh chan *CodeTickerOptions
	getCh   chan chan *CodeToken
	closeCh chan struct{}
}

// Close implements ICodeTicker.
func (codeTicker *CodeTicker) Close() {
	codeTicker.ticker.Stop()
	close(codeTicker.closeCh)
	codeTicker.closeCh = make(chan struct{})
}

// Get implements ICodeTicker.
func (codeTicker *CodeTicker) Get() (ticket *CodeToken) {
	tokenCh := make(chan *CodeToken, 1)
	defer close(tokenCh)

	codeTicker.getCh <- tokenCh
	ticket = <-tokenCh
	return
}

// Reset implements ICodeTicker.
func (codeTicker *CodeTicker) Reset(options *CodeTickerOptions) {
	codeTicker.resetCh <- options
}

// Verify implements ICodeTicker.
func (codeTicker *CodeTicker) Verify(code string) bool {
	return codeTicker.token.Code == code
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
		ticker:  time.NewTicker(options.Expiration),
		resetCh: make(chan *CodeTickerOptions, 1),
		getCh:   make(chan chan *CodeToken, runtime.NumCPU()),
		closeCh: make(chan struct{}),
	}
	go codeTicker.routine()
	return codeTicker, codeTicker, nil
}
