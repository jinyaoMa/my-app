package i18n

import (
	"embed"
	"io/fs"
	"strings"

	"my-app/backend.new/utils"
)

//go:embed translation
var translation embed.FS

const (
	TranslationSuffix  = ".json"
	TranslationDirname = "translation"
)

type I18n struct {
	assetHelper        utils.AssetHelper
	translationMap     map[string]*Translation
	availableLanguages []string
}

func NewI18n(dirLanguages string) *I18n {
	var assetHelper utils.AssetHelper
	if utils.Utils().HasDir(dirLanguages) {
		// Languages Directory available at executable directory
		assetHelper = utils.NewDirFS(dirLanguages)
	} else {
		assetHelper = utils.NewEmbedFS(translation, TranslationDirname)
		assetHelper.Extract(dirLanguages)
	}

	return &I18n{
		assetHelper:    assetHelper,
		translationMap: make(map[string]*Translation),
	}
}

// Translation get translation of the given language
func (i *I18n) Translation(lang string) *Translation {
	if i.translationMap[lang] == nil { // cache
		i.translationMap[lang] = &Translation{}
		err := i.assetHelper.LoadJSON(i.translationMap[lang], lang+".json")
		if err != nil {
			return TranslationPlaceholder()
		}
	}
	return i.translationMap[lang]
}

// HasLanguage check if the language is available
func (i *I18n) HasLanguage(lang string) bool {
	for _, l := range i.AvailableLanguages() {
		if l == lang {
			return true
		}
	}
	return false
}

// AvailableLanguages get available languages
func (i *I18n) AvailableLanguages() []string {
	if i.availableLanguages == nil { // cache
		i.availableLanguages = make([]string, 0, 18)
		err := i.assetHelper.Walk(func(path string, isDir bool, f fs.DirEntry) error {
			if strings.HasSuffix(f.Name(), TranslationSuffix) && !isDir {
				i.availableLanguages = append(i.availableLanguages, strings.TrimSuffix(f.Name(), TranslationSuffix))
			}
			return nil
		})
		if err != nil {
			return nil
		}
	}
	return i.availableLanguages
}
