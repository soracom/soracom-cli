package cmd

import (
	"net/url"
	"testing"
)

func newTestAPIClient() *apiClient {
	return &apiClient{
		apiCredentials: &APICredentials{
			APIKey:   "super-secret-key",
			APIToken: "super-secret-token",
		},
		endpoint: "https://g.api.soracom.io",
		basePath: "/v1",
		language: "en",
	}
}

func TestDryRunOutputRedactsSecretHeaders(t *testing.T) {
	ac := newTestAPIClient()
	params := &apiParams{
		method:      "POST",
		path:        "/groups",
		query:       url.Values{},
		contentType: "application/json",
		body:        `{"tags":{"name":"sensor-01"}}`,
	}

	out, err := ac.dryRunOutput(params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if out["method"] != "POST" {
		t.Errorf("method = %v, want POST", out["method"])
	}
	if out["url"] != "https://g.api.soracom.io/v1/groups" {
		t.Errorf("url = %v (trailing '?' should be trimmed)", out["url"])
	}

	headers, ok := out["headers"].(map[string]interface{})
	if !ok {
		t.Fatalf("headers missing or wrong type: %#v", out["headers"])
	}
	if got := headers["X-Soracom-Api-Key"]; got != "<hidden>" {
		t.Errorf("API key should be redacted, got %v", got)
	}
	if got := headers["X-Soracom-Token"]; got != "<hidden>" {
		t.Errorf("token should be redacted, got %v", got)
	}

	body, ok := out["body"].(map[string]interface{})
	if !ok {
		t.Fatalf("body should be parsed JSON, got %#v", out["body"])
	}
	tags, ok := body["tags"].(map[string]interface{})
	if !ok || tags["name"] != "sensor-01" {
		t.Errorf("body not parsed as expected: %#v", body)
	}
}

func TestDryRunOutputRedactsProfileHeaders(t *testing.T) {
	ac := newTestAPIClient()
	ac.profile = &profile{
		Headers: map[string]string{
			"Authorization":  "Bearer super-secret",
			"X-Custom-Token": "opaque-secret",
		},
	}
	params := &apiParams{method: "GET", path: "/operators/x", query: url.Values{}}

	out, err := ac.dryRunOutput(params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	headers, ok := out["headers"].(map[string]interface{})
	if !ok {
		t.Fatalf("headers missing: %#v", out["headers"])
	}
	for _, h := range []string{"Authorization", "X-Custom-Token"} {
		if headers[h] != "<hidden>" {
			t.Errorf("profile header %q should be redacted, got %v", h, headers[h])
		}
	}
}

func TestIsSecretHeaderName(t *testing.T) {
	secret := []string{"X-Soracom-Api-Key", "x-soracom-api-key", "X-Soracom-Token", "X-SORACOM-TOKEN"}
	for _, h := range secret {
		if !isSecretHeaderName(h) {
			t.Errorf("%q should be treated as secret", h)
		}
	}
	notSecret := []string{"Content-Type", "User-Agent", "X-Soracom-Lang"}
	for _, h := range notSecret {
		if isSecretHeaderName(h) {
			t.Errorf("%q should NOT be treated as secret", h)
		}
	}
}
