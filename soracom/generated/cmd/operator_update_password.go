package cmd

import (
	"encoding/json"
	"io/ioutil"

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

func init() {
	OperatorUpdatePasswordCmd.Flags().StringVar(&OperatorUpdatePasswordCmdCurrentPassword, "current-password", "", TRAPI(""))

	OperatorUpdatePasswordCmd.Flags().StringVar(&OperatorUpdatePasswordCmdNewPassword, "new-password", "", TRAPI(""))

	OperatorUpdatePasswordCmd.Flags().StringVar(&OperatorUpdatePasswordCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	OperatorUpdatePasswordCmd.Flags().StringVar(&OperatorUpdatePasswordCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	OperatorCmd.AddCommand(OperatorUpdatePasswordCmd)
}

// OperatorUpdatePasswordCmd defines 'update-password' subcommand
var OperatorUpdatePasswordCmd = &cobra.Command{
	Use:   "update-password",
	Short: TRAPI("/operators/{operator_id}/password:post:summary"),
	Long:  TRAPI(`/operators/{operator_id}/password:post:description`),
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

		param, err := collectOperatorUpdatePasswordCmdParams()
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

func collectOperatorUpdatePasswordCmdParams() (*apiParams, error) {

	body, err := buildBodyForOperatorUpdatePasswordCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForOperatorUpdatePasswordCmd("/operators/{operator_id}/password"),
		query:       buildQueryForOperatorUpdatePasswordCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForOperatorUpdatePasswordCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", OperatorUpdatePasswordCmdOperatorId, -1)

	return path
}

func buildQueryForOperatorUpdatePasswordCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForOperatorUpdatePasswordCmd() (string, error) {
	if OperatorUpdatePasswordCmdBody != "" {
		if strings.HasPrefix(OperatorUpdatePasswordCmdBody, "@") {
			fname := strings.TrimPrefix(OperatorUpdatePasswordCmdBody, "@")
			// #nosec
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if OperatorUpdatePasswordCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return OperatorUpdatePasswordCmdBody, nil
		}
	}

	result := map[string]interface{}{}

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
