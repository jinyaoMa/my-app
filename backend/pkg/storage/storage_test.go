package storage_test

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"slices"
	"testing"
	"time"

	"majinyao.cn/my-app/backend/pkg/crypto/hasher"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/crc64"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/sha3"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/xxh3"
	"majinyao.cn/my-app/backend/pkg/storage"
	"majinyao.cn/my-app/backend/pkg/test"
)

func TestStorage(t *testing.T) {
	now := time.Now()
	defer test.LogTestingTime(t, now)

	ctx := context.TODO()
	ns := "mjy"
	pwd, _ := os.Getwd()
	apath := filepath.Join(pwd, "testdata")
	options := storage.Options{
		Libraries: []storage.Library{
			{
				Mountpoint: filepath.VolumeName(apath),
				Directory:  apath,
			},
		},
		Temporary:  "tmp",
		BufferSize: 512 * storage.KB,
		Hashers: []hasher.Options{
			{
				Alg:       sha3.Alg,
				BitLength: 256,
			},
			{
				Alg: crc64.Alg,
			},
			{
				Alg: xxh3.Alg,
			},
		},
		MaxPathLength: 254,
	}
	defer os.RemoveAll(apath)

	c, err := hasher.NewCombo(options.Hashers...)
	if err != nil {
		t.Fatal(err)
	}

	checksums := c.HashBase64(ns)

	s, err := storage.New(options)
	if err != nil {
		t.Fatal(err)
	}

	err = s.CacheTemporary(ctx, ns, checksums[0], 3, 2, bytes.NewReader([]byte(ns)[2:]))
	if err != nil {
		t.Fatal(err)
	}

	err = s.CacheTemporary(ctx, ns, checksums[0], 3, 0, bytes.NewReader([]byte(ns)[0:2]))
	if err != nil {
		t.Fatal(err)
	}

	cs, err := s.PersistTemporary(ctx, ns, checksums[0], 3, ".txt")
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(cs, checksums) {
		t.Fatal("checksums not equal")
	}
}
