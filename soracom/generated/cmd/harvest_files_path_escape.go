package cmd

import (
	"net/url"
	"strings"
)

func harvestFilesPathEscape(path string) string {
	result := []string{}
	for _, s := range strings.Split(path, "/") {
		if s == "" {
			continue
		}
		result = append(result, url.PathEscape(s))
	}
	return strings.Join(result, "/")
}
