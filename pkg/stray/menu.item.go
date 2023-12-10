package stray

import (
	"context"
	"my-app/pkg/i18n"

	"github.com/getlantern/systray"
)

type MenuItem[TTranslation i18n.ITranslation] struct {
	bind *systray.MenuItem

	/* style */
	Seperator bool // append separator to systray root menu only
	CanCheck  bool // for linux to use checkbox menuitem

	/* state */
	Visible bool
	Enabled bool
	Checked bool

	OnClick func(ctx context.Context) (quit bool)
	SubMenu []*MenuItem[TTranslation]
}

func (menuItem *MenuItem[TTranslation]) TemplateIcon(translation TTranslation) (templateIconBytes []byte, regularIconBytes []byte) {
	panic("unimplemented")
}

func (menuItem *MenuItem[TTranslation]) Title(translation TTranslation) string {
	panic("unimplemented")
}

func (menuItem *MenuItem[TTranslation]) Tooltip(translation TTranslation) string {
	panic("unimplemented")
}

func NewMenuItem[TTranslation i18n.ITranslation]() *MenuItem[TTranslation] {
	return &MenuItem[TTranslation]{}
}
