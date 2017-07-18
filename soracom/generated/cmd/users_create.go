package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// UsersCreateCmdDescription holds value of 'description' option
var UsersCreateCmdDescription string

// UsersCreateCmdOperatorId holds value of 'operator_id' option
var UsersCreateCmdOperatorId string

// UsersCreateCmdUserName holds value of 'user_name' option
var UsersCreateCmdUserName string

// UsersCreateCmdBody holds contents of request body to be sent
var UsersCreateCmdBody string

func init() {
	UsersCreateCmd.Flags().StringVar(&UsersCreateCmdDescription, "description", "", TRAPI(""))

	UsersCreateCmd.Flags().StringVar(&UsersCreateCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	UsersCreateCmd.Flags().StringVar(&UsersCreateCmdUserName, "user-name", "", TRAPI("user_name"))

	UsersCreateCmd.Flags().StringVar(&UsersCreateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	UsersCmd.AddCommand(UsersCreateCmd)
}

// UsersCreateCmd defines 'create' subcommand
var UsersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: TRAPI("/operators/{operator_id}/users/{user_name}:post:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users/{user_name}:post:description`),
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

		param, err := collectUsersCreateCmdParams()
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

func collectUsersCreateCmdParams() (*apiParams, error) {

	body, err := buildBodyForUsersCreateCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForUsersCreateCmd("/operators/{operator_id}/users/{user_name}"),
		query:       buildQueryForUsersCreateCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForUsersCreateCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", UsersCreateCmdOperatorId, -1)

	path = strings.Replace(path, "{"+"user_name"+"}", UsersCreateCmdUserName, -1)

	return path
}

func buildQueryForUsersCreateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForUsersCreateCmd() (string, error) {
	if UsersCreateCmdBody != "" {
		if strings.HasPrefix(UsersCreateCmdBody, "@") {
			fname := strings.TrimPrefix(UsersCreateCmdBody, "@")
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if UsersCreateCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return UsersCreateCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if UsersCreateCmdDescription != "" {
		result["description"] = UsersCreateCmdDescription
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
