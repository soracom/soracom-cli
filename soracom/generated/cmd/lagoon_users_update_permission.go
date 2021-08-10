// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"

	"fmt"

	"io/ioutil"

	"mime"
	"net/url"
	"os"

	"strings"

	"github.com/itchyny/gojq"
	"github.com/soracom/soracom-cli/generators/lib"
	"github.com/spf13/cobra"
)

// LagoonUsersUpdatePermissionCmdRole holds value of 'role' option
var LagoonUsersUpdatePermissionCmdRole string

// LagoonUsersUpdatePermissionCmdLagoonUserId holds value of 'lagoon_user_id' option
var LagoonUsersUpdatePermissionCmdLagoonUserId int64

// LagoonUsersUpdatePermissionCmdBody holds contents of request body to be sent
var LagoonUsersUpdatePermissionCmdBody string

func init() {
	LagoonUsersUpdatePermissionCmd.Flags().StringVar(&LagoonUsersUpdatePermissionCmdRole, "role", "", TRAPI(""))

	LagoonUsersUpdatePermissionCmd.Flags().Int64Var(&LagoonUsersUpdatePermissionCmdLagoonUserId, "lagoon-user-id", 0, TRAPI("Target ID of the lagoon user"))

	LagoonUsersUpdatePermissionCmd.Flags().StringVar(&LagoonUsersUpdatePermissionCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	LagoonUsersCmd.AddCommand(LagoonUsersUpdatePermissionCmd)
}

// LagoonUsersUpdatePermissionCmd defines 'update-permission' subcommand
var LagoonUsersUpdatePermissionCmd = &cobra.Command{
	Use:   "update-permission",
	Short: TRAPI("/lagoon/users/{lagoon_user_id}/permission:put:summary"),
	Long:  TRAPI(`/lagoon/users/{lagoon_user_id}/permission:put:description`),
	RunE: func(cmd *cobra.Command, args []string) error {

		var jq *gojq.Query
		if jqString != "" {
			var err error
			jq, err = gojq.Parse(jqString)
			if err != nil {
				return err
			}
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

		body, contentType, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		mediaType, _, err := mime.ParseMediaType(contentType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if jq != nil {
			if mediaType == "application/json" {
				return processJQ(jq, body)
			} else {
				lib.WarnfStderr(TRCLI("cli.tried-jq-on-non-json") + "\n")
			}
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
	body, err := buildBodyForLagoonUsersUpdatePermissionCmd()
	if err != nil {
		return nil, err
	}
	contentType := "application/json"

	if LagoonUsersUpdatePermissionCmdLagoonUserId == 0 {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "lagoon-user-id")
		}

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
			b, err = ioutil.ReadFile(fname)
		} else if LagoonUsersUpdatePermissionCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
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
