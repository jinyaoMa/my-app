package menus

import "github.com/getlantern/systray"

type SingleItem struct {
	id          string
	item        *systray.MenuItem
	textUpdater func(updateText func(text string))
}

func NewSingleItem(id string, text string, icon ...[]byte) *SingleItem {
	si := &SingleItem{
		id:   id,
		item: systray.AddMenuItem(text, text),
	}
	if len(icon) > 0 {
		si.item.SetTemplateIcon(icon[0], icon[0])
	}
	return si
}

func (si *SingleItem) GetID() string {
	return si.id
}

func (si *SingleItem) SetTextUpdater(textUpdater func(updateText func(text string))) *SingleItem {
	si.textUpdater = textUpdater
	return si
}

func (si *SingleItem) UpdateText() *SingleItem {
	si.textUpdater(func(text string) {
		si.item.SetTitle(text)
		si.item.SetTooltip(text)
	})
	return si
}

func (si *SingleItem) Show() *SingleItem {
	si.item.Show()
	return si
}

func (si *SingleItem) Hide() *SingleItem {
	si.item.Hide()
	return si
}

func (si *SingleItem) Enable() *SingleItem {
	si.item.Enable()
	return si
}

func (si *SingleItem) Disable() *SingleItem {
	si.item.Disable()
	return si
}

func (si *SingleItem) Check() *SingleItem {
	si.item.Check()
	return si
}

func (si *SingleItem) Uncheck() *SingleItem {
	si.item.Uncheck()
	return si
}

func (si *SingleItem) Clicked() chan struct{} {
	return si.item.ClickedCh
}
