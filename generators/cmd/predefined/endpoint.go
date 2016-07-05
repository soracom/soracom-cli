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
	if e != "" {
		return e
	}

	profile, err := getProfile()
	if err != nil {
		return defaultEndpoint
	}

	if profile.Endpoint == nil {
		return defaultEndpoint
	}

	return *profile.Endpoint
}
