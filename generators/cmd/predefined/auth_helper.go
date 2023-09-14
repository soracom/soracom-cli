package cmd

import (
	"encoding/json"
	"strings"

	"github.com/dvsekhvalnov/jose2go/base64url"
	"github.com/soracom/soracom-cli/generators/lib"
	"github.com/spf13/cobra"
)

const (
	minTokenTimeoutSeconds = 180
	maxTokenTimeoutSeconds = 3600
)

type authRequest struct {
	Email      *string `json:"email,omitempty"`
	Password   *string `json:"password,omitempty"`
	AuthKeyID  *string `json:"authKeyId,omitempty"`
	AuthKey    *string `json:"authKey,omitempty"`
	Username   *string `json:"userName,omitempty"`
	OperatorID *string `json:"operatorId,omitempty"`
	MfaOTPCode *string `json:"mfaOTPCode,omitempty"`
}

func authRequestFromProfile(p *profile) *authRequest {
	return &authRequest{
		Email:      p.Email,
		Password:   p.Password,
		AuthKeyID:  p.AuthKeyID,
		AuthKey:    p.AuthKey,
		Username:   p.Username,
		OperatorID: p.OperatorID,
		MfaOTPCode: p.MfaOTPCode,
	}
}

type authResult struct {
	APIKey     string `json:"apiKey"`
	Token      string `json:"token"`
	OperatorID string `json:"operatorId"`
}

type switchUserRequest struct {
	OperatorID          string `json:"operatorId"`
	UserName            string `json:"userName"`
	TokenTimeoutSeconds *int   `json:"tokenTimeoutSeconds,omitempty"`
}

func authHelper(ac *apiClient, cmd *cobra.Command, args []string) error {
	apiKey, apiToken, operatorID, credentialsProvided := getProvidedCredentials()
	if credentialsProvided {
		ac.APIKey = apiKey
		ac.Token = apiToken
		ac.OperatorID = operatorID
		return nil
	}

	var (
		ares *authResult
		err  error
	)

	if providedAuthKeyID != "" && providedAuthKey != "" {
		ares, err = ac.authenticateWithAuthKey(providedAuthKeyID, providedAuthKey)
		if err != nil {
			return err
		}
	} else {
		profile, err := getProfile()
		if err != nil {
			lib.PrintfStderr("unable to load the profile.\n")
			lib.PrintfStderr("run `soracom configure` first.\n")
			return err
		}

		ares, err = ac.authenticateWithProfile(profile)
		if err != nil {
			return err
		}
	}

	ac.APIKey = ares.APIKey
	ac.Token = ares.Token
	ac.OperatorID = ares.OperatorID

	return nil
}

func getProvidedCredentials() (string, string, string, bool) {
	operatorID := extractOperatorIDFromAPIToken(providedAPIToken)
	return providedAPIKey, providedAPIToken, operatorID, (providedAPIKey != "" && providedAPIToken != "")
}

func getProvidedTokenTimeoutSeconds(profile *profile) *int {
	// TODO: support providing tokenTimeoutSeconds from command line option
	//if providedTokenTimeoutSeconds != 0 {
	//return providedTokenTimeoutSeconds
	//}
	if profile.TokenTimeoutSeconds != nil && isValidTokenTimeoutSeconds(*profile.TokenTimeoutSeconds) {
		return profile.TokenTimeoutSeconds
	}
	return nil
}

func isValidTokenTimeoutSeconds(tokenTimeoutSeconds int) bool {
	if tokenTimeoutSeconds < minTokenTimeoutSeconds {
		return false
	}
	if tokenTimeoutSeconds > maxTokenTimeoutSeconds {
		return false
	}

	return true
}

type jwtPayload struct {
	Operator jwtPayloadOperator `json:"operator"`
}

type jwtPayloadOperator struct {
	OperatorID string `json:"operatorId"`
}

func extractOperatorIDFromAPIToken(apiToken string) string {
	parts := strings.Split(apiToken, ".")
	if len(parts) < 2 {
		return ""
	}

	b64Decoded, err := base64url.Decode(parts[1])
	if err != nil {
		return ""
	}

	var jp jwtPayload
	err = json.Unmarshal(b64Decoded, &jp)
	if err != nil {
		return ""
	}

	return jp.Operator.OperatorID
}

func toJSON(x interface{}) string {
	bodyBytes, err := json.Marshal(x)
	if err != nil {
		return ""
	}
	return string(bodyBytes)
}
