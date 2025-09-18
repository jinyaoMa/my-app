package i18n

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"sync/atomic"

	"github.com/tidwall/gjson"
	"majinyao.cn/my-app/backend/pkg/reactive"
)

type II18n interface {
	SetLocale(code string, sync ...bool) (err error)
	GetTranslation() (t Translation)
	Watch(handler func(t Translation) (err error)) (id int64)
	Unwatch(id int64) (handler func(t Translation) (err error), ok bool)
	AvailableLocales() []Locale
	Reload() error
}

type i18n struct {
	fallback           string
	directory          string
	defineJson         string
	locale             reactive.ITransformable[Locale, Translation]
	currentTranslation atomic.Pointer[Translation]
	availableLocales   []Locale
}

func New(options Options) (II18n, error) {
	return new(i18n).init(options)
}

func (i *i18n) SetLocale(code string, sync ...bool) (err error) {
	index := slices.IndexFunc(i.availableLocales, func(l Locale) bool {
		return l.Code == code
	})
	if index < 0 {
		return fmt.Errorf("locale [%s] not found", code)
	}
	return i.locale.Set(i.availableLocales[index], sync...)
}

func (i *i18n) GetTranslation() (t Translation) {
	return *i.currentTranslation.Load()
}

func (i *i18n) Watch(handler func(t Translation) (err error)) (id int64) {
	return i.locale.Watch(handler)
}

func (i *i18n) Unwatch(id int64) (handler func(t Translation) (err error), ok bool) {
	return i.locale.Unwatch(id)
}

func (i *i18n) AvailableLocales() []Locale {
	return i.availableLocales
}

func (i *i18n) Reload() error {
	defineJsonPath := filepath.Join(i.directory, i.defineJson)
	data, err := os.ReadFile(defineJsonPath)
	if err != nil {
		return errors.Join(fmt.Errorf("error when reading define.json from %s", defineJsonPath), err)
	}

	err = json.Unmarshal(data, &i.availableLocales)
	if err != nil {
		return errors.Join(errors.New("error when unmarshaling define.json"), err)
	}
	if len(i.availableLocales) == 0 {
		return errors.Join(errors.New("no available locales found"), err)
	}

	var fallback Locale
	if index := slices.IndexFunc(i.availableLocales, func(l Locale) bool {
		return l.Code == i.fallback
	}); index >= 0 {
		fallback = i.availableLocales[index]
	}

	i.locale, err = reactive.NewTransformable(i.availableLocales[0], func(locale Locale) (translation Translation, err error) {
		translation.Locale = locale

		var data []byte
		data, err = os.ReadFile(filepath.Join(i.directory, locale.File))
		if err != nil {
			data, err = os.ReadFile(filepath.Join(i.directory, fallback.File))
			if err != nil {
				return
			}
		}

		translation.result = gjson.ParseBytes(data)
		return
	}, i.currentTranslationListener)
	if err != nil {
		return errors.Join(errors.New("error when creating locale transformable"), err)
	}

	i.locale.Filter(func(locale Locale) (ok bool) {
		return slices.ContainsFunc(i.availableLocales, func(l Locale) bool {
			return l.Code == locale.Code
		})
	})
	i.locale.Watch(i.currentTranslationListener)
	return nil
}

func (i *i18n) currentTranslationListener(translation Translation) (err error) {
	i.currentTranslation.Store(&translation)
	return
}

func (i *i18n) init(options Options) (*i18n, error) {
	i.fallback = options.Fallback
	i.directory = options.Directory
	i.defineJson = options.DefineJson
	if i.defineJson == "" {
		i.defineJson = "define.json"
	}

	err := i.Reload()
	if err != nil {
		return nil, err
	}
	return i, nil
}
