package i18n

import (
	"embed"
	"io/fs"
	"strings"

	"my-app/backend/pkg/utils"
)

//go:embed translation
var translation embed.FS

type I18n struct {
	embedFs            *utils.EmbedFs
	translationMap     map[string]*Translation
	availableLanguages []string
}

func NewI18n() *I18n {
	return &I18n{
		embedFs:        utils.NewEmbedFs(translation, "translation"),
		translationMap: make(map[string]*Translation),
	}
}

func (i *I18n) Translation(lang string) *Translation {
	if !i.HasLanguage(lang) {
		return nil
	}
	if i.translationMap[lang] == nil {
		i.translationMap[lang] = &Translation{}
		i.embedFs.LoadJSON(i.translationMap[lang], lang+".json")
	}
	return i.translationMap[lang]
}

func (i *I18n) HasLanguage(lang string) bool {
	for _, l := range i.AvailableLanguages() {
		if l == lang {
			return true
		}
	}
	return false
}

func (i *I18n) AvailableLanguages() []string {
	if i.availableLanguages == nil {
		i.availableLanguages = make([]string, 0, 18)
		i.embedFs.WalkDir(func(path string, isDir bool, f fs.DirEntry) error {
			i.availableLanguages = append(i.availableLanguages, strings.TrimSuffix(f.Name(), ".json"))
			return nil
		})
	}
	return i.availableLanguages
}

func (i *I18n) WalkTranslations(callback func(translation []byte)) error {
	return i.embedFs.WalkDir(func(path string, isDir bool, f fs.DirEntry) error {
		data, err := i.embedFs.GetFileBytes(path)
		if err != nil {
			return err
		}
		callback(data)
		return nil
	})
}
