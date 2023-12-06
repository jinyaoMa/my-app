package ctick

import (
	"my-app/pkg/base"
	"time"
)

type CodeTicker struct {
	options     *CodeTickerOptions
	code        string
	expiredTime time.Time
	ticker      *time.Ticker
	reset       chan *CodeTickerOptions
	done        chan bool
}

// Done implements ICodeTicker.
func (codeTicker *CodeTicker) Done() {
	close(codeTicker.done)
}

// Get implements ICodeTicker.
func (codeTicker *CodeTicker) Get() (code string, expiredTime time.Time) {
	return codeTicker.code, codeTicker.expiredTime
}

// Reset implements ICodeTicker.
func (codeTicker *CodeTicker) Reset(options *CodeTickerOptions) {
	codeTicker.reset <- options
}

// Stop implements ICodeTicker.
func (codeTicker *CodeTicker) Stop() {
	codeTicker.ticker.Stop()
}

// Verify implements ICodeTicker.
func (codeTicker *CodeTicker) Verify(code string) bool {
	return codeTicker.code == code
}

func (codeTicker *CodeTicker) refresh() {
	codeTicker.code = base.GenerateCode(codeTicker.options.Size, codeTicker.options.Chars...)
	codeTicker.expiredTime = time.Now().Add(codeTicker.options.Expiration)
}

func NewCodeTicker(options *CodeTickerOptions) (codeTicker *CodeTicker, iCodeTicker ICodeTicker, err error) {
	options, err = NewCodeTickerOptions(options)
	if err != nil {
		return nil, nil, err
	}

	codeTicker = &CodeTicker{
		options:     options,
		code:        base.GenerateCode(options.Size, options.Chars...),
		expiredTime: time.Now().Add(options.Expiration),
		ticker:      time.NewTicker(options.Expiration),
		reset:       make(chan *CodeTickerOptions),
		done:        make(chan bool),
	}
	go func(codeTicker *CodeTicker) {
		for {
			select {
			case <-codeTicker.done:
				return
			case codeTicker.options = <-codeTicker.reset:
				codeTicker.refresh()
				codeTicker.ticker.Reset(options.Expiration)
			case <-codeTicker.ticker.C:
				codeTicker.refresh()
			}
		}
	}(codeTicker)
	return codeTicker, codeTicker, nil
}
