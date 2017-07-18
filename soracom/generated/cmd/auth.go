package cmd

import (
	"encoding/json"
	"io/ioutil"

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

	AuthCmd.Flags().StringVar(&AuthCmdOperatorId, "operator-id", "", TRAPI(""))

	AuthCmd.Flags().StringVar(&AuthCmdPassword, "password", "", TRAPI(""))

	AuthCmd.Flags().StringVar(&AuthCmdUserName, "user-name", "", TRAPI(""))

	AuthCmd.Flags().Int64Var(&AuthCmdTokenTimeoutSeconds, "token-timeout-seconds", 0, TRAPI(""))

	AuthCmd.Flags().StringVar(&AuthCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	RootCmd.AddCommand(AuthCmd)
}

// AuthCmd defines 'auth' subcommand
var AuthCmd = &cobra.Command{
	Use:   "auth",
	Short: TRAPI("/auth:post:summary"),
	Long:  TRAPI(`/auth:post:description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			BasePath: "/v1",
			Language: getSelectedLanguage(),
		}

		ac := newAPIClient(opt)
		if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
			ac.SetVerbose(true)
		}

		param, err := collectAuthCmdParams()
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

func collectAuthCmdParams() (*apiParams, error) {

	body, err := buildBodyForAuthCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForAuthCmd("/auth"),
		query:       buildQueryForAuthCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForAuthCmd(path string) string {

	return path
}

func buildQueryForAuthCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForAuthCmd() (string, error) {
	if AuthCmdBody != "" {
		if strings.HasPrefix(AuthCmdBody, "@") {
			fname := strings.TrimPrefix(AuthCmdBody, "@")
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if AuthCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return AuthCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if AuthCmdAuthKey != "" {
		result["authKey"] = AuthCmdAuthKey
	}

	if AuthCmdAuthKeyId != "" {
		result["authKeyId"] = AuthCmdAuthKeyId
	}

	if AuthCmdEmail != "" {
		result["email"] = AuthCmdEmail
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

	if AuthCmdTokenTimeoutSeconds != 0 {
		result["tokenTimeoutSeconds"] = AuthCmdTokenTimeoutSeconds
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
