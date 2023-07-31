package assetsio

import (
	"io/fs"
	"path/filepath"
)

type Lang struct {
	Code string `json:"code"` // language code must be the same as the json source filename without extension
	Text string `json:"text"`
}

type I18n[TTranslation ITranslation] struct {
	Interface
}

// LoadI18n implements II18n.
func (a *I18n[TTranslation]) LoadI18n() (availLangs []*Lang, translationMap map[string]TTranslation) {
	var v TTranslation
	a.WalkDir(func(path string, isDir bool, entry fs.DirEntry) (err error) {
		if filepath.Ext(path) == ".json" && a.LoadJSON(v, path) {
			lang := v.Lang()
			availLangs = append(availLangs, lang)
			translationMap[lang.Code] = v
		}
		return nil
	})
	return
}

func NewI18n[TTranslation ITranslation](root string) II18n[TTranslation] {
	return &I18n[TTranslation]{
		Interface: New(root),
	}
}
