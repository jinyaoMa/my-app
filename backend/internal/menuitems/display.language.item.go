package menuitems

import (
	"context"
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
func (*displayLanguageItem) Checked() bool {
	return false
}

// Key implements tray.IMenuItem.
func (i *displayLanguageItem) Key() string {
	return "display.language." + i.lang.Code
}

// OnClick implements tray.IMenuItem.
func (t *displayLanguageItem) OnClick() (quit bool) {
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
		MenuItem: &tray.MenuItem{
			MenuItemBase: &tray.MenuItemBase{
				Ctx: ctx,
			},
		},
		lang: lang,
	}
}
