package ctick

import (
	"my-app/pkg/base"
	"time"
)

func (codeTicker *CodeTicker) refresh() {
	codeTicker.token.Code = base.GenerateCode(codeTicker.options.Size, codeTicker.options.Chars...)
	codeTicker.token.ExpiredTime = time.Now().Add(codeTicker.options.Expiration)
}

func (codeTicker *CodeTicker) reset(options *CodeTickerOptions) {
	codeTicker.refresh()
	codeTicker.ticker.Reset(options.Expiration)
}

func (codeTicker *CodeTicker) routine() {
	for {
		select {
		case <-codeTicker.closeCh:
			return
		case options := <-codeTicker.resetCh:
			codeTicker.reset(options)
		case tokenCh := <-codeTicker.getCh:
			tokenCh <- codeTicker.token
		case <-codeTicker.ticker.C:
			codeTicker.refresh()
		}
	}
}
