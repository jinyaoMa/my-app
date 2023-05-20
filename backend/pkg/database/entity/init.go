package entity

import (
	"my-app/backend/pkg/crypto"
	"my-app/backend/pkg/snowflake"
)

var (
	idGenerator snowflake.Interface
	cipher      crypto.Interface
)

func IdGenerator(abc ...snowflake.Interface) snowflake.Interface {
	if len(abc) == 1 {
		idGenerator = abc[0]
	}
	return idGenerator
}

func Cipher(abc ...crypto.Interface) crypto.Interface {
	if len(abc) == 1 {
		cipher = abc[0]
	}
	return cipher
}
