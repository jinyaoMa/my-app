package ctick

type ICodeTicker interface {
	Get() (ticket *CodeToken)
	Verify(code string) bool
	Reset(options *CodeTickerOptions)
	Stop()
	Done()
}
