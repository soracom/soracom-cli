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

// AuthCmdAuthKey holds value of 'authKey' option
var AuthCmdAuthKey string

// AuthCmdAuthKeyId holds value of 'authKeyId' option
var AuthCmdAuthKeyId string

// AuthCmdEmail holds value of 'email' option
var AuthCmdEmail string

// AuthCmdMfaOTPCode holds value of 'mfaOTPCode' option
var AuthCmdMfaOTPCode string

// AuthCmdOperatorId holds value of 'operatorId' option
var AuthCmdOperatorId string

// AuthCmdPassword holds value of 'password' option
var AuthCmdPassword string

// AuthCmdUserName holds value of 'userName' option
var AuthCmdUserName string

// AuthCmdTokenTimeoutSeconds holds value of 'tokenTimeoutSeconds' option
var AuthCmdTokenTimeoutSeconds int64

// AuthCmdBody holds contents of request body to be sent
var AuthCmdBody string

func init() {
	AuthCmd.Flags().StringVar(&AuthCmdAuthKey, "auth-key", "", TRAPI(""))

	AuthCmd.Flags().StringVar(&AuthCmdAuthKeyId, "auth-key-id", "", TRAPI(""))

	AuthCmd.Flags().StringVar(&AuthCmdEmail, "email", "", TRAPI(""))

	AuthCmd.Flags().StringVar(&AuthCmdMfaOTPCode, "mfa-otpcode", "", TRAPI(""))

	AuthCmd.Flags().StringVar(&AuthCmdOperatorId, "operator-id", "", TRAPI(""))

	AuthCmd.Flags().StringVar(&AuthCmdPassword, "password", "", TRAPI(""))

	AuthCmd.Flags().StringVar(&AuthCmdUserName, "user-name", "", TRAPI(""))

	AuthCmd.Flags().Int64Var(&AuthCmdTokenTimeoutSeconds, "token-timeout-seconds", 86400, TRAPI(""))

	AuthCmd.Flags().StringVar(&AuthCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	RootCmd.AddCommand(AuthCmd)
}

// AuthCmd defines 'auth' subcommand
var AuthCmd = &cobra.Command{
	Use:   "auth",
	Short: TRAPI("/auth:post:summary"),
	Long:  TRAPI(`/auth:post:description`) + "\n\n" + createLinkToAPIReference("Auth", "auth"),
	RunE: func(cmd *cobra.Command, args []string) error {

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

		param, err := collectAuthCmdParams(ac)
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
	},
}

func collectAuthCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForAuthCmd()
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
		path:        buildPathForAuthCmd("/auth"),
		query:       buildQueryForAuthCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForAuthCmd(path string) string {

	return path
}

func buildQueryForAuthCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForAuthCmd() (string, error) {
	var result map[string]interface{}

	if AuthCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(AuthCmdBody, "@") {
			fname := strings.TrimPrefix(AuthCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if AuthCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(AuthCmdBody)
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

	if AuthCmdAuthKey != "" {
		result["authKey"] = AuthCmdAuthKey
	}

	if AuthCmdAuthKeyId != "" {
		result["authKeyId"] = AuthCmdAuthKeyId
	}

	if AuthCmdEmail != "" {
		result["email"] = AuthCmdEmail
	}

	if AuthCmdMfaOTPCode != "" {
		result["mfaOTPCode"] = AuthCmdMfaOTPCode
	}

	if AuthCmdOperatorId != "" {
		result["operatorId"] = AuthCmdOperatorId
	}

	if AuthCmdPassword != "" {
		result["password"] = AuthCmdPassword
	}

	if AuthCmdUserName != "" {
		result["userName"] = AuthCmdUserName
	}

	if AuthCmdTokenTimeoutSeconds != 86400 {

		result["tokenTimeoutSeconds"] = AuthCmdTokenTimeoutSeconds

	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
