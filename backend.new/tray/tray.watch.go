package tray

import "my-app/backend.new/app"

func (t *tray) watch() {
	for {
		select {
		case <-t.openWindow.Clicked():
			app.App().Log().Tray().Println("open window single item clicked")
		case id := <-t.webService.OnGroupClicked():
			app.App().Log().Tray().Printf("web service switch group (on) clicked, id: %s\n", id)
		case id := <-t.webService.OffGroupClicked():
			app.App().Log().Tray().Printf("web service switch group (off) clicked, id: %s\n", id)
		case lang := <-t.displayLanguage.Selected():
			app.App().Log().Tray().Printf("display language select list selected, lang: %s\n", lang)
		case theme := <-t.colorTheme.Selected():
			app.App().Log().Tray().Printf("color theme select list selected, theme: %s\n", theme)
		case <-t.quit.Clicked():
			app.App().Log().Tray().Println("quit single item clicked")
		}
	}
}
