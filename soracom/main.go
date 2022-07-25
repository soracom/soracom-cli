package main

import (
	"os"

	"github.com/soracom/soracom-cli/soracom/generated/cmd"
)

func main() {
	os.Exit(run())
}

func run() int {
	cmd.InitRootCmd()
	err := cmd.RootCmd.Execute()
	if err != nil {
		return -1
	}
	return 0
}
