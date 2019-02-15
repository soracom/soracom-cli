package cmd

import (
	"encoding/json"

	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LagoonUsersUpdatePasswordCmdNewPassword holds value of 'newPassword' option
var LagoonUsersUpdatePasswordCmdNewPassword string

// LagoonUsersUpdatePasswordCmdOldPassword holds value of 'oldPassword' option
var LagoonUsersUpdatePasswordCmdOldPassword string

// LagoonUsersUpdatePasswordCmdUserEmail holds value of 'userEmail' option
var LagoonUsersUpdatePasswordCmdUserEmail string

// LagoonUsersUpdatePasswordCmdLagoonUserId holds value of 'lagoon_user_id' option
var LagoonUsersUpdatePasswordCmdLagoonUserId int64

// LagoonUsersUpdatePasswordCmdBody holds contents of request body to be sent
var LagoonUsersUpdatePasswordCmdBody string

func init() {
	LagoonUsersUpdatePasswordCmd.Flags().StringVar(&LagoonUsersUpdatePasswordCmdNewPassword, "new-password", "", TRAPI(""))

	LagoonUsersUpdatePasswordCmd.Flags().StringVar(&LagoonUsersUpdatePasswordCmdOldPassword, "old-password", "", TRAPI(""))

	LagoonUsersUpdatePasswordCmd.Flags().StringVar(&LagoonUsersUpdatePasswordCmdUserEmail, "user-email", "", TRAPI(""))

	LagoonUsersUpdatePasswordCmd.Flags().Int64Var(&LagoonUsersUpdatePasswordCmdLagoonUserId, "lagoon-user-id", 0, TRAPI("Target ID of the lagoon user"))

	LagoonUsersUpdatePasswordCmd.Flags().StringVar(&LagoonUsersUpdatePasswordCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	LagoonUsersCmd.AddCommand(LagoonUsersUpdatePasswordCmd)
}

// LagoonUsersUpdatePasswordCmd defines 'update-password' subcommand
var LagoonUsersUpdatePasswordCmd = &cobra.Command{
	Use:   "update-password",
	Short: TRAPI("/lagoon/users/{lagoon_user_id}/password:put:summary"),
	Long:  TRAPI(`/lagoon/users/{lagoon_user_id}/password:put:description`),
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

		param, err := collectLagoonUsersUpdatePasswordCmdParams(ac)
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

func collectLagoonUsersUpdatePasswordCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForLagoonUsersUpdatePasswordCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForLagoonUsersUpdatePasswordCmd("/lagoon/users/{lagoon_user_id}/password"),
		query:       buildQueryForLagoonUsersUpdatePasswordCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForLagoonUsersUpdatePasswordCmd(path string) string {

	path = strings.Replace(path, "{"+"lagoon_user_id"+"}", sprintf("%d", LagoonUsersUpdatePasswordCmdLagoonUserId), -1)

	return path
}

func buildQueryForLagoonUsersUpdatePasswordCmd() string {
	result := []string{}

	return strings.Join(result, "&")
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

	if LagoonUsersUpdatePasswordCmdUserEmail != "" {
		result["userEmail"] = LagoonUsersUpdatePasswordCmdUserEmail
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
