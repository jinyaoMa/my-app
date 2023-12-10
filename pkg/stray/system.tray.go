package stray

import (
	"context"
	"my-app/pkg/i18n"
	"sync"

	"github.com/getlantern/systray"
)

var (
	instance any
	once     sync.Once
)

type systemTray[TTranslation i18n.ITranslation] struct {
	options     *SystemTrayOptions[TTranslation]
	translation TTranslation
	ctx         context.Context
}

func (tray *systemTray[TTranslation]) Run() {
	systray.Run(tray.onReady, nil)
}

func (tray *systemTray[TTranslation]) Register() {
	systray.Register(tray.onReady, nil)
}

func (tray *systemTray[TTranslation]) SetContext(ctx context.Context) {
	tray.ctx = ctx
}

func (tray *systemTray[TTranslation]) Update() {
	tray.update()
}

func (tray *systemTray[TTranslation]) SetTranslation(translation TTranslation) {
	tray.translation = translation
	tray.update()
}

func (tray *systemTray[TTranslation]) SetOptions(options *SystemTrayOptions[TTranslation]) (err error) {
	options, err = NewSystemTrayOptions(options)
	if err != nil {
		return err
	}
	tray.options = options
	return
}

func SystemTray[TTranslation i18n.ITranslation]() (tray *systemTray[TTranslation], ok bool) {
	once.Do(func() {
		instance = &systemTray[TTranslation]{}
	})
	tray, ok = instance.(*systemTray[TTranslation])
	return
}
