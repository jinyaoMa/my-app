package base

import "strings"

var (
	CodeDigits = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
)

func GenerateCode(size uint, chars ...string) string {
	builder := new(strings.Builder)
	num := len(chars)
	if num > 0 {
		for i := uint(0); i < size; i++ {
			builder.WriteString(chars[randGenerator.Intn(num)])
		}
	} else {
		num := len(CodeDigits)
		for i := uint(0); i < size; i++ {
			builder.WriteString(CodeDigits[randGenerator.Intn(num)])
		}
	}
	return builder.String()
}
