package stray

import (
	"context"
	"my-app/pkg/i18n"

	"github.com/getlantern/systray"
)

type MenuItem[TTranslation i18n.ITranslation] struct {
	Key  string
	bind *systray.MenuItem

	/* style */
	Seperator bool // append separator to systray root menu only
	CanCheck  bool // for linux to use checkbox menuitem

	/* state */
	Visible bool
	Enabled bool
	Checked bool

	TemplateIcon func(translation TTranslation) (templateIconBytes []byte, regularIconBytes []byte)
	Title        func(translation TTranslation) string
	Tooltip      func(translation TTranslation) string
	OnClick      func(menuitem *MenuItem[TTranslation], ctx context.Context) (quit bool)
	SubMenu      []*MenuItem[TTranslation]
}
