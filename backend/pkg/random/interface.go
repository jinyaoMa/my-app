package random

type Interface interface {
	// GenerateCode generate verification code or any random string
	GenerateCode(size int, chars ...rune) string
}
