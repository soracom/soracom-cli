package cmd

import "os"

var (
	defaultEndpoint = ""
)

func setDefaultEndpoint(endpoint string) {
	defaultEndpoint = endpoint
}

func getSpecifiedEndpoint() string {
	e := os.Getenv("SORACOM_ENDPOINT")
	if e == "" {
		e = defaultEndpoint
	}
	return e
}
