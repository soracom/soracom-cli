package lib

import (
	"fmt"
	"os"
)

func PrintfStderr(format string, args ...interface{}) {
	_, err := fmt.Fprintf(os.Stderr, format, args...)
	if err != nil {
		//fmt.Printf("err: %+v\n", err) // this messes up stdout
	}
}

func WarnfStderr(format string, args ...interface{}) {
	PrintfStderr("WARN: "+format, args...)
}
