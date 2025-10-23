package flag

import (
	"bytes"
	"encoding/base64"
	"slices"
)

type IFlag interface {
	ToBytes() []byte
	ToBase64() string
  ToPositions() []int
	Trim() IFlag
	TurnOn(pos int, noCopy ...bool) IFlag
	TurnOff(pos int, noCopy ...bool) IFlag
	IsOn(pos int) bool
	Or(other IFlag) IFlag
	And(other IFlag) IFlag
	HasOn() bool
	Equals(other IFlag) bool
}

func Make(size int, allOn bool) IFlag {
	if size < 1 {
		return make(flag, 0)
	}

	bIndex := (size - 1) / 8
	bLenNeeded := bIndex + 1
	f := make(flag, bLenNeeded)
	if allOn {
		for i := range f {
			f[i] = 0b11111111
			if i == bIndex && size%8 != 0 {
				f[i] = 0b11111111 << (8 - size%8)
			}
		}
	}
	return f
}

func FromBytes(v []byte) IFlag {
	return flag(v).copy()
}

func FromBase64(b64 string) (IFlag, error) {
	v, err := base64.RawURLEncoding.DecodeString(b64)
	if err != nil {
		return nil, err
	}
	return flag(v), nil
}

type flag []byte

func (f flag) ToBytes() []byte {
	return f.copy()
}

func (f flag) ToBase64() string {
	return base64.RawURLEncoding.EncodeToString(f)
}

func (f flag) ToPositions() []int {
	positions := make([]int, 0)
	for i, b := range f {
		if b == 0 {
			continue
		}
		for j := 0; j < 8; j++ {
			if b&(0b10000000>>j) != 0 {
				positions = append(positions, i*8+j)
			}
		}
	}
	return positions
}

func (f flag) Trim() IFlag {
	for i := len(f) - 1; i >= 0; i-- {
		if f[i] == 0 {
			continue
		}
		return f[:i+1]
	}
	return f[:0]
}

func (f flag) TurnOn(pos int, noCopy ...bool) IFlag {
	nc := slices.Contains(noCopy, true)
	if !nc {
		f = f.copy()
	}
	if pos < 0 {
		return f
	}

	bLen := len(f)
	bIndex := pos / 8
	bLenNeeded := bIndex + 1
	if bLen < bLenNeeded {
		f = append(f, make(flag, bLenNeeded-bLen)...)
	}

	bPos := pos % 8
	f[bIndex] |= 0b10000000 >> bPos
	return f
}

func (f flag) TurnOff(pos int, noCopy ...bool) IFlag {
	nc := slices.Contains(noCopy, true)
	if !nc {
		f = f.copy()
	}
	if pos < 0 {
		return f
	}

	bLen := len(f)
	bIndex := pos / 8
	bLenNeeded := bIndex + 1
	if bLen < bLenNeeded {
		f = append(f, make(flag, bLenNeeded-bLen)...)
	}

	bPos := pos % 8
	f[bIndex] &= ^(0b10000000 >> bPos)
	return f
}

func (f flag) IsOn(pos int) bool {
	if pos < 0 {
		return false
	}

	bLen := len(f)
	bIndex := pos / 8
	bLenNeeded := bIndex + 1
	if bLen < bLenNeeded {
		return false
	}

	bPos := pos % 8
	return f[bIndex]&(0b10000000>>bPos) != 0
}

func (f flag) Or(other IFlag) IFlag {
	f = f.copy()
	if other == nil {
		return f
	}

	oBytes := other.ToBytes()
	bLenNeeded := len(oBytes)
	if bLenNeeded == 0 {
		return f
	}

	bLen := len(f)
	if bLen < bLenNeeded {
		f = append(f, make(flag, bLenNeeded-bLen)...)
	}

	for i := range bLenNeeded {
		f[i] |= oBytes[i]
	}
	return f
}

func (f flag) And(other IFlag) IFlag {
	f = f.copy()
	if other == nil {
		return f
	}

	oBytes := other.ToBytes()
	bLenNeeded := len(oBytes)
	if bLenNeeded == 0 {
		return f
	}

	bLen := len(f)
	if bLen > bLenNeeded {
		f = f[:bLenNeeded]
	} else if bLen < bLenNeeded {
		bLenNeeded = bLen
	}

	for i := range bLenNeeded {
		f[i] &= oBytes[i]
	}
	return f
}

func (f flag) HasOn() bool {
	for _, b := range f {
		if b != 0 {
			return true
		}
	}
	return false
}

func (f flag) Equals(other IFlag) bool {
	return bytes.Equal(f.Trim().ToBytes(), other.Trim().ToBytes())
}

func (f flag) copy() flag {
	c := make(flag, len(f))
	copy(c, f)
	return c
}
