package base

import "strings"

var (
	CodeDigits    = strings.Split("0123456789", "")
	CodeHexDigits = strings.Split("0123456789ABCDEF", "")
	CodeLetters   = strings.Split("abcdefghijklmnopqrstuvwxyz", "")
	CodeULetters  = strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXZ", "")
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
