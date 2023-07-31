package assetsio

import "io/fs"

type Interface interface {
	fs.FS

	GetBytes(paths ...string) (data []byte)

	LoadJSON(v interface{}, paths ...string) (ok bool)

	WalkDir(callback func(path string, isDir bool, entry fs.DirEntry) (err error), paths ...string) (err error)
}

type II18n[TTranslation ITranslation] interface {
	Interface
	LoadTranslation(lang string) (t TTranslation, ok bool)
	LoadI18n() (availLangs []*Lang, translationMap map[string]TTranslation)
}

type ITranslation interface {
	Metadata() *Lang
}
