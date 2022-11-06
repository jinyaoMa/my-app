package menus

import "github.com/getlantern/systray"

type TextUpdater func(id string) (updateText string)

type SingleItem struct {
	id          string
	item        *systray.MenuItem
	textUpdater TextUpdater
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

func (si *SingleItem) SetTextUpdater(textUpdater TextUpdater) *SingleItem {
	si.textUpdater = textUpdater
	return si
}

func (si *SingleItem) UpdateText() *SingleItem {
	text := si.textUpdater(si.id)
	si.item.SetTitle(text)
	si.item.SetTooltip(text)
	if si.isShowed {
		si.item.Show()
	} else {
		si.item.Hide()
	}
	if si.isEnabled {
		si.item.Enable()
	} else {
		si.item.Disable()
	}
	if si.isChecked {
		si.item.Check()
	} else {
		si.item.Uncheck()
	}
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
