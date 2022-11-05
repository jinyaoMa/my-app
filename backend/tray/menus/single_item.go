package menus

import "github.com/getlantern/systray"

type IRefresh interface {
	UpdateText() IRefresh
}

type SingleItem struct {
	IRefresh
	id          string
	item        *systray.MenuItem
	textUpdater func(updateText func(text string))
	isShowed    bool
	isEnabled   bool
	isChecked   bool
}

func NewSingleItem(id string, text string, icon ...[]byte) *SingleItem {
	si := &SingleItem{
		id:        id,
		item:      systray.AddMenuItem(text, text),
		isShowed:  true,
		isEnabled: true,
	}
	si.isChecked = si.item.Checked()
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
		if si.isShowed {
			si.Show()
		} else {
			si.Hide()
		}
		if si.isEnabled {
			si.Enable()
		} else {
			si.Disable()
		}
		if si.isChecked {
			si.Check()
		} else {
			si.Uncheck()
		}
	})
	return si
}

func (si *SingleItem) Show() *SingleItem {
	si.isShowed = true
	si.item.Show()
	return si
}

func (si *SingleItem) Hide() *SingleItem {
	si.isShowed = false
	si.item.Hide()
	return si
}

func (si *SingleItem) Enable() *SingleItem {
	si.isEnabled = true
	si.item.Enable()
	return si
}

func (si *SingleItem) Disable() *SingleItem {
	si.isEnabled = false
	si.item.Disable()
	return si
}

func (si *SingleItem) Check() *SingleItem {
	si.isChecked = true
	si.item.Check()
	return si
}

func (si *SingleItem) Uncheck() *SingleItem {
	si.isChecked = false
	si.item.Uncheck()
	return si
}

func (si *SingleItem) Clicked() chan struct{} {
	return si.item.ClickedCh
}
