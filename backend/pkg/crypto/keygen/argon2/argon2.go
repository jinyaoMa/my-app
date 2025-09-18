package argon2

import (
	"bytes"
	"encoding/base64"
	"runtime"
	"slices"

	a2 "golang.org/x/crypto/argon2"
)

const Alg string = "argon2"

func New(options Options) *argon2 {
	return new(argon2).init(options)
}

type argon2 struct {
	salt      []byte
	time      uint32
	memory    uint32
	threads   uint8
	keyLength uint32
	prefix    []byte
}

func (a *argon2) Derive(password []byte) (key []byte) {
	return a2.IDKey(
		password,
		a.salt,
		a.time,
		a.memory,
		a.threads,
		a.keyLength)
}

func (a *argon2) DeriveBase64(password string, hasPrefix ...bool) (key string) {
	bytes := a.Derive([]byte(password))
	if slices.Contains(hasPrefix, true) {
		return base64.RawURLEncoding.EncodeToString(append(a.prefix, bytes...))
	}
	return base64.RawURLEncoding.EncodeToString(bytes)
}

func (a *argon2) Verify(key []byte, password []byte) bool {
	newKey := a.Derive(password)
	return bytes.Equal(newKey, key)
}

func (a *argon2) VerifyBase64(key string, password string, hasPrefix ...bool) bool {
	newKey := a.DeriveBase64(password, hasPrefix...)
	return newKey == key
}

func (a *argon2) init(options Options) *argon2 {
	numcpu := runtime.NumCPU()
	if numcpu > 254 {
		numcpu = 254
	} else if numcpu < 1 {
		numcpu = 1
	}

	threads := uint8(options.Threads)
	if options.Threads > numcpu {
		threads = uint8(numcpu)
	} else if options.Threads < 1 {
		threads = 1
	}

	keyLength := uint32(options.KeyLength)
	if options.KeyLength > 512 {
		keyLength = 512
	} else if options.KeyLength < 2 {
		keyLength = 2
	}

	prefix := []byte(options.Prefix)
	if len(prefix) == 0 {
		prefix = []byte(Alg + "_")
	}

	a.salt = []byte(options.Salt)
	a.time = 1           // 1 iteration
	a.memory = 64 * 1024 // 64 MB
	a.threads = threads
	a.keyLength = keyLength
	a.prefix = prefix
	return a
}
