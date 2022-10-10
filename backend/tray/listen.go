package tray

import (
	"fmt"
	"my-app/backend/app"
	"my-app/backend/app/config"
	"my-app/backend/web"

	"github.com/getlantern/systray"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// need to run tray click listener in goroutine
func (t *tray) listen() {
	cfg := app.App().Config()
	for {
		select {
		case <-t.openWindow.ClickedCh:
			runtime.Show(t.ctx)
		case <-t.webService.FalseOptions[MniWebServiceStart].ClickedCh:
			if web.Web().Start() {
				runtime.EventsEmit(t.ctx, "onWebServiceStart")
				t.webService.SetFlag(true)
				t.refreshTooltip()
			}
		case <-t.webService.TrueOptions[MniWebServiceVitePress].ClickedCh:
			runtime.BrowserOpenURL(
				t.ctx,
				fmt.Sprintf("https://localhost%s/docs/", cfg.Web().PortHttps),
			)
		case <-t.webService.TrueOptions[MniWebServiceSwagger].ClickedCh:
			runtime.BrowserOpenURL(
				t.ctx,
				fmt.Sprintf("https://localhost%s/swagger/index.html", cfg.Web().PortHttps),
			)
		case <-t.webService.TrueOptions[MniWebServiceStop].ClickedCh:
			if web.Web().Stop() {
				runtime.EventsEmit(t.ctx, "onWebServiceStop")
				t.webService.SetFlag(false)
				t.refreshTooltip()
			}
		case <-t.displayLanguage.Options[config.DisplayLanguageEn].ClickedCh:
			cfg.Update(config.CfgDisplayLanguage, config.DisplayLanguageEn)
			runtime.EventsEmit(t.ctx, "onDisplayLanguageChanged", cfg.DisplayLanguage)
			t.displayLanguage.Select(cfg.DisplayLanguage)
			t.refreshLanguage()
			t.refreshTooltip()
		case <-t.displayLanguage.Options[config.DisplayLanguageZh].ClickedCh:
			cfg.Update(config.CfgDisplayLanguage, config.DisplayLanguageZh)
			runtime.EventsEmit(t.ctx, "onDisplayLanguageChanged", cfg.DisplayLanguage)
			t.displayLanguage.Select(cfg.DisplayLanguage)
			t.refreshLanguage()
			t.refreshTooltip()
		case <-t.colorTheme.Options[config.ColorThemeSystem].ClickedCh:
			cfg.Update(config.CfgColorTheme, config.ColorThemeSystem)
			runtime.EventsEmit(t.ctx, "onColorThemeChanged", cfg.ColorTheme)
			runtime.WindowSetSystemDefaultTheme(t.ctx)
			t.colorTheme.Select(cfg.ColorTheme)
			t.refreshTooltip()
		case <-t.colorTheme.Options[config.ColorThemeLight].ClickedCh:
			cfg.Update(config.CfgColorTheme, config.ColorThemeLight)
			runtime.EventsEmit(t.ctx, "onColorThemeChanged", cfg.ColorTheme)
			runtime.WindowSetLightTheme(t.ctx)
			t.colorTheme.Select(cfg.ColorTheme)
			t.refreshTooltip()
		case <-t.colorTheme.Options[config.ColorThemeDark].ClickedCh:
			cfg.Update(config.CfgColorTheme, config.ColorThemeDark)
			runtime.EventsEmit(t.ctx, "onColorThemeChanged", cfg.ColorTheme)
			runtime.WindowSetDarkTheme(t.ctx)
			t.colorTheme.Select(cfg.ColorTheme)
			t.refreshTooltip()
		case <-t.quit.ClickedCh:
			ct := app.App().CurrentTranslation()
			dialog, err := runtime.MessageDialog(t.ctx, runtime.MessageDialogOptions{
				Type:    runtime.QuestionDialog,
				Title:   ct.AppName,
				Message: ct.QuitDialog.Message,
				Buttons: []string{
					ct.QuitDialog.DefaultButton,
					ct.QuitDialog.CancelButton,
				},
				DefaultButton: ct.QuitDialog.DefaultButton,
				CancelButton:  ct.QuitDialog.CancelButton,
				// Icon:          nil,
			})
			if err != nil {
				app.App().Log().Tray().Fatalf("fail to open quit dialog: %+v\n", err)
			}
			if dialog == "Yes" || dialog == ct.QuitDialog.DefaultButton {
				// when "Yes" or default button is clicked
				web.Web().Stop()
				systray.Quit()
				runtime.Quit(t.ctx)
				return
			}
		}
	}
}
