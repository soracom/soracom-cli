package cmd

import (
	"regexp"
	"strconv"
)

// version strings are in the form of "v1.22.333" or "v0.0.1"
func isNewerThanCurrentVersion(latestVersion string) bool {
	cv := versionInt(version)
	lv := versionInt(latestVersion)
	return cv < lv
}

func versionInt(ver string) uint32 {
	s := splitVersionString(ver)
	if len(s) < 3 {
		return 0
	}

	var n uint32
	shift := uint(24)
	for i := 0; i < 4; i++ {
		if len(s) <= i {
			break
		}
		x, err := strconv.Atoi(s[i])
		if err == nil {
			n |= uint32((x & 0xff) << shift)
		}
		shift -= 8
	}
	return n
}

var versionStringRegexp = regexp.MustCompile("([[:digit:]]+)[[:^digit:]]*")

func splitVersionString(ver string) []string {
	m := versionStringRegexp.FindAllStringSubmatch(ver, -1)
	if len(m) < 2 {
		return []string{}
	}
	result := make([]string, len(m))
	for i, s := range m {
		result[i] = s[1]
	}
	return result
}
