package id

import (
	shortid "github.com/jasonsoft/go-short-id"
	"strings"
)

func GenShortId() string {
	opt := shortid.Options{
		Number:        6,
		StartWithYear: true,
		EndWithHost:   false,
	}
	return strings.ToLower(shortid.Generate(opt))
}
