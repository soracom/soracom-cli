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

// PayerInformationRegisterCmdCompanyName holds value of 'companyName' option
var PayerInformationRegisterCmdCompanyName string

// PayerInformationRegisterCmdDepartment holds value of 'department' option
var PayerInformationRegisterCmdDepartment string

// PayerInformationRegisterCmdFullName holds value of 'fullName' option
var PayerInformationRegisterCmdFullName string

// PayerInformationRegisterCmdBody holds contents of request body to be sent
var PayerInformationRegisterCmdBody string

func InitPayerInformationRegisterCmd() {
	PayerInformationRegisterCmd.Flags().StringVar(&PayerInformationRegisterCmdCompanyName, "company-name", "", TRAPI("Company name"))

	PayerInformationRegisterCmd.Flags().StringVar(&PayerInformationRegisterCmdDepartment, "department", "", TRAPI("Department."))

	PayerInformationRegisterCmd.Flags().StringVar(&PayerInformationRegisterCmdFullName, "full-name", "", TRAPI("Full name."))

	PayerInformationRegisterCmd.Flags().StringVar(&PayerInformationRegisterCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	PayerInformationRegisterCmd.RunE = PayerInformationRegisterCmdRunE

	PayerInformationCmd.AddCommand(PayerInformationRegisterCmd)
}

// PayerInformationRegisterCmd defines 'register' subcommand
var PayerInformationRegisterCmd = &cobra.Command{
	Use:   "register",
	Short: TRAPI("/payment_statements/payer_information:post:summary"),
	Long:  TRAPI(`/payment_statements/payer_information:post:description`) + "\n\n" + createLinkToAPIReference("Payment", "registerPayerInformation"),
}

func PayerInformationRegisterCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectPayerInformationRegisterCmdParams(ac)
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

func collectPayerInformationRegisterCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForPayerInformationRegisterCmd()
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
		path:        buildPathForPayerInformationRegisterCmd("/payment_statements/payer_information"),
		query:       buildQueryForPayerInformationRegisterCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForPayerInformationRegisterCmd(path string) string {

	return path
}

func buildQueryForPayerInformationRegisterCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForPayerInformationRegisterCmd() (string, error) {
	var result map[string]interface{}

	if PayerInformationRegisterCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(PayerInformationRegisterCmdBody, "@") {
			fname := strings.TrimPrefix(PayerInformationRegisterCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if PayerInformationRegisterCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(PayerInformationRegisterCmdBody)
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

	if PayerInformationRegisterCmdCompanyName != "" {
		result["companyName"] = PayerInformationRegisterCmdCompanyName
	}

	if PayerInformationRegisterCmdDepartment != "" {
		result["department"] = PayerInformationRegisterCmdDepartment
	}

	if PayerInformationRegisterCmdFullName != "" {
		result["fullName"] = PayerInformationRegisterCmdFullName
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
