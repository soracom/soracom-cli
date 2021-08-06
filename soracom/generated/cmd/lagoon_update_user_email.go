// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"

	"fmt"

	"io/ioutil"

	"net/url"
	"os"

	"strings"

	"github.com/spf13/cobra"
)

// LagoonUpdateUserEmailCmdUserEmail holds value of 'userEmail' option
var LagoonUpdateUserEmailCmdUserEmail string

// LagoonUpdateUserEmailCmdLagoonUserId holds value of 'lagoon_user_id' option
var LagoonUpdateUserEmailCmdLagoonUserId int64

// LagoonUpdateUserEmailCmdBody holds contents of request body to be sent
var LagoonUpdateUserEmailCmdBody string

func init() {
	LagoonUpdateUserEmailCmd.Flags().StringVar(&LagoonUpdateUserEmailCmdUserEmail, "user-email", "", TRAPI(""))

	LagoonUpdateUserEmailCmd.Flags().Int64Var(&LagoonUpdateUserEmailCmdLagoonUserId, "lagoon-user-id", 0, TRAPI("Target ID of the lagoon user"))

	LagoonUpdateUserEmailCmd.Flags().StringVar(&LagoonUpdateUserEmailCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	LagoonCmd.AddCommand(LagoonUpdateUserEmailCmd)
}

// LagoonUpdateUserEmailCmd defines 'update-user-email' subcommand
var LagoonUpdateUserEmailCmd = &cobra.Command{
	Use:   "update-user-email",
	Short: TRAPI("/lagoon/users/{lagoon_user_id}/email:put:summary"),
	Long:  TRAPI(`/lagoon/users/{lagoon_user_id}/email:put:description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			BasePath: "/v1",
			Language: getSelectedLanguage(),
		}

		ac := newAPIClient(opt)
		if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
			ac.SetVerbose(true)
		}
		err := authHelper(ac, cmd, args)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		param, err := collectLagoonUpdateUserEmailCmdParams(ac)
		if err != nil {
			return err
		}

		body, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		if jqString != "" {
			return processJQ(jqString, body)
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectLagoonUpdateUserEmailCmdParams(ac *apiClient) (*apiParams, error) {
	body, err := buildBodyForLagoonUpdateUserEmailCmd()
	if err != nil {
		return nil, err
	}
	contentType := "application/json"

	if LagoonUpdateUserEmailCmdLagoonUserId == 0 {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "lagoon-user-id")
		}

	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForLagoonUpdateUserEmailCmd("/lagoon/users/{lagoon_user_id}/email"),
		query:       buildQueryForLagoonUpdateUserEmailCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLagoonUpdateUserEmailCmd(path string) string {

	path = strReplace(path, "{"+"lagoon_user_id"+"}", url.PathEscape(sprintf("%d", LagoonUpdateUserEmailCmdLagoonUserId)), -1)

	return path
}

func buildQueryForLagoonUpdateUserEmailCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForLagoonUpdateUserEmailCmd() (string, error) {
	var result map[string]interface{}

	if LagoonUpdateUserEmailCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(LagoonUpdateUserEmailCmdBody, "@") {
			fname := strings.TrimPrefix(LagoonUpdateUserEmailCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if LagoonUpdateUserEmailCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(LagoonUpdateUserEmailCmdBody)
		}

		if err != nil {
			return "", err
		}

		err = json.Unmarshal(b, &result)
		if err != nil {
			return "", err
		}
	}

	if result == nil {
		result = make(map[string]interface{})
	}

	if LagoonUpdateUserEmailCmdUserEmail != "" {
		result["userEmail"] = LagoonUpdateUserEmailCmdUserEmail
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
