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

// SandboxInitCmdAuthKey holds value of 'authKey' option
var SandboxInitCmdAuthKey string

// SandboxInitCmdAuthKeyId holds value of 'authKeyId' option
var SandboxInitCmdAuthKeyId string

// SandboxInitCmdEmail holds value of 'email' option
var SandboxInitCmdEmail string

// SandboxInitCmdPassword holds value of 'password' option
var SandboxInitCmdPassword string

// SandboxInitCmdCoverageTypes holds multiple values of 'coverageTypes' option
var SandboxInitCmdCoverageTypes []string

// SandboxInitCmdRegisterPaymentMethod holds value of 'registerPaymentMethod' option
var SandboxInitCmdRegisterPaymentMethod bool

// SandboxInitCmdBody holds contents of request body to be sent
var SandboxInitCmdBody string

func InitSandboxInitCmd() {
	SandboxInitCmd.Flags().StringVar(&SandboxInitCmdAuthKey, "auth-key", "", TRAPI(""))

	SandboxInitCmd.Flags().StringVar(&SandboxInitCmdAuthKeyId, "auth-key-id", "", TRAPI(""))

	SandboxInitCmd.Flags().StringVar(&SandboxInitCmdEmail, "email", "", TRAPI(""))

	SandboxInitCmd.Flags().StringVar(&SandboxInitCmdPassword, "password", "", TRAPI(""))

	SandboxInitCmd.Flags().StringSliceVar(&SandboxInitCmdCoverageTypes, "coverage-types", []string{}, TRAPI("Coverage type.- 'g': Global coverage- 'jp': Japan coverage"))

	SandboxInitCmd.Flags().BoolVar(&SandboxInitCmdRegisterPaymentMethod, "register-payment-method", true, TRAPI(""))

	SandboxInitCmd.Flags().StringVar(&SandboxInitCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	SandboxInitCmd.RunE = SandboxInitCmdRunE

	SandboxCmd.AddCommand(SandboxInitCmd)
}

// SandboxInitCmd defines 'init' subcommand
var SandboxInitCmd = &cobra.Command{
	Use:   "init",
	Short: TRAPI("/sandbox/init:post:summary"),
	Long:  TRAPI(`/sandbox/init:post:description`) + "\n\n" + createLinkToAPIReference("API Sandbox: Operator", "sandboxInitializeOperator"),
}

func SandboxInitCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSandboxInitCmdParams(ac)
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

func collectSandboxInitCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForSandboxInitCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("authKey", "auth-key", "body", parsedBody, SandboxInitCmdAuthKey)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("authKeyId", "auth-key-id", "body", parsedBody, SandboxInitCmdAuthKeyId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("email", "email", "body", parsedBody, SandboxInitCmdEmail)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("password", "password", "body", parsedBody, SandboxInitCmdPassword)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSandboxInitCmd("/sandbox/init"),
		query:       buildQueryForSandboxInitCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSandboxInitCmd(path string) string {

	return path
}

func buildQueryForSandboxInitCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSandboxInitCmd() (string, error) {
	var result map[string]interface{}

	if SandboxInitCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SandboxInitCmdBody, "@") {
			fname := strings.TrimPrefix(SandboxInitCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if SandboxInitCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(SandboxInitCmdBody)
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

	if SandboxInitCmdAuthKey != "" {
		result["authKey"] = SandboxInitCmdAuthKey
	}

	if SandboxInitCmdAuthKeyId != "" {
		result["authKeyId"] = SandboxInitCmdAuthKeyId
	}

	if SandboxInitCmdEmail != "" {
		result["email"] = SandboxInitCmdEmail
	}

	if SandboxInitCmdPassword != "" {
		result["password"] = SandboxInitCmdPassword
	}

	if len(SandboxInitCmdCoverageTypes) != 0 {
		result["coverageTypes"] = SandboxInitCmdCoverageTypes
	}

	if SandboxInitCmd.Flags().Lookup("register-payment-method").Changed {
		result["registerPaymentMethod"] = SandboxInitCmdRegisterPaymentMethod
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
