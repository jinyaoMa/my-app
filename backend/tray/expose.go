package tray

import (
	"my-app/backend/app/config"
	"my-app/backend/web"
)

func (t *tray) IsWebServiceRunning() bool {
	return web.Web().IsRunning()
}

func (t *tray) StartWebService() {
	t.webService.FalseOptions[MniWebServiceStart].ClickedCh <- struct{}{}
}

func (t *tray) OpenVitePress() {
	t.webService.TrueOptions[MniWebServiceVitePress].ClickedCh <- struct{}{}
}

func (t *tray) OpenSwagger() {
	t.webService.TrueOptions[MniWebServiceSwagger].ClickedCh <- struct{}{}
}

func (t *tray) StopWebService() {
	t.webService.TrueOptions[MniWebServiceStop].ClickedCh <- struct{}{}
}

func (t *tray) ChangeDisplayLanguage(lang string) {
	switch lang {
	default:
		t.displayLanguage.Options[config.DisplayLanguageEn].ClickedCh <- struct{}{}
	case config.DisplayLanguageZh:
		t.displayLanguage.Options[config.DisplayLanguageZh].ClickedCh <- struct{}{}
	}
}

func (t *tray) ChangeColorTheme(theme string) {
	switch theme {
	default:
		t.colorTheme.Options[config.ColorThemeSystem].ClickedCh <- struct{}{}
	case config.ColorThemeLight:
		t.colorTheme.Options[config.ColorThemeLight].ClickedCh <- struct{}{}
	case config.ColorThemeDark:
		t.colorTheme.Options[config.ColorThemeDark].ClickedCh <- struct{}{}
	}
}
