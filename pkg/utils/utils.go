package utils

import (
	"strings"
)

func URL(server string, uri string) string {
	if strings.HasSuffix(server, "/") {
		return server[:len(server)-1] + uri
	}
	return server + uri
}
