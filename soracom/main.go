package main

// generate i18n data to be embedded to cli
//go:generate go-assets-builder --strip-prefix="/generators/assets" --output="./generated/cmd/i18n_data.go" --package="cmd" ../generators/assets

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
