package app

import "my-app/backend/pkg/assetsio"

type Translation struct {
	lang *assetsio.Lang
}

// Lang implements assetsio.ITranslation.
func (t *Translation) Lang() *assetsio.Lang {
	return t.lang
}

func DefaultTranslation() *Translation {
	return &Translation{
		lang: &assetsio.Lang{
			Code: "",
			Text: "",
		},
	}
}

func NewTranslation() assetsio.ITranslation {
	return &Translation{}
}
