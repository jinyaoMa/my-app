package base

import (
	"math/rand"
	"time"
)

var (
	randGenerator *rand.Rand
)

func init() {
	randGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func GetRandGenerator() *rand.Rand {
	return randGenerator
}
