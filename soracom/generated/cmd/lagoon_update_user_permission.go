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

// LagoonUpdateUserPermissionCmdRole holds value of 'role' option
var LagoonUpdateUserPermissionCmdRole string

// LagoonUpdateUserPermissionCmdLagoonUserId holds value of 'lagoon_user_id' option
var LagoonUpdateUserPermissionCmdLagoonUserId int64

// LagoonUpdateUserPermissionCmdBody holds contents of request body to be sent
var LagoonUpdateUserPermissionCmdBody string

func init() {
	LagoonUpdateUserPermissionCmd.Flags().StringVar(&LagoonUpdateUserPermissionCmdRole, "role", "", TRAPI("A role that represents the permission."))

	LagoonUpdateUserPermissionCmd.Flags().Int64Var(&LagoonUpdateUserPermissionCmdLagoonUserId, "lagoon-user-id", 0, TRAPI("Target ID of the lagoon user"))

	LagoonUpdateUserPermissionCmd.Flags().StringVar(&LagoonUpdateUserPermissionCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	LagoonCmd.AddCommand(LagoonUpdateUserPermissionCmd)
}

// LagoonUpdateUserPermissionCmd defines 'update-user-permission' subcommand
var LagoonUpdateUserPermissionCmd = &cobra.Command{
	Use:   "update-user-permission",
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

		param, err := collectLagoonUpdateUserPermissionCmdParams(ac)
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

func collectLagoonUpdateUserPermissionCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForLagoonUpdateUserPermissionCmd()
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

	err = checkIfRequiredIntegerParameterIsSupplied("lagoon_user_id", "lagoon-user-id", "path", parsedBody, LagoonUpdateUserPermissionCmdLagoonUserId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForLagoonUpdateUserPermissionCmd("/lagoon/users/{lagoon_user_id}/permission"),
		query:       buildQueryForLagoonUpdateUserPermissionCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLagoonUpdateUserPermissionCmd(path string) string {

	path = strReplace(path, "{"+"lagoon_user_id"+"}", url.PathEscape(sprintf("%d", LagoonUpdateUserPermissionCmdLagoonUserId)), -1)

	return path
}

func buildQueryForLagoonUpdateUserPermissionCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForLagoonUpdateUserPermissionCmd() (string, error) {
	var result map[string]interface{}

	if LagoonUpdateUserPermissionCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(LagoonUpdateUserPermissionCmdBody, "@") {
			fname := strings.TrimPrefix(LagoonUpdateUserPermissionCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if LagoonUpdateUserPermissionCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(LagoonUpdateUserPermissionCmdBody)
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

	if LagoonUpdateUserPermissionCmdRole != "" {
		result["role"] = LagoonUpdateUserPermissionCmdRole
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
