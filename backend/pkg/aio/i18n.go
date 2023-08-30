package aio

import (
	"io/fs"
	"path/filepath"
)

type Lang struct {
	Code string `json:"code"` // language code must be the same as the json source filename without extension
	Text string `json:"text"`
}

type I18n[TTranslation ITranslation] struct {
	*AIO
}

// CheckLang implements II18n.
func (a *I18n[TTranslation]) LoadTranslation(lang string) (t TTranslation, ok bool) {
	a.WalkDir(func(path string, isDir bool, entry fs.DirEntry) (err error) {
		if !isDir && filepath.Base(path) == lang+".json" && a.LoadJSON(&t, path) {
			ok = true
			return fs.SkipAll
		}
		return nil
	})
	return
}

// LoadI18n implements II18n.
func (a *I18n[TTranslation]) LoadI18n() (availLangs []Lang, translationMap map[string]TTranslation) {
	var v TTranslation
	err := a.WalkDir(func(path string, isDir bool, entry fs.DirEntry) (err error) {
		println("=========================LoadI18n")
		println(isDir)
		println(filepath.Ext(path))
		if !isDir && filepath.Ext(path) == ".json" && a.LoadJSON(&v, path) {
			lang := v.Metadata()
			availLangs = append(availLangs, lang)
			translationMap[lang.Code] = v
		}
		return nil
	})
	println(err)
	return
}

func NewI18n[TTranslation ITranslation](root string) *I18n[TTranslation] {
	return &I18n[TTranslation]{
		AIO: New(root),
	}
}

func NewII18n[TTranslation ITranslation](root string) II18n[TTranslation] {
	return NewI18n[TTranslation](root)
}
