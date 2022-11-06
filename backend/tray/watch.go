package tray

import (
	"my-app/backend/app"
)

// watch listen to menu events
func (t *tray) watch() *tray {
	go func() {
		for {
			select {
			case <-t.openWindow.Clicked():
				t.clickOpenWindow()
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
				if t.clickQuit() {
					return
				}
			}
		}
	}()
	app.App().Log().Tray().Println("TRAY IS ON WATCH")
	return t
}
