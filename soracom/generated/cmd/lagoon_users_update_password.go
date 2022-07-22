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

// LagoonUsersUpdatePasswordCmdNewPassword holds value of 'newPassword' option
var LagoonUsersUpdatePasswordCmdNewPassword string

// LagoonUsersUpdatePasswordCmdOldPassword holds value of 'oldPassword' option
var LagoonUsersUpdatePasswordCmdOldPassword string

// LagoonUsersUpdatePasswordCmdLagoonUserId holds value of 'lagoon_user_id' option
var LagoonUsersUpdatePasswordCmdLagoonUserId int64

// LagoonUsersUpdatePasswordCmdBody holds contents of request body to be sent
var LagoonUsersUpdatePasswordCmdBody string

func init() {
	LagoonUsersUpdatePasswordCmd.Flags().StringVar(&LagoonUsersUpdatePasswordCmdNewPassword, "new-password", "", TRAPI(""))

	LagoonUsersUpdatePasswordCmd.Flags().StringVar(&LagoonUsersUpdatePasswordCmdOldPassword, "old-password", "", TRAPI(""))

	LagoonUsersUpdatePasswordCmd.Flags().Int64Var(&LagoonUsersUpdatePasswordCmdLagoonUserId, "lagoon-user-id", 0, TRAPI("Target ID of the lagoon user"))

	LagoonUsersUpdatePasswordCmd.Flags().StringVar(&LagoonUsersUpdatePasswordCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	LagoonUsersCmd.AddCommand(LagoonUsersUpdatePasswordCmd)
}

// LagoonUsersUpdatePasswordCmd defines 'update-password' subcommand
var LagoonUsersUpdatePasswordCmd = &cobra.Command{
	Use:   "update-password",
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

		param, err := collectLagoonUsersUpdatePasswordCmdParams(ac)
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

func collectLagoonUsersUpdatePasswordCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForLagoonUsersUpdatePasswordCmd()
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

	err = checkIfRequiredIntegerParameterIsSupplied("lagoon_user_id", "lagoon-user-id", "path", parsedBody, LagoonUsersUpdatePasswordCmdLagoonUserId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForLagoonUsersUpdatePasswordCmd("/lagoon/users/{lagoon_user_id}/password"),
		query:       buildQueryForLagoonUsersUpdatePasswordCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLagoonUsersUpdatePasswordCmd(path string) string {

	path = strReplace(path, "{"+"lagoon_user_id"+"}", url.PathEscape(sprintf("%d", LagoonUsersUpdatePasswordCmdLagoonUserId)), -1)

	return path
}

func buildQueryForLagoonUsersUpdatePasswordCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForLagoonUsersUpdatePasswordCmd() (string, error) {
	var result map[string]interface{}

	if LagoonUsersUpdatePasswordCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(LagoonUsersUpdatePasswordCmdBody, "@") {
			fname := strings.TrimPrefix(LagoonUsersUpdatePasswordCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if LagoonUsersUpdatePasswordCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(LagoonUsersUpdatePasswordCmdBody)
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

	if LagoonUsersUpdatePasswordCmdNewPassword != "" {
		result["newPassword"] = LagoonUsersUpdatePasswordCmdNewPassword
	}

	if LagoonUsersUpdatePasswordCmdOldPassword != "" {
		result["oldPassword"] = LagoonUsersUpdatePasswordCmdOldPassword
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
