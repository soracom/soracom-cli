package main

import (
	"os"

	"github.com/soracom/soracom-cli/soracom/generated/cmd"
)

func main() {
	os.Exit(run())
}

func run() int {
	err := cmd.RootCmd.Execute()
	if err != nil {
		return -1
	}
	return 0
}
