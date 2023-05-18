package interfaces

type IRandom interface {
	// GenerateCode generate verification code or any random string
	GenerateCode(size int, chars ...rune) string
}
