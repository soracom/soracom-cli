package cmd

import (
	"fmt"
	"strings"
)

// it was difficult to use fmt.Sprintf() in template codes
// because we need to import only if fmt.Sprintf() is called
// but the condition was too complicated. So we separated
// fmt.Sprintf() to a func which can be referenced from
// inside the complicated condition.

func sprintf(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

func strReplace(s, old, new string, n int) string {
	return strings.Replace(s, old, new, n)
}
