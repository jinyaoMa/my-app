package main

import (
	"context"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/quic-go/quic-go/http3"
	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/icons"
	"majinyao.cn/my-app/backend/cmd/wails/assets"
	"majinyao.cn/my-app/backend/internal/app"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/i18n"
)

const (
	SystrayWindowName = "Systray Window"
)

var (
	Systray                         *application.SystemTray
	SystrayTooltip                  string
	SystrayMenu                     *application.Menu
	SystrayMenuApi                  *application.MenuItem
	SystrayMenuApid                 *application.MenuItem
	SystrayMenuDocs                 *application.MenuItem
	SystrayMenuDisplayLanguageTitle *application.MenuItem
	SystrayMenuDisplayLanguageMap   map[string]*application.MenuItem
	SystrayMenuColorThemeTitle      *application.MenuItem
	SystrayMenuColorThemeSystem     *application.MenuItem
	SystrayMenuColorThemeLight      *application.MenuItem
	SystrayMenuColorThemeDark       *application.MenuItem
	SystrayMenuQuit                 *application.MenuItem
	SystrayWindow                   *application.WebviewWindow
)

func newSystray(a *application.App) *application.SystemTray {
	translation := app.I18N.GetTranslation()
	app.I18N.Watch(func(t i18n.Translation) (err error) {
		if SystrayMenuQuit == nil {
			return
		}
		if app.H3S.IsRunning() {
			SystrayMenuApi.SetLabel(t.Get("systray.menu.api.start"))
		} else {
			SystrayMenuApi.SetLabel(t.Get("systray.menu.api.stop"))
		}
		SystrayMenuApid.SetLabel(t.Get("systray.menu.apid"))
		SystrayMenuDocs.SetLabel(t.Get("systray.menu.docs"))
		SystrayMenuDisplayLanguageTitle.SetLabel(t.Get("systray.menu.display.language.title"))
		SystrayMenuColorThemeTitle.SetLabel(t.Get("systray.menu.color.theme.title"))
		SystrayMenuColorThemeSystem.SetLabel(t.Get("systray.menu.color.theme.system"))
		SystrayMenuColorThemeLight.SetLabel(t.Get("systray.menu.color.theme.light"))
		SystrayMenuColorThemeDark.SetLabel(t.Get("systray.menu.color.theme.dark"))
		SystrayMenuQuit.SetLabel(t.Get("systray.menu.quit"))
		for localeCode, menuitem := range SystrayMenuDisplayLanguageMap {
			menuitem.SetChecked(t.Code == localeCode)
		}
		updateSystrayTooltip(t)
		return
	})

	app.THEME.Watch(func(value string) (err error) {
		switch value {
		case entity.OptionColorThemeLight:
			SystrayMenuColorThemeSystem.SetChecked(false)
			SystrayMenuColorThemeLight.SetChecked(true)
			SystrayMenuColorThemeDark.SetChecked(false)
		case entity.OptionColorThemeDark:
			SystrayMenuColorThemeSystem.SetChecked(false)
			SystrayMenuColorThemeLight.SetChecked(false)
			SystrayMenuColorThemeDark.SetChecked(true)
		default:
			SystrayMenuColorThemeSystem.SetChecked(true)
			SystrayMenuColorThemeLight.SetChecked(false)
			SystrayMenuColorThemeDark.SetChecked(false)
		}
		updateSystrayTooltip(app.I18N.GetTranslation(), value)
		return
	})

	app.H3S.OnRun(func(h1, h2 *http.Server, h3 *http3.Server) {
		SystrayMenuApi.SetEnabled(true)

		SystrayMenuApid.SetHidden(false)
		SystrayMenuDocs.SetHidden(false)

		SystrayMenuApi.SetBitmap(assets.IconStart)
		SystrayMenuApid.SetBitmap(assets.IconApid)
		SystrayMenuDocs.SetBitmap(assets.IconDocs)

		t := app.I18N.GetTranslation()
		SystrayMenuApi.SetLabel(t.Get("systray.menu.api.start"))
		updateSystrayTooltip(t)
	})
	app.H3S.OnShutddown(func(h1, h2 *http.Server, h3 *http3.Server) {
		SystrayMenuApi.SetEnabled(true)

		SystrayMenuApid.SetHidden(true)
		SystrayMenuDocs.SetHidden(true)

		SystrayMenuApi.SetBitmap(assets.IconStop)

		t := app.I18N.GetTranslation()
		SystrayMenuApi.SetLabel(t.Get("systray.menu.api.stop"))
		updateSystrayTooltip(t)
	})

	SystrayMenu = a.Menu.New()

	if app.H3S.IsRunning() {
		SystrayMenuApi = SystrayMenu.
			Add(translation.Get("systray.menu.api.start")).
			SetBitmap(assets.IconStart)
	} else {
		SystrayMenuApi = SystrayMenu.
			Add(translation.Get("systray.menu.api.stop")).
			SetBitmap(assets.IconStop)
	}
	SystrayMenuApi.OnClick(func(ctx *application.Context) {
		if SystrayMenuApi.Enabled() {
			c, cancel := context.WithTimeout(a.Context(), 5*time.Second)
			defer cancel()
			SystrayMenuApi.SetEnabled(false)
			if app.H3S.IsRunning() {
				app.SHUTDOWN_H3S(c)
			} else {
				app.RUN_H3S(c, false)
			}
		}
	})

	SystrayMenuApid = SystrayMenu.
		Add(translation.Get("systray.menu.apid")).
		SetBitmap(assets.IconApid).
		SetHidden(!app.H3S.IsRunning()).
		OnClick(func(ctx *application.Context) {
			a.Browser.OpenURL("https://localhost" + app.H3S.SecureAddr() + app.API.GetDocsPath())
		})

	SystrayMenuDocs = SystrayMenu.
		Add(translation.Get("systray.menu.docs")).
		SetBitmap(assets.IconDocs).
		SetHidden(!app.H3S.IsRunning()).
		OnClick(func(ctx *application.Context) {
			a.Browser.OpenURL("https://localhost" + app.H3S.SecureAddr() + "/docs")
		})

	SystrayMenu.AddSeparator()

	SystrayMenuDisplayLanguageTitle = SystrayMenu.Add(translation.Get("systray.menu.display.language.title")).SetEnabled(false)
	SystrayMenuDisplayLanguageMap = make(map[string]*application.MenuItem)
	for _, locale := range app.I18N.AvailableLocales() {
		SystrayMenuDisplayLanguageMap[locale.Code] = SystrayMenu.
			AddCheckbox(locale.Text, translation.Code == locale.Code).
			OnClick(func(ctx *application.Context) {
				err := app.I18N.SetLocale(locale.Code, true)
				if err != nil {
					app.LOG.Println(err)
				}
			})
	}

	SystrayMenu.AddSeparator()

	theme := app.THEME.Get()
	SystrayMenuColorThemeTitle = SystrayMenu.Add(translation.Get("systray.menu.color.theme.title")).SetEnabled(false)
	SystrayMenuColorThemeSystem = SystrayMenu.
		AddCheckbox(translation.Get("systray.menu.color.theme.system"), theme == entity.OptionColorThemeSystem).
		OnClick(func(ctx *application.Context) {
			err := app.THEME.Set(entity.OptionColorThemeSystem, true)
			if err != nil {
				app.LOG.Println(err)
			}
		})
	SystrayMenuColorThemeLight = SystrayMenu.
		AddCheckbox(translation.Get("systray.menu.color.theme.light"), theme == entity.OptionColorThemeLight).
		OnClick(func(ctx *application.Context) {
			err := app.THEME.Set(entity.OptionColorThemeLight, true)
			if err != nil {
				app.LOG.Println(err)
			}
		})
	SystrayMenuColorThemeDark = SystrayMenu.
		AddCheckbox(translation.Get("systray.menu.color.theme.dark"), theme == entity.OptionColorThemeDark).
		OnClick(func(ctx *application.Context) {
			err := app.THEME.Set(entity.OptionColorThemeDark, true)
			if err != nil {
				app.LOG.Println(err)
			}
		})

	SystrayMenu.AddSeparator()

	SystrayMenu.Add("© 2025 jinyaoMa").SetEnabled(false)

	SystrayMenu.AddSeparator()

	SystrayMenuQuit = SystrayMenu.
		Add(translation.Get("systray.menu.quit")).
		OnClick(func(ctx *application.Context) {
			a.Quit()
		})

	SystrayWindow = a.Window.NewWithOptions(application.WebviewWindowOptions{
		Name:                     SystrayWindowName,
		Width:                    300,
		Height:                   375,
		Frameless:                true,
		AlwaysOnTop:              true,
		Hidden:                   true,
		DisableResize:            true,
		ContentProtectionEnabled: true,
		Windows: application.WindowsWindow{
			DisableIcon:     true,
			HiddenOnTaskbar: true,
		},
		KeyBindings: map[string]func(window application.Window){
			"F12": func(window application.Window) {

			},
		},
		URL: "/app/?hasReservedUsersOnStartup=" + strconv.FormatBool(app.HasReservedUsersOnStartup) + "#/systray",
	})

	Systray = a.SystemTray.New().
		AttachWindow(SystrayWindow).
		WindowOffset(5).
		SetIcon(assets.Icon).
		SetDarkModeIcon(assets.IconDark)
	if runtime.GOOS == "darwin" {
		Systray.SetTemplateIcon(icons.SystrayMacTemplate)
	}

	Systray.SetMenu(SystrayMenu).OnRightClick(func() {
		Systray.OpenMenu()
	})

	updateSystrayTooltip(translation, theme)
	return Systray
}

func updateSystrayTooltip(t i18n.Translation, theme ...string) {
	SystrayTooltip = t.Get("systray.tooltip.appname")
	if app.H3S.IsRunning() {
		SystrayTooltip += "\n" + t.Get("systray.tooltip.api.start")
	} else {
		SystrayTooltip += "\n" + t.Get("systray.tooltip.api.stop")
	}
	SystrayTooltip += "\n" + t.Get("systray.tooltip.language") + t.Text
	SystrayTooltip += "\n" + t.Get("systray.tooltip.theme")
	if len(theme) > 0 {
		SystrayTooltip += t.Get("systray.menu.color.theme." + strings.Join(theme, "."))
	} else {
		SystrayTooltip += t.Get("systray.menu.color.theme." + app.THEME.Get())
	}
	Systray.SetTooltip(SystrayTooltip)
}
