package cmd

import (
	"fmt"
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
