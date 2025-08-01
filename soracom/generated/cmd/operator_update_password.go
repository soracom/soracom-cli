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

// OperatorUpdatePasswordCmdCurrentPassword holds value of 'currentPassword' option
var OperatorUpdatePasswordCmdCurrentPassword string

// OperatorUpdatePasswordCmdNewPassword holds value of 'newPassword' option
var OperatorUpdatePasswordCmdNewPassword string

// OperatorUpdatePasswordCmdOperatorId holds value of 'operator_id' option
var OperatorUpdatePasswordCmdOperatorId string

// OperatorUpdatePasswordCmdBody holds contents of request body to be sent
var OperatorUpdatePasswordCmdBody string

func InitOperatorUpdatePasswordCmd() {
	OperatorUpdatePasswordCmd.Flags().StringVar(&OperatorUpdatePasswordCmdCurrentPassword, "current-password", "", TRAPI(""))

	OperatorUpdatePasswordCmd.Flags().StringVar(&OperatorUpdatePasswordCmdNewPassword, "new-password", "", TRAPI(""))

	OperatorUpdatePasswordCmd.Flags().StringVar(&OperatorUpdatePasswordCmdOperatorId, "operator-id", "", TRAPI("Operator ID."))

	OperatorUpdatePasswordCmd.Flags().StringVar(&OperatorUpdatePasswordCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	OperatorUpdatePasswordCmd.RunE = OperatorUpdatePasswordCmdRunE

	OperatorCmd.AddCommand(OperatorUpdatePasswordCmd)
}

// OperatorUpdatePasswordCmd defines 'update-password' subcommand
var OperatorUpdatePasswordCmd = &cobra.Command{
	Use:   "update-password",
	Short: TRAPI("/operators/{operator_id}/password:post:summary"),
	Long:  TRAPI(`/operators/{operator_id}/password:post:description`) + "\n\n" + createLinkToAPIReference("Operator", "updateOperatorPassword"),
}

func OperatorUpdatePasswordCmdRunE(cmd *cobra.Command, args []string) error {

	if len(args) > 0 {
		return fmt.Errorf("unexpected arguments passed => %v", args)
	}

	opt := &apiClientOptions{
		BasePath: "/v1",
		Language: getSelectedLanguage(),
		Profile:  getProfileIfExists(),
	}

	ac := newAPIClient(opt)
	if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
		ac.SetVerbose(true)
	}
	err := ac.getAPICredentials()
	if err != nil {
		cmd.SilenceUsage = true
		return err
	}

	param, err := collectOperatorUpdatePasswordCmdParams(ac)
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

func collectOperatorUpdatePasswordCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	if OperatorUpdatePasswordCmdOperatorId == "" {
		OperatorUpdatePasswordCmdOperatorId = ac.apiCredentials.getOperatorID()
	}

	body, err = buildBodyForOperatorUpdatePasswordCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("currentPassword", "current-password", "body", parsedBody, OperatorUpdatePasswordCmdCurrentPassword)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("newPassword", "new-password", "body", parsedBody, OperatorUpdatePasswordCmdNewPassword)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForOperatorUpdatePasswordCmd("/operators/{operator_id}/password"),
		query:       buildQueryForOperatorUpdatePasswordCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForOperatorUpdatePasswordCmd(path string) string {

	escapedOperatorId := url.PathEscape(OperatorUpdatePasswordCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	return path
}

func buildQueryForOperatorUpdatePasswordCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForOperatorUpdatePasswordCmd() (string, error) {
	var result map[string]interface{}

	if OperatorUpdatePasswordCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(OperatorUpdatePasswordCmdBody, "@") {
			fname := strings.TrimPrefix(OperatorUpdatePasswordCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if OperatorUpdatePasswordCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(OperatorUpdatePasswordCmdBody)
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

	if OperatorUpdatePasswordCmdCurrentPassword != "" {
		result["currentPassword"] = OperatorUpdatePasswordCmdCurrentPassword
	}

	if OperatorUpdatePasswordCmdNewPassword != "" {
		result["newPassword"] = OperatorUpdatePasswordCmdNewPassword
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
