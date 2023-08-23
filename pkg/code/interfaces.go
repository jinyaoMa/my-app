package code

type ICode interface {
	// Generate generate verification code or any random string
	Generate(size int, chars ...rune) string
}
