package base

import "time"

type IDynamicCode interface {
	Get() string
	Verify(code string) bool
	Reset(expired time.Duration, size uint, runes ...rune)
	Stop()
}

type DynamicCode struct {
	size   uint
	runes  []rune
	code   string
	ticker *time.Ticker
	done   chan bool
}

// Get implements IDynamicCode.
func (dynamicCode *DynamicCode) Get() string {
	return dynamicCode.code
}

// Verify implements IDynamicCode.
func (dynamicCode *DynamicCode) Verify(code string) bool {
	return dynamicCode.code == code
}

// Reset implements IDynamicCode.
func (dynamicCode *DynamicCode) Reset(expired time.Duration, size uint, runes ...rune) {
	dynamicCode.Stop()
	dynamicCode.size = size
	dynamicCode.runes = runes
	dynamicCode.ticker.Reset(expired)
	go func(this *DynamicCode) {
		for {
			select {
			case <-this.done:
				return
			case <-this.ticker.C:
				this.code = GenerateCode(this.size, this.runes...)
			}
		}
	}(dynamicCode)
}

// Stop implements IDynamicCode.
func (dynamicCode *DynamicCode) Stop() {
	close(dynamicCode.done)
	dynamicCode.ticker.Stop()
	dynamicCode.done = make(chan bool)
}

func NewDynamicCode(expired time.Duration, size uint, runes ...rune) (dynamicCode *DynamicCode, iDynamicCode IDynamicCode) {
	dynamicCode = &DynamicCode{
		size:   size,
		runes:  runes,
		code:   GenerateCode(size, runes...),
		ticker: time.NewTicker(expired),
		done:   make(chan bool),
	}
	return dynamicCode, dynamicCode
}
