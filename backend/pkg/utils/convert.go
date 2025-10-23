package utils

import (
	"encoding/base64"
	"encoding/binary"
	"strconv"
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

func ConvertInt64ToBase64(v int64) string {
	return base64.RawURLEncoding.EncodeToString(ConvertInt64ToBytes(v))
}

func ConvertBase64ToInt64(s string) (int64, error) {
	buf, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return 0, err
	}
	return ConvertBytesToInt64(buf), nil
}

func ConvertInt64ToB36(v int64) string {
	return strconv.FormatUint(uint64(v), 36)
}

func ConvertB36ToInt64(s string) (int64, error) {
	v, err := strconv.ParseUint(s, 36, 64)
	return int64(v), err
}
