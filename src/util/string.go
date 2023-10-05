package util

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}

	return cases.Title(language.Und).String(s[0:1]) + s[1:]
}
