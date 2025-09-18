package codegen

import (
	"crypto/rand"
	"errors"
	"math/big"
	r2 "math/rand"
	"strings"
)

type ICodegen interface {
	Generate(length int) string
}

func MustNew(options Options) ICodegen {
	c, err := New(options)
	if err != nil {
		panic(err)
	}
	return c
}

func New(options Options) (ICodegen, error) {
	return new(codegen).init(options)
}

type codegen struct {
	characters    []rune
	lenCharacters *big.Int
}

func (c *codegen) Generate(length int) string {
	if length <= 0 {
		return ""
	}

	var builder strings.Builder
	for range length {
		n, _ := rand.Int(rand.Reader, c.lenCharacters)
		if n == nil {
			m := r2.Int63n(c.lenCharacters.Int64())
			builder.WriteRune(c.characters[m])
		} else {
			builder.WriteRune(c.characters[n.Int64()])
		}
	}
	return builder.String()
}

func (c *codegen) init(options Options) (*codegen, error) {
	c.characters = []rune(options.Characters)
	length := len(c.characters)
	if length == 0 {
		return nil, errors.New("characters length must be greater than 0")
	}
	c.lenCharacters = big.NewInt(int64(length))
	return c, nil
}
