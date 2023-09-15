// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"

	"strings"

	"github.com/spf13/cobra"
)

// UsersDefaultPermissionsUpdateCmdOperatorId holds value of 'operator_id' option
var UsersDefaultPermissionsUpdateCmdOperatorId string

// UsersDefaultPermissionsUpdateCmdPermissions holds value of 'permissions' option
var UsersDefaultPermissionsUpdateCmdPermissions string

// UsersDefaultPermissionsUpdateCmdBody holds contents of request body to be sent
var UsersDefaultPermissionsUpdateCmdBody string

func InitUsersDefaultPermissionsUpdateCmd() {
	UsersDefaultPermissionsUpdateCmd.Flags().StringVar(&UsersDefaultPermissionsUpdateCmdOperatorId, "operator-id", "", TRAPI("Operator ID"))

	UsersDefaultPermissionsUpdateCmd.Flags().StringVar(&UsersDefaultPermissionsUpdateCmdPermissions, "permissions", "", TRAPI("JSON string of permissions"))

	UsersDefaultPermissionsUpdateCmd.Flags().StringVar(&UsersDefaultPermissionsUpdateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	UsersDefaultPermissionsUpdateCmd.RunE = UsersDefaultPermissionsUpdateCmdRunE

	UsersDefaultPermissionsCmd.AddCommand(UsersDefaultPermissionsUpdateCmd)
}

// UsersDefaultPermissionsUpdateCmd defines 'update' subcommand
var UsersDefaultPermissionsUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: TRAPI("/operators/{operator_id}/users/default_permissions:put:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users/default_permissions:put:description`) + "\n\n" + createLinkToAPIReference("User", "updateDefaultPermissions"),
}

func UsersDefaultPermissionsUpdateCmdRunE(cmd *cobra.Command, args []string) error {

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
	err := ac.getAPICredentials()
	if err != nil {
		cmd.SilenceUsage = true
		return err
	}

	param, err := collectUsersDefaultPermissionsUpdateCmdParams(ac)
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
}

func collectUsersDefaultPermissionsUpdateCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	if UsersDefaultPermissionsUpdateCmdOperatorId == "" {
		UsersDefaultPermissionsUpdateCmdOperatorId = ac.apiCredentials.getOperatorID()
	}

	body, err = buildBodyForUsersDefaultPermissionsUpdateCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("permissions", "permissions", "body", parsedBody, UsersDefaultPermissionsUpdateCmdPermissions)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForUsersDefaultPermissionsUpdateCmd("/operators/{operator_id}/users/default_permissions"),
		query:       buildQueryForUsersDefaultPermissionsUpdateCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForUsersDefaultPermissionsUpdateCmd(path string) string {

	escapedOperatorId := url.PathEscape(UsersDefaultPermissionsUpdateCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	return path
}

func buildQueryForUsersDefaultPermissionsUpdateCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForUsersDefaultPermissionsUpdateCmd() (string, error) {
	var result map[string]interface{}

	if UsersDefaultPermissionsUpdateCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(UsersDefaultPermissionsUpdateCmdBody, "@") {
			fname := strings.TrimPrefix(UsersDefaultPermissionsUpdateCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if UsersDefaultPermissionsUpdateCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(UsersDefaultPermissionsUpdateCmdBody)
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

	if UsersDefaultPermissionsUpdateCmdPermissions != "" {
		result["permissions"] = UsersDefaultPermissionsUpdateCmdPermissions
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
