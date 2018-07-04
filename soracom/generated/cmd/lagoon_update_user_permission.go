package cmd

import (
	"encoding/json"
	"io/ioutil"

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
	LagoonUpdateUserPermissionCmd.Flags().StringVar(&LagoonUpdateUserPermissionCmdRole, "role", "", TRAPI(""))

	LagoonUpdateUserPermissionCmd.Flags().Int64Var(&LagoonUpdateUserPermissionCmdLagoonUserId, "lagoon-user-id", 0, TRAPI("Target ID of the lagoon user"))

	LagoonUpdateUserPermissionCmd.Flags().StringVar(&LagoonUpdateUserPermissionCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	LagoonCmd.AddCommand(LagoonUpdateUserPermissionCmd)
}

// LagoonUpdateUserPermissionCmd defines 'update-user-permission' subcommand
var LagoonUpdateUserPermissionCmd = &cobra.Command{
	Use:   "update-user-permission",
	Short: TRAPI("/lagoon/users/{lagoon_user_id}/permission:put:summary"),
	Long:  TRAPI(`/lagoon/users/{lagoon_user_id}/permission:put:description`),
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

		param, err := collectLagoonUpdateUserPermissionCmdParams(ac)
		if err != nil {
			return err
		}

		result, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if result == "" {
			return nil
		}

		return prettyPrintStringAsJSON(result)
	},
}

func collectLagoonUpdateUserPermissionCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForLagoonUpdateUserPermissionCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForLagoonUpdateUserPermissionCmd("/lagoon/users/{lagoon_user_id}/permission"),
		query:       buildQueryForLagoonUpdateUserPermissionCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForLagoonUpdateUserPermissionCmd(path string) string {

	path = strings.Replace(path, "{"+"lagoon_user_id"+"}", sprintf("%d", LagoonUpdateUserPermissionCmdLagoonUserId), -1)

	return path
}

func buildQueryForLagoonUpdateUserPermissionCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForLagoonUpdateUserPermissionCmd() (string, error) {
	if LagoonUpdateUserPermissionCmdBody != "" {
		if strings.HasPrefix(LagoonUpdateUserPermissionCmdBody, "@") {
			fname := strings.TrimPrefix(LagoonUpdateUserPermissionCmdBody, "@")
			// #nosec
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if LagoonUpdateUserPermissionCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return LagoonUpdateUserPermissionCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if LagoonUpdateUserPermissionCmdRole != "" {
		result["role"] = LagoonUpdateUserPermissionCmdRole
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
