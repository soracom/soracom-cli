package cmd

import (
	"encoding/json"
	"io/ioutil"

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

func init() {
	PayerInformationRegisterCmd.Flags().StringVar(&PayerInformationRegisterCmdCompanyName, "company-name", "", TRAPI(""))

	PayerInformationRegisterCmd.Flags().StringVar(&PayerInformationRegisterCmdDepartment, "department", "", TRAPI(""))

	PayerInformationRegisterCmd.Flags().StringVar(&PayerInformationRegisterCmdFullName, "full-name", "", TRAPI(""))

	PayerInformationRegisterCmd.Flags().StringVar(&PayerInformationRegisterCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	PayerInformationCmd.AddCommand(PayerInformationRegisterCmd)
}

// PayerInformationRegisterCmd defines 'register' subcommand
var PayerInformationRegisterCmd = &cobra.Command{
	Use:   "register",
	Short: TRAPI("/payment_statements/payer_information:post:summary"),
	Long:  TRAPI(`/payment_statements/payer_information:post:description`),
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

		param, err := collectPayerInformationRegisterCmdParams()
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

func collectPayerInformationRegisterCmdParams() (*apiParams, error) {

	body, err := buildBodyForPayerInformationRegisterCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForPayerInformationRegisterCmd("/payment_statements/payer_information"),
		query:       buildQueryForPayerInformationRegisterCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForPayerInformationRegisterCmd(path string) string {

	return path
}

func buildQueryForPayerInformationRegisterCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForPayerInformationRegisterCmd() (string, error) {
	if PayerInformationRegisterCmdBody != "" {
		if strings.HasPrefix(PayerInformationRegisterCmdBody, "@") {
			fname := strings.TrimPrefix(PayerInformationRegisterCmdBody, "@")
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if PayerInformationRegisterCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return PayerInformationRegisterCmdBody, nil
		}
	}

	result := map[string]interface{}{}

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
