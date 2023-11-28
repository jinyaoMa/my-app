package base

import "dario.cat/mergo"

func SimpleMerge[T any](src T, dst T) T {
	err := mergo.Merge(dst, src)
	if err != nil {
		return src
	}
	return dst
}
