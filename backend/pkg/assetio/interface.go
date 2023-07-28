package assetio

import "io/fs"

type Interface[TI18n any] interface {
	fs.FS

	GetBytes(paths ...string) (data []byte)

	LoadJSON(v interface{}, paths ...string) (ok bool)

	WalkDir(callback func(path string, isDir bool, entry fs.DirEntry) (err error), paths ...string) (err error)

	LoadI18n(v TI18n, paths ...string) (availLangs []*Lang, translationMap map[string]TI18n)
}
