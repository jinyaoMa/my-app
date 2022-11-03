package tray

import (
	"fmt"
	"my-app/backend.new/app"
	"my-app/backend.new/app/types"
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
	runtime.Show(app.App().Ctx())
	app.App().Log().Tray().Println("OPEN WINDOW CLICKED")
}

func (t *tray) ClickOpenVitePress() {
	port := types.ParsePort(app.App().Cfg().Get(types.ConfigNameWebPortHttps))
	runtime.BrowserOpenURL(
		app.App().Ctx(),
		fmt.Sprintf("https://localhost:%d/docs/", port),
	)
	app.App().Log().Tray().Println("OPEN VITEPRESS CLICKED")
}

func (t *tray) ClickOpenSwagger() {
	port := types.ParsePort(app.App().Cfg().Get(types.ConfigNameWebPortHttps))
	runtime.BrowserOpenURL(
		app.App().Ctx(),
		fmt.Sprintf("https://localhost:%d/swagger/", port),
	)
	app.App().Log().Tray().Println("OPEN SWAGGER CLICKED")
}

func (t *tray) ChangeWebServiceState(state bool) (ok bool) {
	if state {
		ok = web.Web().Start()
	} else {
		ok = web.Web().Stop()
	}
	if ok {
		runtime.EventsEmit(app.App().Ctx(), EventNameOnWebServiceStateChanged, state)
		t.webService.Switch(state)
		t.updateIconTooltip()
		app.App().Log().Tray().Printf("web service state changed to: %t\n", state)
	}
	return
}

func (t *tray) ChangeDisplayLanguage(lang string) {
	if app.App().I18n().HasLanguage(lang) && app.App().Cfg().Set(types.ConfigNameDisplayLanguage, lang) {
		runtime.EventsEmit(app.App().Ctx(), EventNameOnDisplayLanguageChanged, lang)
		t.displayLanguage.Check(lang)
		t.updateLanguage().updateIconTooltip()
		app.App().Log().Tray().Printf("display language changed to: %s\n", lang)
	}
}

func (t *tray) ChangeColorTheme(theme string) {
	theme = types.ParseColorTheme(theme).ToString()
	if app.App().Cfg().Set(types.ConfigNameColorTheme, theme) {
		runtime.EventsEmit(app.App().Ctx(), EventNameOnColorThemeChanged, theme)
		switch theme {
		default:
			runtime.WindowSetSystemDefaultTheme(app.App().Ctx())
		case types.ColorThemeLight.ToString():
			runtime.WindowSetLightTheme(app.App().Ctx())
		case types.ColorThemeDark.ToString():
			runtime.WindowSetDarkTheme(app.App().Ctx())
		}
		t.colorTheme.Check(theme)
		t.updateIconTooltip()
		app.App().Log().Tray().Printf("color theme changed to: %s\n", theme)
	}
}

func (t *tray) ClickQuit() {
	app.App().Log().Tray().Println("QUIT CLICKED")
	T := app.App().T()
	dialog, err := runtime.MessageDialog(app.App().Ctx(), runtime.MessageDialogOptions{
		Type:    runtime.QuestionDialog,
		Title:   T.AppName,
		Message: T.QuitDialog.Message,
		Buttons: []string{
			T.QuitDialog.DefaultButton,
			T.QuitDialog.CancelButton,
		},
		DefaultButton: T.QuitDialog.DefaultButton,
		CancelButton:  T.QuitDialog.CancelButton,
		// Icon:          nil,
	})
	if err != nil {
		app.App().Log().Tray().Fatalf("fail to open quit dialog: %+v\n", err)
	}
	if dialog == "Yes" || dialog == T.QuitDialog.DefaultButton {
		// when "Yes" or default button is clicked
		app.App().Log().Tray().Println("APPLICATION IS EXITING...")
		web.Web().Stop()
		systray.Quit()
		runtime.Quit(app.App().Ctx())
	}
}
