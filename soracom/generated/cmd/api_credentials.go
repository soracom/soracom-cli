package cmd

import (
	"encoding/json"
	"strings"

	"github.com/dvsekhvalnov/jose2go/base64url"
)

type APICredentials struct {
	APIKey   string
	APIToken string
}

func apiCredentialsFromAuthResult(ares *authResult) *APICredentials {
	return &APICredentials{
		APIKey:   ares.APIKey,
		APIToken: ares.Token,
	}
}

func (c *APICredentials) getOperatorID() string {
	if c == nil {
		return ""
	}

	return extractOperatorIDFromAPIToken(c.APIToken)
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
