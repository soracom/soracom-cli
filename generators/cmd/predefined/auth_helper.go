package cmd

import (
	"bytes"
	"encoding/json"

	"github.com/spf13/cobra"
)

type authRequest struct {
	Email      *string `json:"email,omitempty"`
	Password   *string `json:"password,omitempty"`
	AuthKeyID  *string `json:"authKeyId,omitempty"`
	AuthKey    *string `json:"authKey,omitempty"`
	Username   *string `json:"userName,omitempty"`
	OperatorID *string `json:"operatorId,omitempty"`
}

type authResult struct {
	APIKey     string `json:"apiKey"`
	Token      string `json:"token"`
	OperatorID string `json:"operatorId"`
}

func authHelper(ac *apiClient, cmd *cobra.Command, args []string) error {
	apiKey, apiToken, credentialsProvided := getProvidedCredentials()
	if credentialsProvided {
		ac.APIKey = apiKey
		ac.Token = apiToken
		return nil
	}

	profile, err := getProfile()
	if err != nil {
		printfStderr("unable to load the profile.\n")
		printfStderr("run `soracom configure` first.\n")
		return err
	}

	areq := &authRequest{
		Email:      profile.Email,
		Password:   profile.Password,
		AuthKeyID:  profile.AuthKeyID,
		AuthKey:    profile.AuthKey,
		Username:   profile.Username,
		OperatorID: profile.OperatorID,
	}

	params := &apiParams{
		method:         "POST",
		path:           "/auth",
		query:          map[string][]string{},
		contentType:    "application/json",
		body:           toJSON(areq),
		noVersionCheck: true,
	}

	res, err := ac.callAPI(params)
	if err != nil {
		return err
	}

	dec := json.NewDecoder(bytes.NewBufferString(res))
	var ares authResult
	err = dec.Decode(&ares)
	if err != nil {
		return err
	}

	ac.APIKey = ares.APIKey
	ac.Token = ares.Token
	ac.OperatorID = ares.OperatorID
	return nil
}

func getProvidedCredentials() (string, string, bool) {
	return providedAPIKey, providedAPIToken, (providedAPIKey != "" && providedAPIToken != "")
}

func toJSON(x interface{}) string {
	bodyBytes, err := json.Marshal(x)
	if err != nil {
		return ""
	}
	return string(bodyBytes)
}
