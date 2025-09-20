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

func SetSnowflake(db *gorm.DB, options snowflake.Options) (s snowflake.ISnowflake, err error) {
	db.Statement.Context, s, err = SetSnowflakeToContext(db.Statement.Context, options)
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

func GetSnowflake(db *gorm.DB) (s snowflake.ISnowflake, ok bool) {
	return GetSnowflakeFromContext(db.Statement.Context)
}

func GetSnowflakeFromContext(ctx context.Context) (s snowflake.ISnowflake, ok bool) {
	v := ctx.Value(keySnowflake)
	if v == nil {
		return nil, false
	}

	s, ok = v.(snowflake.ISnowflake)
	return
}

func SetKeygen(db *gorm.DB, options keygen.Options) (k keygen.IKeygen, err error) {
	db.Statement.Context, k, err = SetKeygenToContext(db.Statement.Context, options)
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

func GetKeygen(db *gorm.DB) (k keygen.IKeygen, ok bool) {
	return GetKeygenFromContext(db.Statement.Context)
}

func GetKeygenFromContext(ctx context.Context) (k keygen.IKeygen, ok bool) {
	v := ctx.Value(keyKeygen)
	if v == nil {
		return nil, false
	}

	k, ok = v.(keygen.IKeygen)
	return
}

func SetHasher(db *gorm.DB, options hasher.Options) (h hasher.IHasher, err error) {
	db.Statement.Context, h, err = SetHasherToContext(db.Statement.Context, options)
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

func GetHasher(db *gorm.DB) (h hasher.IHasher, ok bool) {
	return GetHasherFromContext(db.Statement.Context)
}

func GetHasherFromContext(ctx context.Context) (h hasher.IHasher, ok bool) {
	v := ctx.Value(keyHasher)
	if v == nil {
		return nil, false
	}

	h, ok = v.(hasher.IHasher)
	return
}

func SetCipher(db *gorm.DB, options cipher.Options) (c cipher.ICipher, err error) {
	db.Statement.Context, c, err = SetCipherToContext(db.Statement.Context, options)
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

func GetCipher(db *gorm.DB) (c cipher.ICipher, ok bool) {
	return GetCipherFromContext(db.Statement.Context)
}

func GetCipherFromContext(ctx context.Context) (c cipher.ICipher, ok bool) {
	v := ctx.Value(keyCipher)
	if v == nil {
		return nil, false
	}

	c, ok = v.(cipher.ICipher)
	return
}
