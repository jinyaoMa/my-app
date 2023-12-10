package base

import (
	"math/rand"
	"time"
)

var (
	randGenerator *rand.Rand
)

func init() {
	randGenerator = rand.New(rand.NewSource(time.Now().UnixMicro()))
}

func GetRandGenerator() *rand.Rand {
	return randGenerator
}
