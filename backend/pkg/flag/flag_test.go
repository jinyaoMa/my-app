package flag_test

import (
	"bytes"
	"testing"
	"time"

	"majinyao.cn/my-app/backend/pkg/flag"
	"majinyao.cn/my-app/backend/pkg/test"
)

func TestFlag(t *testing.T) {
	now := time.Now()
	defer test.LogTestingTime(t, now)

	f := flag.Make(0, true)

	if !bytes.Equal(f.ToBytes(), []byte{}) {
		t.Fatal("expected true, got false")
	}

	f = flag.Make(1, true)

	if !bytes.Equal(f.ToBytes(), []byte{0b10000000}) {
		t.Fatal("expected true, got false")
	}

	f = flag.Make(7, true)

	if !bytes.Equal(f.ToBytes(), []byte{0b11111110}) {
		t.Fatal("expected true, got false")
	}

	f = flag.Make(8, true)

	if !bytes.Equal(f.ToBytes(), []byte{0b11111111}) {
		t.Fatal("expected true, got false")
	}

	f = flag.Make(9, true)

	if !bytes.Equal(f.ToBytes(), []byte{0b11111111, 0b10000000}) {
		t.Fatal("expected true, got false")
	}

	f = flag.Make(15, true)

	if !bytes.Equal(f.ToBytes(), []byte{0b11111111, 0b11111110}) {
		t.Fatal("expected true, got false")
	}

	f = flag.Make(16, true)

	if !bytes.Equal(f.ToBytes(), []byte{0b11111111, 0b11111111}) {
		t.Fatal("expected true, got false")
	}

	f = flag.FromBytes([]byte{0, 1, 0})

	if f.IsOn(1) {
		t.Fatal("expected false, got true")
	}

	if f.IsOn(8) {
		t.Fatal("expected false, got true")
	}

	if !f.IsOn(15) {
		t.Fatal("expected true, got false")
	}

	f = f.TurnOn(0)
	if !f.IsOn(0) {
		t.Fatal("expected true, got false")
	}

	f = f.TurnOn(8)
	if !f.IsOn(8) {
		t.Fatal("expected true, got false")
	}

	f2 := flag.FromBytes(f.ToBytes())
	if !f.Equals(f2) {
		t.Fatal("expected true, got false")
	}

	f2 = f2.TurnOff(15)
	if f.Equals(f2) {
		t.Fatal("expected false, got true")
	}

	f2 = flag.FromBytes([]byte{0, 0, 0b10000000})
	f3 := f.Or(f2)
	if f3.Equals(f) {
		t.Fatal("expected false, got true")
	}

	f3 = f.And(f2)
	if f3.Equals(f) {
		t.Fatal("expected false, got true")
	}

	f2 = f2.TurnOn(0)
	if !f.Or(f2).HasOn() {
		t.Fatal("expected true, got false")
	}

	f2 = f2.TurnOn(8)
	f2 = f2.TurnOn(15)
	if !f.And(f2).HasOn() {
		t.Fatal("expected true, got false")
	}
}
