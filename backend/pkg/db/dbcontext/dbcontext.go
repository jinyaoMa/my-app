package dbcontext

import (
	"context"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/crypto/cipher"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher"
	"majinyao.cn/my-app/backend/pkg/crypto/keygen"
	"majinyao.cn/my-app/backend/pkg/snowflake"
)

type key int

const (
	keySnowflake key = iota + 1
	keyKeygen
	keyHasher
	keyCipher
)

func SetSnowflake(tx *gorm.DB, options snowflake.Options) (s snowflake.ISnowflake, err error) {
	tx.Statement.Context, s, err = SetSnowflakeToContext(tx.Statement.Context, options)
	return
}

func SetSnowflakeToContext(ctx context.Context, options snowflake.Options) (newCtx context.Context, s snowflake.ISnowflake, err error) {
	s, err = snowflake.New(options)
	if err != nil {
		return
	}
	newCtx = context.WithValue(ctx, keySnowflake, s)
	return
}

func GetSnowflake(tx *gorm.DB) (s snowflake.ISnowflake, ok bool) {
	return GetSnowflakeFromContext(tx.Statement.Context)
}

func GetSnowflakeFromContext(ctx context.Context) (s snowflake.ISnowflake, ok bool) {
	v := ctx.Value(keySnowflake)
	if v == nil {
		return nil, false
	}

	s, ok = v.(snowflake.ISnowflake)
	return
}

func SetKeygen(tx *gorm.DB, options keygen.Options) (k keygen.IKeygen, err error) {
	tx.Statement.Context, k, err = SetKeygenToContext(tx.Statement.Context, options)
	return
}

func SetKeygenToContext(ctx context.Context, options keygen.Options) (newCtx context.Context, k keygen.IKeygen, err error) {
	k, err = keygen.New(options)
	if err != nil {
		return
	}
	newCtx = context.WithValue(ctx, keyKeygen, k)
	return
}

func GetKeygen(tx *gorm.DB) (k keygen.IKeygen, ok bool) {
	return GetKeygenFromContext(tx.Statement.Context)
}

func GetKeygenFromContext(ctx context.Context) (k keygen.IKeygen, ok bool) {
	v := ctx.Value(keyKeygen)
	if v == nil {
		return nil, false
	}

	k, ok = v.(keygen.IKeygen)
	return
}

func SetHasher(tx *gorm.DB, options hasher.Options) (h hasher.IHasher, err error) {
	tx.Statement.Context, h, err = SetHasherToContext(tx.Statement.Context, options)
	return
}

func SetHasherToContext(ctx context.Context, options hasher.Options) (newCtx context.Context, h hasher.IHasher, err error) {
	h, err = hasher.New(options)
	if err != nil {
		return
	}
	newCtx = context.WithValue(ctx, keyHasher, h)
	return
}

func GetHasher(tx *gorm.DB) (h hasher.IHasher, ok bool) {
	return GetHasherFromContext(tx.Statement.Context)
}

func GetHasherFromContext(ctx context.Context) (h hasher.IHasher, ok bool) {
	v := ctx.Value(keyHasher)
	if v == nil {
		return nil, false
	}

	h, ok = v.(hasher.IHasher)
	return
}

func SetCipher(tx *gorm.DB, options cipher.Options) (c cipher.ICipher, err error) {
	tx.Statement.Context, c, err = SetCipherToContext(tx.Statement.Context, options)
	return
}

func SetCipherToContext(ctx context.Context, options cipher.Options) (newCtx context.Context, c cipher.ICipher, err error) {
	c, err = cipher.New(options)
	if err != nil {
		return
	}
	newCtx = context.WithValue(ctx, keyCipher, c)
	return
}

func GetCipher(tx *gorm.DB) (c cipher.ICipher, ok bool) {
	return GetCipherFromContext(tx.Statement.Context)
}

func GetCipherFromContext(ctx context.Context) (c cipher.ICipher, ok bool) {
	v := ctx.Value(keyCipher)
	if v == nil {
		return nil, false
	}

	c, ok = v.(cipher.ICipher)
	return
}
