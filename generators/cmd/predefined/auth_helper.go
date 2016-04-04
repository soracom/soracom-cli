package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/spf13/cobra"
)

type soracomAuthInfoWriter struct {
	apiKey   string
	apiToken string
}

func (w *soracomAuthInfoWriter) AuthenticateRequest(req client.Request, _ strfmt.Registry) error {
	err := req.SetHeaderParam("X-Soracom-API-Key", w.apiKey)
	if err != nil {
		return err
	}
	err = req.SetHeaderParam("X-Soracom-Token", w.apiToken)
	if err != nil {
		return err
	}
	return nil
}

type authRequest struct {
	Email      *string `json:"email,omitempty"`
	Password   *string `json:"password,omitempty"`
	AuthKeyID  *string `json:"authKeyId,omitempty"`
	AuthKey    *string `json:"authKey,omitempty"`
	OperatorID *string `json:"operatorId,omitempty"`
}

type authResult struct {
	APIKey string `json:"apiKey"`
	Token  string `json:"token"`
}

func authHelper(ac *apiClient, cmd *cobra.Command, args []string) error {
	pn := getSpecifiedProfileName()
	if pn == "" {
		pn = "default"
	}

	profile, err := loadProfile(pn)
	if err != nil {
		fmt.Println("unable to load the specified profile: " + pn)
		fmt.Println("run `soracom configure` first.")
		return err
	}

	areq := &authRequest{
		Email:      profile.Email,
		Password:   profile.Password,
		AuthKeyID:  profile.AuthKeyID,
		AuthKey:    profile.AuthKey,
		OperatorID: profile.OperatorID,
	}

	params := &apiParams{
		method:      "POST",
		path:        "/auth",
		query:       "",
		contentType: "application/json",
		body:        toJSON(areq),
	}

	res, err := ac.callAPI(params)
	if err != nil {
		fmt.Println("auth failed: ", err.Error())
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
	return nil
}

func toJSON(x interface{}) string {
	bodyBytes, err := json.Marshal(x)
	if err != nil {
		return ""
	}
	return string(bodyBytes)
}
