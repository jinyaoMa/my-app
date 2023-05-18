package utility

import (
	"math/rand"
	"my-app/backend/pkg/utility/interfaces"
	"strings"
	"time"
)

type Random struct {
	rand       *rand.Rand
	digitRunes []rune
}

// GenerateCode implements interfaces.IRandom
func (r *Random) GenerateCode(size int, chars ...rune) string {
	choices := append(r.digitRunes, chars...)
	max := len(choices)
	builder := new(strings.Builder)
	for i := 0; i < size; i++ {
		builder.WriteRune(choices[r.rand.Intn(max)])
	}
	return builder.String()
}

func NewRandom() interfaces.IRandom {
	return &Random{
		rand:       rand.New(rand.NewSource(time.Now().UnixNano())),
		digitRunes: []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'},
	}
}
