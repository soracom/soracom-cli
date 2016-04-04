package lib

import (
	"bytes"
	"regexp"
	"strings"
)

var wordRegex = regexp.MustCompile("[0-9A-Za-z]+")
var optionRegex = regexp.MustCompile("([a-z])([A-Z])")

// CamelCase converts strings like "space separated", "dash-separated", "snake_case" to "camelCase"
// but does not convert "TitleCase" to "camelCase"
func CamelCase(src string) string {
	byteSrc := bytes.NewBufferString(src).Bytes()
	chunks := wordRegex.FindAll(byteSrc, -1)
	for idx, val := range chunks {
		if idx > 0 {
			chunks[idx] = bytes.Title(val)
		}
	}
	return bytes.NewBuffer(bytes.Join(chunks, nil)).String()
}

// TitleCase converts strings like "space separated", "dash-separated", "snake_case", "camelCase" to "TitleCase"
func TitleCase(src string) string {
	byteSrc := bytes.NewBufferString(src).Bytes()
	chunks := wordRegex.FindAll(byteSrc, -1)
	for idx, val := range chunks {
		chunks[idx] = bytes.Title(val)
	}
	return bytes.NewBuffer(bytes.Join(chunks, nil)).String()
}

// SnakeCase converts strings like "space separated", "dash-separated", "camelCase", "TitleCase" to "snake_case"
func SnakeCase(s string) string {
	s1 := optionRegex.ReplaceAll([]byte(s), []byte("$1-$2"))
	s2 := strings.Replace(string(s1), " ", "_", -1)
	s3 := strings.Replace(s2, "-", "_", -1)
	s4 := bytes.ToLower([]byte(s3))
	return string(s4)
}

// OptionCase converts strings like "space separated", "snake_case", "camelCase", "TitleCase" to "option-case"
func OptionCase(s string) string {
	s1 := optionRegex.ReplaceAll([]byte(s), []byte("$1-$2"))
	s2 := strings.Replace(string(s1), " ", "-", -1)
	s3 := strings.Replace(s2, "_", "-", -1)
	s4 := bytes.ToLower([]byte(s3))
	return string(s4)
}
