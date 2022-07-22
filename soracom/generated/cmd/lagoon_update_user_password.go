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

// LagoonUpdateUserPasswordCmdNewPassword holds value of 'newPassword' option
var LagoonUpdateUserPasswordCmdNewPassword string

// LagoonUpdateUserPasswordCmdOldPassword holds value of 'oldPassword' option
var LagoonUpdateUserPasswordCmdOldPassword string

// LagoonUpdateUserPasswordCmdLagoonUserId holds value of 'lagoon_user_id' option
var LagoonUpdateUserPasswordCmdLagoonUserId int64

// LagoonUpdateUserPasswordCmdBody holds contents of request body to be sent
var LagoonUpdateUserPasswordCmdBody string

func init() {
	LagoonUpdateUserPasswordCmd.Flags().StringVar(&LagoonUpdateUserPasswordCmdNewPassword, "new-password", "", TRAPI(""))

	LagoonUpdateUserPasswordCmd.Flags().StringVar(&LagoonUpdateUserPasswordCmdOldPassword, "old-password", "", TRAPI(""))

	LagoonUpdateUserPasswordCmd.Flags().Int64Var(&LagoonUpdateUserPasswordCmdLagoonUserId, "lagoon-user-id", 0, TRAPI("Target ID of the lagoon user"))

	LagoonUpdateUserPasswordCmd.Flags().StringVar(&LagoonUpdateUserPasswordCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	LagoonCmd.AddCommand(LagoonUpdateUserPasswordCmd)
}

// LagoonUpdateUserPasswordCmd defines 'update-user-password' subcommand
var LagoonUpdateUserPasswordCmd = &cobra.Command{
	Use:   "update-user-password",
	Short: TRAPI("/lagoon/users/{lagoon_user_id}/password:put:summary"),
	Long:  TRAPI(`/lagoon/users/{lagoon_user_id}/password:put:description`) + "\n\n" + createLinkToAPIReference("Lagoon", "updateLagoonUserPassword"),
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) > 0 {
			return fmt.Errorf("unexpected arguments passed => %v", args)
		}

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

		param, err := collectLagoonUpdateUserPasswordCmdParams(ac)
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

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectLagoonUpdateUserPasswordCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForLagoonUpdateUserPasswordCmd()
	if err != nil {
		return nil, err
	}
	contentType := "application/json"

	if contentType == "application/json" {
		err = json.Unmarshal([]byte(body), &parsedBody)
		if err != nil {
			return nil, fmt.Errorf("invalid json format specified for `--body` parameter: %s", err)
		}
	}

	err = checkIfRequiredIntegerParameterIsSupplied("lagoon_user_id", "lagoon-user-id", "path", parsedBody, LagoonUpdateUserPasswordCmdLagoonUserId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForLagoonUpdateUserPasswordCmd("/lagoon/users/{lagoon_user_id}/password"),
		query:       buildQueryForLagoonUpdateUserPasswordCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLagoonUpdateUserPasswordCmd(path string) string {

	path = strReplace(path, "{"+"lagoon_user_id"+"}", url.PathEscape(sprintf("%d", LagoonUpdateUserPasswordCmdLagoonUserId)), -1)

	return path
}

func buildQueryForLagoonUpdateUserPasswordCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForLagoonUpdateUserPasswordCmd() (string, error) {
	var result map[string]interface{}

	if LagoonUpdateUserPasswordCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(LagoonUpdateUserPasswordCmdBody, "@") {
			fname := strings.TrimPrefix(LagoonUpdateUserPasswordCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if LagoonUpdateUserPasswordCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(LagoonUpdateUserPasswordCmdBody)
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

	if LagoonUpdateUserPasswordCmdNewPassword != "" {
		result["newPassword"] = LagoonUpdateUserPasswordCmdNewPassword
	}

	if LagoonUpdateUserPasswordCmdOldPassword != "" {
		result["oldPassword"] = LagoonUpdateUserPasswordCmdOldPassword
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
