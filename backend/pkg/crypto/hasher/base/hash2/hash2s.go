package hash2

import (
	"bytes"

	"golang.org/x/sync/errgroup"
	"majinyao.cn/my-app/backend/pkg/utils"
)

type Hash2s []IHash2

func (hs Hash2s) Algs() []string {
	return utils.SliceMap(hs, func(h IHash2) string {
		return h.Alg()
	})
}

func (hs Hash2s) Serials() []string {
	return utils.SliceMap(hs, func(h IHash2) string {
		return h.Serial()
	})
}

func (hs Hash2s) CheckSerials(serials []string) bool {
	if len(hs) != len(serials) {
		return false
	}
	for i := range hs {
		if !hs[i].CheckSerial(serials[i]) {
			return false
		}
	}
	return true
}

func (hs Hash2s) Write(p []byte) (n int, err error) {
	var g errgroup.Group
	for i := range hs {
		index := i
		g.Go(func() error {
			_, errWrite := hs[index].Write(p)
			return errWrite
		})
	}
	if err = g.Wait(); err != nil {
		return 0, err
	}
	return len(p), nil
}

func (hs Hash2s) Sum() (sums [][]byte) {
	sums = make([][]byte, 0, len(hs))
	for i := range hs {
		sums = append(sums, hs[i].Sum())
	}
	return
}

func (hs Hash2s) SumBase64() (b64s []string) {
	b64s = make([]string, 0, len(hs))
	for i := range hs {
		b64s = append(b64s, hs[i].SumBase64())
	}
	return
}

func (hs Hash2s) Check(sums [][]byte) (ok bool) {
	newSum := hs.Sum()
	if len(newSum) != len(sums) {
		return false
	}
	for i := range newSum {
		if !bytes.Equal(newSum[i], sums[i]) {
			return false
		}
	}
	return true
}

func (hs Hash2s) CheckBase64(b64s []string) (ok bool) {
	newB64 := hs.SumBase64()
	if len(newB64) != len(b64s) {
		return false
	}
	for i := range newB64 {
		if newB64[i] != b64s[i] {
			return false
		}
	}
	return true
}

func (hs Hash2s) Reset() {
	for i := range hs {
		hs[i].Reset()
	}
}
