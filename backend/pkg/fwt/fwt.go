package fwt

import (
	"encoding/base64"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/apache/fory/go/fory"
	"majinyao.cn/my-app/backend/pkg/codegen"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher"
	"majinyao.cn/my-app/backend/pkg/snowflake"
)

type IFwt[T IdentityGetter] interface {
	Generate(data T) (accessToken string, refreshToken string, expiredAt time.Time, err error)
	Refresh(claims Claims[T], refreshToken string, handlers ...func(data T) (newData T, err error)) (newAccessToken string, newRefreshToken string, newExpiredAt time.Time, err error)
	ValidateClaims(claims Claims[T]) (err error)
	ParseAccessToken(accessToken string) (claims Claims[T], err error)
	BuildAccessToken(claims Claims[T]) (accessToken string, err error)
}

func New[T IdentityGetter](options Options, registers ...func(f *fory.Fory) error) (IFwt[T], error) {
	return new(fwt[T]).init(options, registers...)
}

type refreshClaims struct {
	id       int64
	string   string
	expireAt time.Time
}

type fwt[T IdentityGetter] struct {
	snowflake     snowflake.ISnowflake
	hasher        hasher.IHasher
	codegen       codegen.ICodegen
	pool          sync.Pool
	idMap         sync.Map // id => refreshClaims
	issuer        string
	subject       string
	epoch         time.Time
	expiredAge    time.Duration
	refreshAge    time.Duration
	refreshLength int
}

func (f *fwt[T]) Generate(data T) (accessToken string, refreshToken string, expiredAt time.Time, err error) {
	now := time.Now()
	claims := Claims[T]{
		Id:        f.snowflake.Generate(),
		Issuer:    f.issuer,
		Subject:   f.subject,
		IssuedAt:  now,
		ExpiredAt: now.Add(f.expiredAge),
		Data:      data,
	}

	accessToken, err = f.BuildAccessToken(claims)
	if err != nil {
		err = errors.Join(errors.New("fwt generate access token failed"), err)
		return
	}

	expiredAt = claims.ExpiredAt
	refreshToken = f.codegen.Generate(f.refreshLength)
	f.idMap.Store(data.GetIdentity(), &refreshClaims{
		id:       claims.Id,
		string:   refreshToken,
		expireAt: now.Add(f.refreshAge),
	})
	return
}

func (f *fwt[T]) Refresh(claims Claims[T], refreshToken string, handlers ...func(data T) (newData T, err error)) (
	newAccessToken string, newRefreshToken string, newExpiredAt time.Time, err error,
) {
	value, ok := f.idMap.Load(claims.Data.GetIdentity())
	if !ok {
		err = errors.New("fwt data identity invalid " + claims.Data.GetIdentity())
		return
	}

	now := time.Now()
	rClaims := value.(*refreshClaims)
	if rClaims.id != claims.Id {
		err = errors.New("fwt claims id not matched")
		return
	}
	if rClaims.string != refreshToken {
		err = errors.New("fwt refresh token wrong")
		return
	}
	if rClaims.expireAt.Before(now) {
		f.idMap.CompareAndDelete(claims.Data.GetIdentity(), value)
		err = errors.New("fwt refresh token expired")
		return
	}

	for _, handler := range handlers {
		if handler != nil {
			claims.Data, err = handler(claims.Data)
			if err != nil {
				err = errors.Join(errors.New("fwt refresh token: handler failed"), err)
				return
			}
		}
	}

	newAccessToken, newRefreshToken, newExpiredAt, err = f.Generate(claims.Data)
	if err != nil {
		err = errors.Join(errors.New("fwt refresh token: generate new token failed"), err)
		return
	}
	return
}

func (f *fwt[T]) ValidateClaims(claims Claims[T]) (err error) {
	return claims.validate(f.issuer, f.subject, f.epoch)
}

func (f *fwt[T]) ParseAccessToken(accessToken string) (claims Claims[T], err error) {
	parts := strings.Split(accessToken, ",")
	if len(parts) != 2 {
		err = errors.New("fwt access token format wrong")
		return
	}

	content, checksum := parts[0], parts[1]
	if !f.hasher.VerifyBase64(checksum, content) {
		err = errors.New("fwt access token checksum failed")
		return
	}

	data, err := base64.RawURLEncoding.DecodeString(content)
	if err != nil {
		err = errors.Join(errors.New("fwt access token decode failed"), err)
		return
	}

	f2 := f.pool.Get().(*fory.Fory)
	defer func() {
		f2.Reset()
		f.pool.Put(f2)
	}()

	err = f2.Unmarshal(data, &claims)
	if err != nil {
		err = errors.Join(errors.New("fwt access token unmarshal failed"), err)
		return
	}
	return
}

func (f *fwt[T]) BuildAccessToken(claims Claims[T]) (accessToken string, err error) {
	f2 := f.pool.Get().(*fory.Fory)
	defer func() {
		f2.Reset()
		f.pool.Put(f2)
	}()

	var data []byte
	data, err = f2.Marshal(claims)
	if err != nil {
		err = errors.Join(errors.New("fwt access token marshal failed"), err)
		return
	}

	content := base64.RawURLEncoding.EncodeToString(data)
	checksum := f.hasher.HashBase64(content)
	return content + "," + checksum, nil
}

func (f *fwt[T]) init(options Options, registers ...func(f *fory.Fory) error) (*fwt[T], error) {
	snowflake, err := snowflake.New(options.Snowflake)
	if err != nil {
		return nil, errors.Join(errors.New("fwt init snowflake failed"), err)
	}

	hasher, err := hasher.New(options.Hasher)
	if err != nil {
		return nil, errors.Join(errors.New("fwt init hasher failed"), err)
	}

	codegen, err := codegen.New(options.Codegen)
	if err != nil {
		return nil, errors.Join(errors.New("fwt init codegen failed"), err)
	}

	f2 := fory.NewFory(true)

	var t T
	tTag := reflect.TypeOf(t).Name()
	if err := f2.RegisterTagType(tTag, t); err != nil {
		return nil, errors.Join(fmt.Errorf("fwt register tag type %s failed", tTag), err)
	}

	var c Claims[T]
	cTag := reflect.TypeOf(c).Name()
	if err := f2.RegisterTagType(cTag, c); err != nil {
		return nil, errors.Join(fmt.Errorf("fwt register tag type %s failed", cTag), err)
	}

	for _, register := range registers {
		if err := register(f2); err != nil {
			return nil, errors.Join(errors.New("fwt register extra failed"), err)
		}
	}

	f.pool = sync.Pool{
		New: func() any {
			fory := fory.NewFory(true)
			fory.RegisterTagType(tTag, t)
			fory.RegisterTagType(cTag, c)
			for _, register := range registers {
				register(fory)
			}
			return fory
		},
	}
	f.pool.Put(f2)

	f.snowflake = snowflake
	f.hasher = hasher
	f.codegen = codegen
	f.issuer = options.Issuer
	f.subject = options.Subject
	f.epoch = options.Epoch
	f.expiredAge = time.Duration(options.ExpiredAge) * time.Second
	f.refreshAge = time.Duration(options.RefreshAge) * time.Second
	f.refreshLength = options.RefreshLength
	return f, nil
}
