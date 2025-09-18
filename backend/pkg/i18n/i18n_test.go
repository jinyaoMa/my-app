package i18n_test

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"majinyao.cn/my-app/backend/pkg/i18n"
	"majinyao.cn/my-app/backend/pkg/test"
)

func TestI18n(t *testing.T) {
	now := time.Now()
	defer test.LogTestingTime(t, now)

	pwd, _ := os.Getwd()
	apath := filepath.Join(pwd, "testdata")
	m, err := i18n.New(i18n.Options{
		Fallback:  "zh-CN",
		Directory: apath,
	})
	if err != nil {
		t.Fatal(err)
	}

	translation := m.GetTranslation()
	msg := translation.Get("Test", "123")
	if msg != "测试123" {
		t.Fatal(`msg != "测试123"`)
	}

	_ = m.Watch(func(value i18n.Translation) (err error) {
		if value.Locale.File != "en-US.json" {
			t.Fatal(`value.Locale.File != "en-US.json"`)
		}
		return
	})

	err = m.SetLocale(m.AvailableLocales()[1].Code, true)
	if err != nil {
		t.Fatal(err)
	}

	translation = m.GetTranslation()
	msg = translation.Get("Test", "123")
	if msg != "Test 123" {
		t.Fatal(`msg != "Test 123"`)
	}
}
