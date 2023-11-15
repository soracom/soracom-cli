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

// SandboxOrdersShipCmdOperatorId holds value of 'operatorId' option
var SandboxOrdersShipCmdOperatorId string

// SandboxOrdersShipCmdOrderId holds value of 'orderId' option
var SandboxOrdersShipCmdOrderId string

// SandboxOrdersShipCmdBody holds contents of request body to be sent
var SandboxOrdersShipCmdBody string

func InitSandboxOrdersShipCmd() {
	SandboxOrdersShipCmd.Flags().StringVar(&SandboxOrdersShipCmdOperatorId, "operator-id", "", TRAPI(""))

	SandboxOrdersShipCmd.Flags().StringVar(&SandboxOrdersShipCmdOrderId, "order-id", "", TRAPI(""))

	SandboxOrdersShipCmd.Flags().StringVar(&SandboxOrdersShipCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	SandboxOrdersShipCmd.RunE = SandboxOrdersShipCmdRunE

	SandboxOrdersCmd.AddCommand(SandboxOrdersShipCmd)
}

// SandboxOrdersShipCmd defines 'ship' subcommand
var SandboxOrdersShipCmd = &cobra.Command{
	Use:   "ship",
	Short: TRAPI("/sandbox/orders/ship:post:summary"),
	Long:  TRAPI(`/sandbox/orders/ship:post:description`) + "\n\n" + createLinkToAPIReference("API Sandbox: Order", "sandboxShipOrder"),
}

func SandboxOrdersShipCmdRunE(cmd *cobra.Command, args []string) error {

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
	err := ac.getAPICredentials()
	if err != nil {
		cmd.SilenceUsage = true
		return err
	}

	param, err := collectSandboxOrdersShipCmdParams(ac)
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

func collectSandboxOrdersShipCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	if SandboxOrdersShipCmdOperatorId == "" {
		SandboxOrdersShipCmdOperatorId = ac.apiCredentials.getOperatorID()
	}

	body, err = buildBodyForSandboxOrdersShipCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("orderId", "order-id", "body", parsedBody, SandboxOrdersShipCmdOrderId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSandboxOrdersShipCmd("/sandbox/orders/ship"),
		query:       buildQueryForSandboxOrdersShipCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSandboxOrdersShipCmd(path string) string {

	return path
}

func buildQueryForSandboxOrdersShipCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSandboxOrdersShipCmd() (string, error) {
	var result map[string]interface{}

	if SandboxOrdersShipCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SandboxOrdersShipCmdBody, "@") {
			fname := strings.TrimPrefix(SandboxOrdersShipCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if SandboxOrdersShipCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(SandboxOrdersShipCmdBody)
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

	if SandboxOrdersShipCmdOperatorId != "" {
		result["operatorId"] = SandboxOrdersShipCmdOperatorId
	}

	if SandboxOrdersShipCmdOrderId != "" {
		result["orderId"] = SandboxOrdersShipCmdOrderId
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
