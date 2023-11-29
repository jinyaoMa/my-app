package base

import "strings"

var (
	CodeDigitRunes = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
)

func GenerateCode(size int, runes ...rune) string {
	builder := new(strings.Builder)
	num := len(runes)
	if num > 0 {
		for i := 0; i < size; i++ {
			builder.WriteRune(runes[randGenerator.Intn(num)])
		}
	} else {
		num := len(CodeDigitRunes)
		for i := 0; i < size; i++ {
			builder.WriteRune(CodeDigitRunes[randGenerator.Intn(num)])
		}
	}
	return builder.String()
}
