package cmd

import "os"

// API Endpoint Specification Precedence:
// 1. SORACOM_ENDPOINT env var
// 2. --coverage-type argument (use coverage type default)
// 3. endpoint in profile
// 4. coverageType in profile (use coverage type default)

func getSpecifiedEndpoint() string {
	e := os.Getenv("SORACOM_ENDPOINT")
	if e != "" {
		return e
	}

	if specifiedCoverageType != "" {
		return getDefaultEndpointForCoverageType(specifiedCoverageType)
	}

	profile, err := getProfile()
	if err != nil {
		return ""
	}

	if profile.Endpoint != nil {
		return *profile.Endpoint
	}

	return getDefaultEndpointForCoverageType(profile.CoverageType)
}

func getDefaultEndpointForCoverageType(ct string) string {
	if ct == "g" {
		return "https://g.api.soracom.io"
	}
	return "https://api.soracom.io"
}
