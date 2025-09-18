package i18n

import (
	"fmt"

	"github.com/tidwall/gjson"
)

type Translation struct {
	Locale
	result gjson.Result
}

func (t *Translation) Get(path string, args ...any) string {
	format := t.result.Get(path).String()
	if format == "" {
		return fmt.Sprintf("[%s]", path)
	}
	return fmt.Sprintf(format, args...)
}
