package code

import (
	"math/rand"
	"strings"
	"time"
)

type Code struct {
	rand       *rand.Rand
	digitRunes []rune
}

// Generate implements ICode
func (c *Code) Generate(size int, chars ...rune) string {
	choices := append(c.digitRunes, chars...)
	max := len(choices)
	builder := new(strings.Builder)
	for i := 0; i < size; i++ {
		builder.WriteRune(choices[c.rand.Intn(max)])
	}
	return builder.String()
}

func New() *Code {
	return &Code{
		rand:       rand.New(rand.NewSource(time.Now().UnixNano())),
		digitRunes: []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'},
	}
}

func NewICode() ICode {
	return New()
}
