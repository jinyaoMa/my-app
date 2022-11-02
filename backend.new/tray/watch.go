package tray

import (
	"my-app/backend.new/app"
)

// menu ids
const (
	MenuIdOpenWindow      = "OpenWindow"
	MenuIdOpenVitePress   = "OpenVitePress"
	MenuIdOpenSwagger     = "OpenSwagger"
	MenuIdStopWeb         = "StopWeb"
	MenuIdStartWeb        = "StartWeb"
	MenuIdDisplayLanguage = "DisplayLanguage"
	MenuIdColorTheme      = "ColorTheme"
	MenuIdCopyright       = "Copyright"
	MenuIdQuit            = "Quit"
)

func (t *tray) watch() {
	app.App().Log().Tray().Println("TRAY IS ON WATCH")
	for {
		select {
		case <-t.openWindow.Clicked():
			t.ClickOpenWindow()
		case id := <-t.webService.OnGroupClicked():
			switch id {
			case MenuIdOpenVitePress:
				t.ClickOpenVitePress()
			case MenuIdOpenSwagger:
				t.ClickOpenSwagger()
			case MenuIdStopWeb:
				t.ChangeWebServiceState(false)
			}
		case id := <-t.webService.OffGroupClicked():
			switch id {
			case MenuIdStartWeb:
				t.ChangeWebServiceState(true)
			}
		case lang := <-t.displayLanguage.Selected():
			t.ChangeDisplayLanguage(lang)
		case theme := <-t.colorTheme.Selected():
			t.ChangeColorTheme(theme)
		case <-t.quit.Clicked():
			t.ClickQuit()
			return
		}
	}
}
