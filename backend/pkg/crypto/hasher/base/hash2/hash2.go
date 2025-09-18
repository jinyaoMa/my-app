package hash2

import (
	"bytes"
	"encoding/base64"
	"hash"
	"slices"
)

func New(h hash.Hash, alg string, options Options) IHash2 {
	return new(hash2).init(h, alg, options)
}

type IHash2 interface {
	Alg() string
	Serial() string
	CheckSerial(serial string) bool
	Write(p []byte) (n int, err error)
	Sum() (sum []byte)
	SumBase64(hasPrefix ...bool) (b64 string)
	Check(sum []byte) (ok bool)
	CheckBase64(b64 string, hasPrefix ...bool) (ok bool)
	Reset()
}

type hash2 struct {
	hash   hash.Hash
	alg    string
	serial string
	salt   []byte
	prefix []byte
}

func (h *hash2) Alg() string {
	return h.alg
}

func (h *hash2) Serial() string {
	return h.serial
}

func (h *hash2) CheckSerial(serial string) bool {
	return h.serial == serial
}

func (h *hash2) Write(p []byte) (n int, err error) {
	return h.hash.Write(p)
}

func (h *hash2) Sum() []byte {
	return h.hash.Sum(h.salt)
}

func (h *hash2) SumBase64(hasPrefix ...bool) (b64 string) {
	if slices.Contains(hasPrefix, true) {
		return base64.RawURLEncoding.EncodeToString(append(h.prefix, h.Sum()...))
	}
	return base64.RawURLEncoding.EncodeToString(h.Sum())
}

func (h *hash2) Check(sum []byte) (ok bool) {
	newSum := h.Sum()
	return bytes.Equal(newSum, sum)
}

func (h *hash2) CheckBase64(b64 string, hasPrefix ...bool) (ok bool) {
	newB64 := h.SumBase64(hasPrefix...)
	return newB64 == b64
}

func (h *hash2) Reset() {
	h.hash.Reset()
}

func (h *hash2) init(h2 hash.Hash, alg string, options Options) *hash2 {
	prefix := []byte(options.Prefix)
	if len(prefix) == 0 {
		prefix = []byte(alg + "_")
	}

	h.hash = h2
	h.alg = alg
	h.salt = []byte(options.Salt)
	h.prefix = prefix
	h.serial = base64.RawURLEncoding.EncodeToString([]byte(alg)) + "," +
		base64.RawURLEncoding.EncodeToString(h.salt) + "," +
		base64.RawURLEncoding.EncodeToString(h.prefix)
	return h
}
