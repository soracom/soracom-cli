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

// LagoonUsersCreateCmdRole holds value of 'role' option
var LagoonUsersCreateCmdRole string

// LagoonUsersCreateCmdUserEmail holds value of 'userEmail' option
var LagoonUsersCreateCmdUserEmail string

// LagoonUsersCreateCmdUserPassword holds value of 'userPassword' option
var LagoonUsersCreateCmdUserPassword string

// LagoonUsersCreateCmdBody holds contents of request body to be sent
var LagoonUsersCreateCmdBody string

func InitLagoonUsersCreateCmd() {
	LagoonUsersCreateCmd.Flags().StringVar(&LagoonUsersCreateCmdRole, "role", "", TRAPI("A role that represents the permission."))

	LagoonUsersCreateCmd.Flags().StringVar(&LagoonUsersCreateCmdUserEmail, "user-email", "", TRAPI(""))

	LagoonUsersCreateCmd.Flags().StringVar(&LagoonUsersCreateCmdUserPassword, "user-password", "", TRAPI(""))

	LagoonUsersCreateCmd.Flags().StringVar(&LagoonUsersCreateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	LagoonUsersCreateCmd.RunE = LagoonUsersCreateCmdRunE

	LagoonUsersCmd.AddCommand(LagoonUsersCreateCmd)
}

// LagoonUsersCreateCmd defines 'create' subcommand
var LagoonUsersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: TRAPI("/lagoon/users:post:summary"),
	Long:  TRAPI(`/lagoon/users:post:description`) + "\n\n" + createLinkToAPIReference("Lagoon", "createLagoonUser"),
}

func LagoonUsersCreateCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectLagoonUsersCreateCmdParams(ac)
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

func collectLagoonUsersCreateCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForLagoonUsersCreateCmd()
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

	return &apiParams{
		method:      "POST",
		path:        buildPathForLagoonUsersCreateCmd("/lagoon/users"),
		query:       buildQueryForLagoonUsersCreateCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLagoonUsersCreateCmd(path string) string {

	return path
}

func buildQueryForLagoonUsersCreateCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForLagoonUsersCreateCmd() (string, error) {
	var result map[string]interface{}

	if LagoonUsersCreateCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(LagoonUsersCreateCmdBody, "@") {
			fname := strings.TrimPrefix(LagoonUsersCreateCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if LagoonUsersCreateCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(LagoonUsersCreateCmdBody)
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

	if LagoonUsersCreateCmdRole != "" {
		result["role"] = LagoonUsersCreateCmdRole
	}

	if LagoonUsersCreateCmdUserEmail != "" {
		result["userEmail"] = LagoonUsersCreateCmdUserEmail
	}

	if LagoonUsersCreateCmdUserPassword != "" {
		result["userPassword"] = LagoonUsersCreateCmdUserPassword
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
