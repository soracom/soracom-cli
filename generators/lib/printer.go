package lib

import (
	"fmt"
	"os"
)

// PrintfStderr formats according to a format specifier and writes to standard error
func PrintfStderr(format string, args ...interface{}) {
	_, err := fmt.Fprintf(os.Stderr, format, args...)
	if err != nil {
		//fmt.Printf("err: %+v\n", err) // this messes up stdout
	}
}

// WarnfStderr formats according to a format specifier and writes to standard error with `WARN: ` prefix
func WarnfStderr(format string, args ...interface{}) {
	PrintfStderr("WARN: "+format, args...)
}
