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

// OperatorCreateCmdEmail holds value of 'email' option
var OperatorCreateCmdEmail string

// OperatorCreateCmdPassword holds value of 'password' option
var OperatorCreateCmdPassword string

// OperatorCreateCmdBody holds contents of request body to be sent
var OperatorCreateCmdBody string

func InitOperatorCreateCmd() {
	OperatorCreateCmd.Flags().StringVar(&OperatorCreateCmdEmail, "email", "", TRAPI(""))

	OperatorCreateCmd.Flags().StringVar(&OperatorCreateCmdPassword, "password", "", TRAPI("Requirements for password: At least 8 characters (max 100), a mixture of lowercase (a-z), uppercase (A-Z), number (0-9) letters. You can use symbols."))

	OperatorCreateCmd.Flags().StringVar(&OperatorCreateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	OperatorCreateCmd.RunE = OperatorCreateCmdRunE

	OperatorCmd.AddCommand(OperatorCreateCmd)
}

// OperatorCreateCmd defines 'create' subcommand
var OperatorCreateCmd = &cobra.Command{
	Use:   "create",
	Short: TRAPI("/operators:post:summary"),
	Long:  TRAPI(`/operators:post:description`) + "\n\n" + createLinkToAPIReference("Operator", "createOperator"),
}

func OperatorCreateCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectOperatorCreateCmdParams(ac)
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

func collectOperatorCreateCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForOperatorCreateCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("email", "email", "body", parsedBody, OperatorCreateCmdEmail)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("password", "password", "body", parsedBody, OperatorCreateCmdPassword)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForOperatorCreateCmd("/operators"),
		query:       buildQueryForOperatorCreateCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForOperatorCreateCmd(path string) string {

	return path
}

func buildQueryForOperatorCreateCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForOperatorCreateCmd() (string, error) {
	var result map[string]interface{}

	if OperatorCreateCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(OperatorCreateCmdBody, "@") {
			fname := strings.TrimPrefix(OperatorCreateCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if OperatorCreateCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(OperatorCreateCmdBody)
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
