package i18n

import (
	"embed"
	"io/fs"
	"strings"

	"my-app/backend/utils"
)

//go:embed translations
var translations embed.FS

type I18n struct {
	log                *utils.Logger
	assetHelper        utils.AssetHelper
	translationMap     map[string]*Translation
	availableLanguages []string
}

func NewI18n(dirLanguages string, log *utils.Logger) *I18n {
	// setup asset helper
	var assetHelper utils.AssetHelper
	if utils.Utils().HasDir(dirLanguages) {
		// Languages Directory available at executable directory
		assetHelper = utils.NewDirFS(dirLanguages)
		log.Printf("I18N LOAD ASSET FROM dirLanguages: %s\n", dirLanguages)
	} else {
		assetHelper = utils.NewEmbedFS(translations, "translations")
		if err := assetHelper.Extract(dirLanguages); err != nil {
			log.Fatalf("failed to extract embed translations into dirLanguages (%s): %+v\n", dirLanguages, err)
		}
		log.Println("I18N LOAD ASSET FROM embed: backend/app/i18n/translations")
	}

	// load translations
	translationMap := make(map[string]*Translation)
	countAvailableLanguages := 0
	if err := assetHelper.Walk(func(path string, isDir bool, f fs.DirEntry) error {
		if strings.HasSuffix(f.Name(), ".json") && !isDir {
			t := &Translation{}
			if err := assetHelper.LoadJSON(t, path); err != nil { // load translation from json file
				return err
			}
			translationMap[t.Lang.Code] = t
			countAvailableLanguages++
		}
		return nil
	}); err != nil {
		log.Fatalf("failed to load i18n: %+v\n", err)
	}

	// fill available languages
	availableLanguages := make([]string, 0, countAvailableLanguages)
	for lang := range translationMap {
		availableLanguages = append(availableLanguages, lang)
	}

	return &I18n{
		log:                log,
		assetHelper:        assetHelper,
		translationMap:     translationMap,
		availableLanguages: availableLanguages,
	}
}

// Translation get translation of the given language
func (i *I18n) Translation(lang string) *Translation {
	if t, ok := i.translationMap[lang]; ok {
		// if the language is available
		return t
	}
	return &TranslationPlaceholder
}

// ParseLanguage get valid language
func (i *I18n) ParseLanguage(lang string) string {
	if _, ok := i.translationMap[lang]; ok {
		return lang
	}
	if len(i.availableLanguages) > 0 {
		// default language
		return i.availableLanguages[0]
	}
	return LanguagePlaceholder
}

// HasLanguage check if language is available
func (i *I18n) HasLanguage(lang string) bool {
	_, ok := i.translationMap[lang]
	return ok
}

// AvailableLanguages get available languages
func (i *I18n) AvailableLanguages() []string {
	return i.availableLanguages
}
