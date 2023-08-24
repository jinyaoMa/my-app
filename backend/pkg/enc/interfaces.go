package enc

type ICipher interface {
	// Encrypt encrypt plaintext to hexadecimal encoded ciphertext
	Encrypt(plaintext string) (ciphertext string, err error)

	// Decrypt decrypt hexadecimal encoded ciphertext to plaintext
	Decrypt(ciphertext string) (plaintext string, err error)
}
