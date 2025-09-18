package utils

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
)

func ConvertInt64ToBytes(v int64) []byte {
	buf := make([]byte, 8)
	binary.PutVarint(buf, v)
	return buf
}

func ConvertBytesToInt64(buf []byte) int64 {
	v, _ := binary.Varint(buf)
	return v
}

func ConvertUint64ToBytes(v uint64) []byte {
	buf := make([]byte, 8)
	binary.PutUvarint(buf, v)
	return buf
}

func ConvertBytesToUint64(buf []byte) uint64 {
	v, _ := binary.Uvarint(buf)
	return v
}

func ConvertUint64ToBase64(v uint64) string {
	return base64.RawURLEncoding.EncodeToString(ConvertUint64ToBytes(v))
}

func ConvertBase64ToUint64(s string) (uint64, error) {
	buf, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return 0, err
	}
	return ConvertBytesToUint64(buf), nil
}

func ConvertInt64ToHex(v int64) string {
	return hex.EncodeToString(ConvertInt64ToBytes(v))
}

func ConvertHexToInt64(s string) (int64, error) {
	buf, err := hex.DecodeString(s)
	if err != nil {
		return 0, err
	}
	return ConvertBytesToInt64(buf), nil
}
