package utils

import (
	"regexp"
	"strings"
)

var cleanSpecialRegex = regexp.MustCompile(`[^a-zA-Z0-9\-]`)
var slugifyLDRegex = regexp.MustCompile(`([A-Z]|[0-9]+)`)

func SlugifyLDCamelCaseVariableName(v string, delims ...string) string {
	slug := "-"
	clean := cleanSpecialRegex.ReplaceAllString(v, "")
	slugified := slugifyLDRegex.ReplaceAllString(clean, slug+"${1}")
	lowered := strings.ToLower(slugified)
	trimed := strings.TrimPrefix(lowered, slug)
	for _, delim := range delims {
		trimed = strings.ReplaceAll(trimed, slug, delim)
		break // only use the first delimiter if there are multiple
	}
	return trimed
}

func SpaceLDCamelCaseVariableName(v string, delims ...string) string {
	space := " "
	clean := cleanSpecialRegex.ReplaceAllString(v, "")
	spaced := slugifyLDRegex.ReplaceAllString(clean, space+"${1}")
	trimed := strings.TrimPrefix(spaced, space)
	for _, delim := range delims {
		trimed = strings.ReplaceAll(trimed, space, delim)
		break // only use the first delimiter if there are multiple
	}
	return trimed
}
