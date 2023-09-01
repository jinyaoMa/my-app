package menuitems

import (
	"context"
	"my-app/backend/internal/app"
	"my-app/backend/pkg/aio"
	"my-app/backend/pkg/tray"
)

type displayLanguageItem struct {
	*tray.MenuItem
	lang aio.Lang
}

// CanCheck implements tray.IMenuItem.
func (*displayLanguageItem) CanCheck() bool {
	return true
}

// Checked implements tray.IMenuItem.
func (i *displayLanguageItem) Checked() bool {
	return app.LANG() == i.lang.Code
}

// Key implements tray.IMenuItem.
func (i *displayLanguageItem) Key() string {
	return "display.language." + i.lang.Code
}

// CanClick implements tray.IMenuItem.
func (*displayLanguageItem) CanClick() bool {
	return true
}

// OnClick implements tray.IMenuItem.
func (i *displayLanguageItem) OnClick() (quit bool) {
	app.LANG(i.lang.Code)
	tray.Update(_root, false)
	return false
}

// Title implements tray.IMenuItem.
func (i *displayLanguageItem) Title() string {
	return i.lang.Text
}

// Tooltip implements tray.IMenuItem.
func (i *displayLanguageItem) Tooltip() string {
	return i.lang.Text
}

func newdisplayLanguageItem(ctx context.Context, lang aio.Lang) tray.IMenuItem {
	return &displayLanguageItem{
		MenuItem: tray.NewMenuItem(ctx),
		lang:     lang,
	}
}
