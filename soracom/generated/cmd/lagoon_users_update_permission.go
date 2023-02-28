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

// LagoonUsersUpdatePermissionCmdRole holds value of 'role' option
var LagoonUsersUpdatePermissionCmdRole string

// LagoonUsersUpdatePermissionCmdLagoonUserId holds value of 'lagoon_user_id' option
var LagoonUsersUpdatePermissionCmdLagoonUserId int64

// LagoonUsersUpdatePermissionCmdBody holds contents of request body to be sent
var LagoonUsersUpdatePermissionCmdBody string

func init() {
	LagoonUsersUpdatePermissionCmd.Flags().StringVar(&LagoonUsersUpdatePermissionCmdRole, "role", "", TRAPI("A role that represents the permission."))

	LagoonUsersUpdatePermissionCmd.Flags().Int64Var(&LagoonUsersUpdatePermissionCmdLagoonUserId, "lagoon-user-id", 0, TRAPI("Target ID of the lagoon user"))

	LagoonUsersUpdatePermissionCmd.Flags().StringVar(&LagoonUsersUpdatePermissionCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	LagoonUsersCmd.AddCommand(LagoonUsersUpdatePermissionCmd)
}

// LagoonUsersUpdatePermissionCmd defines 'update-permission' subcommand
var LagoonUsersUpdatePermissionCmd = &cobra.Command{
	Use:   "update-permission",
	Short: TRAPI("/lagoon/users/{lagoon_user_id}/permission:put:summary"),
	Long:  TRAPI(`/lagoon/users/{lagoon_user_id}/permission:put:description`) + "\n\n" + createLinkToAPIReference("Lagoon", "updateLagoonUserPermission"),
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

		param, err := collectLagoonUsersUpdatePermissionCmdParams(ac)
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

func collectLagoonUsersUpdatePermissionCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForLagoonUsersUpdatePermissionCmd()
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

	err = checkIfRequiredIntegerParameterIsSupplied("lagoon_user_id", "lagoon-user-id", "path", parsedBody, LagoonUsersUpdatePermissionCmdLagoonUserId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForLagoonUsersUpdatePermissionCmd("/lagoon/users/{lagoon_user_id}/permission"),
		query:       buildQueryForLagoonUsersUpdatePermissionCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLagoonUsersUpdatePermissionCmd(path string) string {

	path = strReplace(path, "{"+"lagoon_user_id"+"}", url.PathEscape(sprintf("%d", LagoonUsersUpdatePermissionCmdLagoonUserId)), -1)

	return path
}

func buildQueryForLagoonUsersUpdatePermissionCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForLagoonUsersUpdatePermissionCmd() (string, error) {
	var result map[string]interface{}

	if LagoonUsersUpdatePermissionCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(LagoonUsersUpdatePermissionCmdBody, "@") {
			fname := strings.TrimPrefix(LagoonUsersUpdatePermissionCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if LagoonUsersUpdatePermissionCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(LagoonUsersUpdatePermissionCmdBody)
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

	if LagoonUsersUpdatePermissionCmdRole != "" {
		result["role"] = LagoonUsersUpdatePermissionCmdRole
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
