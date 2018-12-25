package cmd

import (
	"encoding/json"

	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// UsersMfaVerifyCmdMfaOTPCode holds value of 'mfaOTPCode' option
var UsersMfaVerifyCmdMfaOTPCode string

// UsersMfaVerifyCmdOperatorId holds value of 'operator_id' option
var UsersMfaVerifyCmdOperatorId string

// UsersMfaVerifyCmdUserName holds value of 'user_name' option
var UsersMfaVerifyCmdUserName string

// UsersMfaVerifyCmdBody holds contents of request body to be sent
var UsersMfaVerifyCmdBody string

func init() {
	UsersMfaVerifyCmd.Flags().StringVar(&UsersMfaVerifyCmdMfaOTPCode, "mfa-otpcode", "", TRAPI(""))

	UsersMfaVerifyCmd.Flags().StringVar(&UsersMfaVerifyCmdOperatorId, "operator-id", "", TRAPI("Operator ID"))

	UsersMfaVerifyCmd.Flags().StringVar(&UsersMfaVerifyCmdUserName, "user-name", "", TRAPI("SAM user name"))

	UsersMfaVerifyCmd.Flags().StringVar(&UsersMfaVerifyCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	UsersMfaCmd.AddCommand(UsersMfaVerifyCmd)
}

// UsersMfaVerifyCmd defines 'verify' subcommand
var UsersMfaVerifyCmd = &cobra.Command{
	Use:   "verify",
	Short: TRAPI("/operators/{operator_id}/users/{user_name}/mfa/verify:post:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users/{user_name}/mfa/verify:post:description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			BasePath: "/v1",
			Language: getSelectedLanguage(),
		}

		ac := newAPIClient(opt)
		if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
			ac.SetVerbose(true)
		}

		param, err := collectUsersMfaVerifyCmdParams(ac)
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

func collectUsersMfaVerifyCmdParams(ac *apiClient) (*apiParams, error) {

	if UsersMfaVerifyCmdOperatorId == "" {
		UsersMfaVerifyCmdOperatorId = ac.OperatorID
	}

	body, err := buildBodyForUsersMfaVerifyCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForUsersMfaVerifyCmd("/operators/{operator_id}/users/{user_name}/mfa/verify"),
		query:       buildQueryForUsersMfaVerifyCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForUsersMfaVerifyCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", UsersMfaVerifyCmdOperatorId, -1)

	path = strings.Replace(path, "{"+"user_name"+"}", UsersMfaVerifyCmdUserName, -1)

	return path
}

func buildQueryForUsersMfaVerifyCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForUsersMfaVerifyCmd() (string, error) {
	var result map[string]interface{}

	if UsersMfaVerifyCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(UsersMfaVerifyCmdBody, "@") {
			fname := strings.TrimPrefix(UsersMfaVerifyCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if UsersMfaVerifyCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(UsersMfaVerifyCmdBody)
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

	if UsersMfaVerifyCmdMfaOTPCode != "" {
		result["mfaOTPCode"] = UsersMfaVerifyCmdMfaOTPCode
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
