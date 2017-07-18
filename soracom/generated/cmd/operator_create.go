package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// OperatorCreateCmdEmail holds value of 'email' option
var OperatorCreateCmdEmail string

// OperatorCreateCmdPassword holds value of 'password' option
var OperatorCreateCmdPassword string

// OperatorCreateCmdBody holds contents of request body to be sent
var OperatorCreateCmdBody string

func init() {
	OperatorCreateCmd.Flags().StringVar(&OperatorCreateCmdEmail, "email", "", TRAPI(""))

	OperatorCreateCmd.Flags().StringVar(&OperatorCreateCmdPassword, "password", "", TRAPI(""))

	OperatorCreateCmd.Flags().StringVar(&OperatorCreateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	OperatorCmd.AddCommand(OperatorCreateCmd)
}

// OperatorCreateCmd defines 'create' subcommand
var OperatorCreateCmd = &cobra.Command{
	Use:   "create",
	Short: TRAPI("/operators:post:summary"),
	Long:  TRAPI(`/operators:post:description`),
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

		param, err := collectOperatorCreateCmdParams()
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

func collectOperatorCreateCmdParams() (*apiParams, error) {

	body, err := buildBodyForOperatorCreateCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForOperatorCreateCmd("/operators"),
		query:       buildQueryForOperatorCreateCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForOperatorCreateCmd(path string) string {

	return path
}

func buildQueryForOperatorCreateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForOperatorCreateCmd() (string, error) {
	if OperatorCreateCmdBody != "" {
		if strings.HasPrefix(OperatorCreateCmdBody, "@") {
			fname := strings.TrimPrefix(OperatorCreateCmdBody, "@")
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if OperatorCreateCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return OperatorCreateCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if OperatorCreateCmdEmail != "" {
		result["email"] = OperatorCreateCmdEmail
	}

	if OperatorCreateCmdPassword != "" {
		result["password"] = OperatorCreateCmdPassword
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
