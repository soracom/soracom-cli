package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// resetEndpointTestState isolates the package-level state that endpoint
// resolution reads, and restores it when the test finishes.
func resetEndpointTestState(t *testing.T) {
	t.Helper()
	t.Setenv("SORACOM_ENDPOINT", "")
	t.Setenv("SORACOM_PROFILE_DIR", t.TempDir())

	origProfileName, origCoverageType, origProfileCommand := specifiedProfileName, specifiedCoverageType, providedProfileCommand
	t.Cleanup(func() {
		specifiedProfileName, specifiedCoverageType, providedProfileCommand = origProfileName, origCoverageType, origProfileCommand
		loadedProfile = nil
		externalCommandProfiles = map[string]externalCommandProfileResult{}
	})
	specifiedProfileName, specifiedCoverageType, providedProfileCommand = "", "", ""
	loadedProfile = nil
	externalCommandProfiles = map[string]externalCommandProfileResult{}
}

func writeTestProfile(t *testing.T, name, content string) {
	t.Helper()
	path := filepath.Join(os.Getenv("SORACOM_PROFILE_DIR"), name+".json")
	assert.NoError(t, os.WriteFile(path, []byte(content), 0600))
}

// writeProfileScript writes a script that prints the given profile JSON and
// appends a line to a counter file each time it runs, so tests can check how
// many times the command was executed.
func writeProfileScript(t *testing.T, profileJSON string) (command string, counterPath string) {
	t.Helper()
	dir := t.TempDir()
	counterPath = filepath.Join(dir, "count")
	scriptPath := filepath.Join(dir, "profile.sh")
	script := fmt.Sprintf("#!/bin/sh\necho x >> '%s'\ncat <<'EOF'\n%s\nEOF\n", counterPath, profileJSON)
	assert.NoError(t, os.WriteFile(scriptPath, []byte(script), 0700))
	return scriptPath, counterPath
}

func countExecutions(t *testing.T, counterPath string) int {
	t.Helper()
	b, err := os.ReadFile(counterPath)
	if os.IsNotExist(err) {
		return 0
	}
	assert.NoError(t, err)
	return strings.Count(string(b), "\n")
}

func TestGetSpecifiedEndpoint_ProfileCommandCoverageType(t *testing.T) {
	resetEndpointTestState(t)
	// The default profile points to Japan coverage; the profile command must win.
	writeTestProfile(t, "default", `{"coverageType":"jp"}`)
	command, _ := writeProfileScript(t, `{"coverageType":"g","authKeyId":"keyId-x","authKey":"secret-x"}`)
	providedProfileCommand = command

	assert.Equal(t, "https://g.api.soracom.io", getSpecifiedEndpoint())
}

func TestGetSpecifiedEndpoint_ProfileCommandSandbox(t *testing.T) {
	resetEndpointTestState(t)
	command, _ := writeProfileScript(t, `{"coverageType":"g","sandbox":true}`)
	providedProfileCommand = command

	assert.Equal(t, "https://g.api-sandbox.soracom.io", getSpecifiedEndpoint())
}

func TestGetSpecifiedEndpoint_ProfileCommandEndpoint(t *testing.T) {
	resetEndpointTestState(t)
	command, _ := writeProfileScript(t, `{"coverageType":"jp","endpoint":"https://example.com"}`)
	providedProfileCommand = command

	assert.Equal(t, "https://example.com", getSpecifiedEndpoint())
}

func TestGetSpecifiedEndpoint_ProfileFileWithProfileCommand(t *testing.T) {
	resetEndpointTestState(t)
	command, _ := writeProfileScript(t, `{"coverageType":"g","authKeyId":"keyId-x","authKey":"secret-x"}`)
	writeTestProfile(t, "agent", fmt.Sprintf(`{"profileCommand":%q}`, command))
	specifiedProfileName = "agent"

	assert.Equal(t, "https://g.api.soracom.io", getSpecifiedEndpoint())
}

func TestGetSpecifiedEndpoint_CoverageTypeFlagOverridesProfileCommand(t *testing.T) {
	resetEndpointTestState(t)
	command, _ := writeProfileScript(t, `{"coverageType":"g"}`)
	providedProfileCommand = command
	specifiedCoverageType = "jp"

	assert.Equal(t, "https://api.soracom.io", getSpecifiedEndpoint())
}

func TestGetSpecifiedEndpoint_ProfileCommandFailureFallsBackToDefault(t *testing.T) {
	resetEndpointTestState(t)
	providedProfileCommand = "/nonexistent/profile-command"

	assert.Equal(t, "https://api.soracom.io", getSpecifiedEndpoint())
}

func TestLoadProfileFromExternalCommand_RunsCommandOnce(t *testing.T) {
	resetEndpointTestState(t)
	command, counterPath := writeProfileScript(t, `{"coverageType":"g","authKeyId":"keyId-x","authKey":"secret-x"}`)
	providedProfileCommand = command

	getSpecifiedEndpoint()
	p, err := loadProfileFromExternalCommand(command)

	assert.NoError(t, err)
	assert.Equal(t, "keyId-x", *p.AuthKeyID)
	assert.Equal(t, 1, countExecutions(t, counterPath))
}

func TestLoadProfileFromExternalCommand_EmptyCommand(t *testing.T) {
	resetEndpointTestState(t)

	p, err := loadProfileFromExternalCommand("   ")

	assert.Nil(t, p)
	assert.Error(t, err)
}
