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

// UsersAttachRoleCmdOperatorId holds value of 'operator_id' option
var UsersAttachRoleCmdOperatorId string

// UsersAttachRoleCmdRoleId holds value of 'roleId' option
var UsersAttachRoleCmdRoleId string

// UsersAttachRoleCmdUserName holds value of 'user_name' option
var UsersAttachRoleCmdUserName string

// UsersAttachRoleCmdBody holds contents of request body to be sent
var UsersAttachRoleCmdBody string

func init() {
	UsersAttachRoleCmd.Flags().StringVar(&UsersAttachRoleCmdOperatorId, "operator-id", "", TRAPI("Operator ID"))

	UsersAttachRoleCmd.Flags().StringVar(&UsersAttachRoleCmdRoleId, "role-id", "", TRAPI(""))

	UsersAttachRoleCmd.Flags().StringVar(&UsersAttachRoleCmdUserName, "user-name", "", TRAPI("user_name"))

	UsersAttachRoleCmd.Flags().StringVar(&UsersAttachRoleCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	UsersCmd.AddCommand(UsersAttachRoleCmd)
}

// UsersAttachRoleCmd defines 'attach-role' subcommand
var UsersAttachRoleCmd = &cobra.Command{
	Use:   "attach-role",
	Short: TRAPI("/operators/{operator_id}/users/{user_name}/roles:post:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users/{user_name}/roles:post:description`) + "\n\n" + createLinkToAPIReference("Role", "attachRole"),
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

		param, err := collectUsersAttachRoleCmdParams(ac)
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

func collectUsersAttachRoleCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	if UsersAttachRoleCmdOperatorId == "" {
		UsersAttachRoleCmdOperatorId = ac.OperatorID
	}

	body, err = buildBodyForUsersAttachRoleCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("user_name", "user-name", "path", parsedBody, UsersAttachRoleCmdUserName)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForUsersAttachRoleCmd("/operators/{operator_id}/users/{user_name}/roles"),
		query:       buildQueryForUsersAttachRoleCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForUsersAttachRoleCmd(path string) string {

	escapedOperatorId := url.PathEscape(UsersAttachRoleCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	escapedUserName := url.PathEscape(UsersAttachRoleCmdUserName)

	path = strReplace(path, "{"+"user_name"+"}", escapedUserName, -1)

	return path
}

func buildQueryForUsersAttachRoleCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForUsersAttachRoleCmd() (string, error) {
	var result map[string]interface{}

	if UsersAttachRoleCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(UsersAttachRoleCmdBody, "@") {
			fname := strings.TrimPrefix(UsersAttachRoleCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if UsersAttachRoleCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(UsersAttachRoleCmdBody)
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

	if UsersAttachRoleCmdRoleId != "" {
		result["roleId"] = UsersAttachRoleCmdRoleId
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
