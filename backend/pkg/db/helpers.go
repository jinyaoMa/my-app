package db

import (
	"context"
	"time"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/crypto/cipher"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher"
	"majinyao.cn/my-app/backend/pkg/crypto/keygen"
	"majinyao.cn/my-app/backend/pkg/snowflake"
	"majinyao.cn/my-app/backend/pkg/utils"
)

type key int

const (
	keySnowflake key = iota + 1
	keyKeygen
	keyHasher
	keyCipher
)

func setSnowflake(tx *gorm.DB, options snowflake.Options) (s snowflake.ISnowflake, err error) {
	tx.Statement.Context, s, err = setSnowflakeToContext(tx.Statement.Context, options)
	return
}

func setSnowflakeToContext(ctx context.Context, options snowflake.Options) (newCtx context.Context, s snowflake.ISnowflake, err error) {
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

func setKeygen(tx *gorm.DB, options keygen.Options) (k keygen.IKeygen, err error) {
	tx.Statement.Context, k, err = setKeygenToContext(tx.Statement.Context, options)
	return
}

func setKeygenToContext(ctx context.Context, options keygen.Options) (newCtx context.Context, k keygen.IKeygen, err error) {
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

func setHasher(tx *gorm.DB, options hasher.Options) (h hasher.IHasher, err error) {
	tx.Statement.Context, h, err = setHasherToContext(tx.Statement.Context, options)
	return
}

func setHasherToContext(ctx context.Context, options hasher.Options) (newCtx context.Context, h hasher.IHasher, err error) {
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

func setCipher(tx *gorm.DB, options cipher.Options) (c cipher.ICipher, err error) {
	tx.Statement.Context, c, err = setCipherToContext(tx.Statement.Context, options)
	return
}

func setCipherToContext(ctx context.Context, options cipher.Options) (newCtx context.Context, c cipher.ICipher, err error) {
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

// session with new timer context which will be deadline exceeded after the given timeout
func SectionWithTimeout(tx *gorm.DB, timeout time.Duration) (*gorm.DB, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(tx.Statement.Context, timeout)
	return tx.Session(&gorm.Session{
		NewDB:   true,
		Context: ctx,
	}), cancel
}

// session with new cancel context which can be canceled after calling cancel function
func SectionWithCancelCause(tx *gorm.DB) (*gorm.DB, context.CancelCauseFunc) {
	ctx, cancel := context.WithCancelCause(tx.Statement.Context)
	return tx.Session(&gorm.Session{
		NewDB:   true,
		Context: ctx,
	}), cancel
}

// session with new timer context which will be deadline exceeded after the given timeout,
// but the timeout of the new timer context won't affect the given context
func SectionUnderContextWithTimeout(ctx context.Context, tx *gorm.DB, timeout time.Duration) (*gorm.DB, context.CancelFunc) {
	tx, cancel := SectionWithTimeout(tx, timeout)
	stop := context.AfterFunc(ctx, cancel)
	return tx, func() {
		stop()
		cancel()
	}
}

// session with new cancel context which can be canceled by the given context,
// but the cancellation of the new cancel context won't affect the given context
func SectionUnderContextWithCancel(ctx context.Context, tx *gorm.DB) (*gorm.DB, context.CancelFunc) {
	tx, cancel := SectionWithCancelCause(tx)
	stop := context.AfterFunc(ctx, func() {
		cancel(context.Cause(ctx))
	})
	return tx, func() {
		stop()
		cancel(context.Canceled)
	}
}

func ConvertStringToId(id string) (int64, error) {
	return utils.ConvertHexToInt64(id)
}

func ConvertIdToString(id int64) string {
	return utils.ConvertInt64ToHex(id)
}
