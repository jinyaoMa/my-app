package tray

import (
	"context"
	"fmt"
	"my-app/backend.new/app"
	"my-app/backend.new/app/i18n"
	"my-app/backend.new/app/types"
	"my-app/backend.new/model"
	"my-app/backend.new/web"

	"github.com/getlantern/systray"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// event names
const (
	EventNameOnWebServiceStateChanged = "OnWebServiceStateChanged"
	EventNameOnDisplayLanguageChanged = "OnDisplayLanguageChanged"
	EventNameOnColorThemeChanged      = "OnColorThemeChanged"
)

func (t *tray) ClickOpenWindow() {
	app.App().UseContext(func(ctx context.Context) {
		runtime.Show(ctx)
	}).Log().Tray().Println("OPEN WINDOW CLICKED")
}

func (t *tray) ClickOpenVitePress() {
	app.App().UseContextAndConfig(func(ctx context.Context, cfg *app.Config) {
		runtime.BrowserOpenURL(
			ctx,
			fmt.Sprintf("https://localhost%s/docs/", cfg.Get(model.OptionNameWebPortHttp)),
		)
	}).Log().Tray().Println("OPEN VITEPRESS CLICKED")
}

func (t *tray) ClickOpenSwagger() {
	app.App().UseContextAndConfig(func(ctx context.Context, cfg *app.Config) {
		runtime.BrowserOpenURL(
			ctx,
			fmt.Sprintf("https://localhost%s/swagger/index.html", cfg.Get(model.OptionNameWebPortHttps)),
		)
	}).Log().Tray().Println("OPEN SWAGGER CLICKED")
}

func (t *tray) ChangeWebServiceState(state bool) (ok bool) {
	if state {
		ok = web.Web().Start()
	} else {
		ok = web.Web().Stop()
	}
	if ok {
		app.App().UseContextAndConfig(func(ctx context.Context, cfg *app.Config) {
			runtime.EventsEmit(ctx, EventNameOnWebServiceStateChanged, state)
			t.webService.Switch(state)
			t.updateIconTooltip()
		}).Log().Tray().Printf("web service state changed to: %t\n", state)
	}
	return
}

func (t *tray) ChangeDisplayLanguage(lang string) {
	app.App().UseContextAndConfigAndI18n(func(ctx context.Context, cfg *app.Config, T func() *i18n.Translation, i18n *i18n.I18n) {
		if i18n.HasLanguage(lang) && cfg.Set(model.OptionNameDisplayLanguage, lang) {
			runtime.EventsEmit(ctx, EventNameOnDisplayLanguageChanged, lang)
			t.displayLanguage.Check(lang)
			t.updateLanguage().updateIconTooltip()
			app.App().Log().Tray().Printf("display language changed to: %s\n", lang)
		}
	})
}

func (t *tray) ChangeColorTheme(theme string) {
	theme = types.NewColorTheme(theme).ToString()
	app.App().UseContextAndConfigAndI18n(func(ctx context.Context, cfg *app.Config, T func() *i18n.Translation, i18n *i18n.I18n) {
		if cfg.Set(model.OptionNameColorTheme, theme) {
			runtime.EventsEmit(ctx, EventNameOnColorThemeChanged, theme)
			switch theme {
			default:
				runtime.WindowSetSystemDefaultTheme(ctx)
			case types.ColorThemeLight.ToString():
				runtime.WindowSetLightTheme(ctx)
			case types.ColorThemeDark.ToString():
				runtime.WindowSetDarkTheme(ctx)
			}
			t.colorTheme.Check(theme)
			t.updateIconTooltip()
			app.App().Log().Tray().Printf("color theme changed to: %s\n", theme)
		}
	})
}

func (t *tray) ClickQuit() {
	app.App().Log().Tray().Println("QUIT CLICKED")
	app.App().UseContextAndI18n(func(ctx context.Context, T func() *i18n.Translation, i18n *i18n.I18n) {
		dialog, err := runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
			Type:    runtime.QuestionDialog,
			Title:   T().AppName,
			Message: T().QuitDialog.Message,
			Buttons: []string{
				T().QuitDialog.DefaultButton,
				T().QuitDialog.CancelButton,
			},
			DefaultButton: T().QuitDialog.DefaultButton,
			CancelButton:  T().QuitDialog.CancelButton,
			// Icon:          nil,
		})
		if err != nil {
			app.App().Log().Tray().Fatalf("fail to open quit dialog: %+v\n", err)
		}
		if dialog == "Yes" || dialog == T().QuitDialog.DefaultButton {
			// when "Yes" or default button is clicked
			app.App().Log().Tray().Println("APPLICATION IS EXITING...")
			web.Web().Stop()
			systray.Quit()
			runtime.Quit(ctx)
		}
	})
}
