package main

// generate i18n data to be embedded to cli
//go:generate go-bindata -o ./generated/cmd/i18n_data.go -pkg cmd ../generators/assets/ ../generators/assets/cli/

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
