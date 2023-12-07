package ctick

import (
	"my-app/pkg/base"
	"time"
)

type CodeTicker struct {
	options   *CodeTickerOptions
	token     *CodeToken
	ticker    *time.Ticker
	resetChan chan *CodeTickerOptions
	getChan   chan chan *CodeToken
	doneChan  chan bool
}

type CodeToken struct {
	Code        string
	ExpiredTime time.Time
}

// Done implements ICodeTicker.
func (codeTicker *CodeTicker) Done() {
	codeTicker.Stop()
	close(codeTicker.doneChan)
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

// Stop implements ICodeTicker.
func (codeTicker *CodeTicker) Stop() {
	codeTicker.ticker.Stop()
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
		resetChan: make(chan *CodeTickerOptions),
		getChan:   make(chan chan *CodeToken),
		doneChan:  make(chan bool),
	}
	go func(codeTicker *CodeTicker) {
		for {
			select {
			case <-codeTicker.doneChan:
				return
			case options := <-codeTicker.resetChan:
				codeTicker.reset(options)
			case tokenChan := <-codeTicker.getChan:
				tokenChan <- codeTicker.token
			case <-codeTicker.ticker.C:
				codeTicker.refresh()
			}
		}
	}(codeTicker)
	return codeTicker, codeTicker, nil
}
