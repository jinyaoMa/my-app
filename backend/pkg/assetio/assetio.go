package assetio

import (
	"encoding/json"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

type Assetio[TI18n I18n] struct {
	fs.FS
	root string
}

// LoadI18n implements Interface.
func (a *Assetio[TI18n]) LoadI18n(v TI18n, paths ...string) (availLangs []*Lang, translationMap map[string]TI18n) {
	a.WalkDir(func(path string, isDir bool, entry fs.DirEntry) (err error) {
		if filepath.Ext(path) == ".json" && a.LoadJSON(v, path) {
			lang := v.Lang()
			availLangs = append(availLangs, lang)
			translationMap[lang.Code] = v
		}
		return nil
	}, paths...)
	return
}

// GetBytes implements Interface.
func (a *Assetio[TI18n]) GetBytes(paths ...string) (data []byte) {
	if file, err := a.Open(filepath.Join(paths...)); err == nil {
		defer file.Close()
		if data, err = io.ReadAll(file); err != nil {
			return nil
		}
	}
	return
}

// LoadJSON implements Interface.
func (a *Assetio[TI18n]) LoadJSON(v interface{}, paths ...string) (ok bool) {
	data := a.GetBytes(paths...)
	if data == nil {
		return false
	}
	err := json.Unmarshal(data, v)
	return err == nil
}

// WalkDir implements Interface.
func (a *Assetio[TI18n]) WalkDir(callback func(path string, isDir bool, entry fs.DirEntry) (err error), paths ...string) (err error) {
	err = fs.WalkDir(a, filepath.Join(paths...), func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		return callback(path, d.IsDir(), d)
	})
	return
}

func (a *Assetio[TI18n]) Root() string {
	return a.root
}

func NewAssetio[TI18n I18n](root string) (a Interface[TI18n], err error) {
	return &Assetio[TI18n]{
		FS:   os.DirFS(root),
		root: root,
	}, nil
}
