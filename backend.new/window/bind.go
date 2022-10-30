package window

import "my-app/backend.new/app"

type Bind struct{}

func NewBind() *Bind {
	return &Bind{}
}

// get AppName from i18n translation
func (b *Bind) GetAppName() string {
	return app.App().CurrentTranslation().AppName
}
