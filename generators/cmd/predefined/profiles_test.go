package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProfileFromExternalCommand(t *testing.T) {
	email := "test@example.com"
	password := "testpassword"
	authKeyID := "testKeyID"
	authKey := "testAuthKey"
	username := "testUsername"
	operatorID := "testOperatorID"

	command := fmt.Sprintf("/bin/echo \"{\\\"email\\\": \\\"%s\\\",\\\"password\\\": \\\"%s\\\",\\\"authKeyId\\\": \\\"%s\\\",\\\"authKey\\\": \\\"%s\\\",\\\"userName\\\": \\\"%s\\\",\\\"operatorId\\\": \\\"%s\\\"}\"", email, password, authKeyID, authKey, username, operatorID)
	p, err := getProfileFromExternalCommand(command)

	if !assert.NoError(t, err) {
		t.FailNow()
	}
	assert.EqualValues(t, email, *(p.Email))
	assert.EqualValues(t, password, *(p.Password))
	assert.EqualValues(t, authKeyID, *(p.AuthKeyID))
	assert.EqualValues(t, authKey, *(p.AuthKey))
	assert.EqualValues(t, username, *(p.Username))
	assert.EqualValues(t, operatorID, *(p.OperatorID))
}

func TestGetAuthBodyFromProfile_AuthKey(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("SORACOM_PROFILE_DIR", tmpDir)

	profileJSON := `{"authKeyId":"key-123","authKey":"secret-456","coverageType":"jp"}`
	err := os.WriteFile(filepath.Join(tmpDir, "test-auth-key.json"), []byte(profileJSON), 0600)
	assert.NoError(t, err)

	body, err := getAuthBodyFromProfile("test-auth-key")
	assert.NoError(t, err)
	assert.Equal(t, "key-123", body["authKeyId"])
	assert.Equal(t, "secret-456", body["authKey"])
	assert.Equal(t, "jp", body["coverageType"])
}

func TestGetAuthBodyFromProfile_SAMUser(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("SORACOM_PROFILE_DIR", tmpDir)

	profileJSON := `{"operatorId":"OP123","username":"sam-user","password":"sam-password"}`
	err := os.WriteFile(filepath.Join(tmpDir, "test-sam.json"), []byte(profileJSON), 0600)
	assert.NoError(t, err)

	body, err := getAuthBodyFromProfile("test-sam")
	assert.NoError(t, err)
	assert.Equal(t, "OP123", body["operatorId"])
	assert.Equal(t, "sam-user", body["userName"])
	_, hasOldKey := body["username"]
	assert.False(t, hasOldKey)
	assert.Equal(t, "sam-password", body["password"])
}

func TestGetAuthBodyFromProfile_ProfileCommand(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("SORACOM_PROFILE_DIR", tmpDir)

	profileJSON := `{"profileCommand":"echo '{\"authKeyId\":\"cmd-key\",\"authKey\":\"cmd-secret\"}'"}`
	err := os.WriteFile(filepath.Join(tmpDir, "test-cmd.json"), []byte(profileJSON), 0600)
	assert.NoError(t, err)

	body, err := getAuthBodyFromProfile("test-cmd")
	assert.NoError(t, err)
	assert.Equal(t, "cmd-key", body["authKeyId"])
	assert.Equal(t, "cmd-secret", body["authKey"])
}

func TestGetAuthBodyFromProfile_NonExistent(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("SORACOM_PROFILE_DIR", tmpDir)

	_, err := getAuthBodyFromProfile("non-existent")
	assert.Error(t, err)
}

func TestGetAuthBodyFromProfile_SourceProfile(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("SORACOM_PROFILE_DIR", tmpDir)

	profileJSON := `{"operatorId":"OP123","userName":"sam-user","sourceProfile":"parent-profile"}`
	err := os.WriteFile(filepath.Join(tmpDir, "test-switch.json"), []byte(profileJSON), 0600)
	assert.NoError(t, err)

	_, err = getAuthBodyFromProfile("test-switch")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "sourceProfile")
}

func TestGetAuthBodyFromProfile_ProfileCommand_SourceProfile(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("SORACOM_PROFILE_DIR", tmpDir)

	profileJSON := `{"profileCommand":"echo '{\"operatorId\":\"OP123\",\"sourceProfile\":\"parent-profile\"}'"}`
	err := os.WriteFile(filepath.Join(tmpDir, "test-cmd-switch.json"), []byte(profileJSON), 0600)
	assert.NoError(t, err)

	_, err = getAuthBodyFromProfile("test-cmd-switch")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "sourceProfile")
}
