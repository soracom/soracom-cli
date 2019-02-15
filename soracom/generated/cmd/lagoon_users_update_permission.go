package cmd

import (
	"encoding/json"

	"io/ioutil"

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

		_, body, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		return prettyPrintStringAsJSON(body)
	},
}

func collectLagoonUsersUpdatePermissionCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForLagoonUsersUpdatePermissionCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForLagoonUsersUpdatePermissionCmd("/lagoon/users/{lagoon_user_id}/permission"),
		query:       buildQueryForLagoonUsersUpdatePermissionCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForLagoonUsersUpdatePermissionCmd(path string) string {

	path = strings.Replace(path, "{"+"lagoon_user_id"+"}", sprintf("%d", LagoonUsersUpdatePermissionCmdLagoonUserId), -1)

	return path
}

func buildQueryForLagoonUsersUpdatePermissionCmd() string {
	result := []string{}

	return strings.Join(result, "&")
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
