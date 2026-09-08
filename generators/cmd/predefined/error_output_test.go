package cmd

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"
)

func TestPrintErrorWrapsPlainError(t *testing.T) {
	var sb strings.Builder
	printErrorTo(&sb, errors.New("required parameter 'imsi' is not specified"))

	got := strings.TrimSpace(sb.String())
	var parsed struct {
		Error struct {
			Message string `json:"message"`
		} `json:"error"`
	}
	if err := json.Unmarshal([]byte(got), &parsed); err != nil {
		t.Fatalf("output is not valid JSON: %q (%v)", got, err)
	}
	if parsed.Error.Message != "required parameter 'imsi' is not specified" {
		t.Errorf("unexpected message: %q", parsed.Error.Message)
	}
}

func TestPrintErrorPassesThroughAPIErrorJSON(t *testing.T) {
	apiBody := `{"code":"SEM0001","message":"resource not found"}`
	var sb strings.Builder
	printErrorTo(&sb, newAPIError(apiBody))

	got := strings.TrimSpace(sb.String())
	if got != apiBody {
		t.Errorf("API error body should be emitted as-is.\n got: %q\nwant: %q", got, apiBody)
	}
}

func TestPrintErrorWrapsNonJSONAPIError(t *testing.T) {
	var sb strings.Builder
	printErrorTo(&sb, newAPIError("plain text failure"))

	got := strings.TrimSpace(sb.String())
	if !json.Valid([]byte(got)) {
		t.Errorf("non-JSON API error should be wrapped into valid JSON, got: %q", got)
	}
}

func TestPrintErrorNil(t *testing.T) {
	var sb strings.Builder
	printErrorTo(&sb, nil)
	if sb.String() != "" {
		t.Errorf("nil error should produce no output, got: %q", sb.String())
	}
}

func TestJSONErrorsEnabled(t *testing.T) {
	t.Setenv("SORACOM_JSON_ERRORS", "")
	if jsonErrorsEnabled() {
		t.Error("should be disabled by default (empty env var)")
	}
	t.Setenv("SORACOM_JSON_ERRORS", "1")
	if !jsonErrorsEnabled() {
		t.Error("should be enabled when SORACOM_JSON_ERRORS is set")
	}
}
