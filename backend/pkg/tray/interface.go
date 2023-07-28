package tray

type Interface interface {
	Key() string

	Separator() bool

	Icon() []byte

	Title() string

	Tooltip() string

	Visible() bool

	Enabled() bool

	Checked() bool

	OnClick()

	Items() []Interface
}
