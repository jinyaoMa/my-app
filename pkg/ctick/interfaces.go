package ctick

import "time"

type ICodeTicker interface {
	Get() (code string, expiredTime time.Time)
	Verify(code string) bool
	Reset(options *CodeTickerOptions)
	Stop()
	Done()
}
