// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"

	"fmt"

	"io/ioutil"

	"net/url"
	"os"

	"strings"

	"github.com/spf13/cobra"
)

// OperatorCreateCompanyInformationCmdAddressLine1 holds value of 'addressLine1' option
var OperatorCreateCompanyInformationCmdAddressLine1 string

// OperatorCreateCompanyInformationCmdAddressLine2 holds value of 'addressLine2' option
var OperatorCreateCompanyInformationCmdAddressLine2 string

// OperatorCreateCompanyInformationCmdBuilding holds value of 'building' option
var OperatorCreateCompanyInformationCmdBuilding string

// OperatorCreateCompanyInformationCmdCity holds value of 'city' option
var OperatorCreateCompanyInformationCmdCity string

// OperatorCreateCompanyInformationCmdCompanyName holds value of 'companyName' option
var OperatorCreateCompanyInformationCmdCompanyName string

// OperatorCreateCompanyInformationCmdContactPersonName holds value of 'contactPersonName' option
var OperatorCreateCompanyInformationCmdContactPersonName string

// OperatorCreateCompanyInformationCmdCountryCode holds value of 'countryCode' option
var OperatorCreateCompanyInformationCmdCountryCode string

// OperatorCreateCompanyInformationCmdDepartment holds value of 'department' option
var OperatorCreateCompanyInformationCmdDepartment string

// OperatorCreateCompanyInformationCmdOperatorId holds value of 'operator_id' option
var OperatorCreateCompanyInformationCmdOperatorId string

// OperatorCreateCompanyInformationCmdPhoneNumber holds value of 'phoneNumber' option
var OperatorCreateCompanyInformationCmdPhoneNumber string

// OperatorCreateCompanyInformationCmdState holds value of 'state' option
var OperatorCreateCompanyInformationCmdState string

// OperatorCreateCompanyInformationCmdVatIdentificationNumber holds value of 'vatIdentificationNumber' option
var OperatorCreateCompanyInformationCmdVatIdentificationNumber string

// OperatorCreateCompanyInformationCmdZipCode holds value of 'zipCode' option
var OperatorCreateCompanyInformationCmdZipCode string

// OperatorCreateCompanyInformationCmdBody holds contents of request body to be sent
var OperatorCreateCompanyInformationCmdBody string

func init() {
	OperatorCreateCompanyInformationCmd.Flags().StringVar(&OperatorCreateCompanyInformationCmdAddressLine1, "address-line1", "", TRAPI(""))

	OperatorCreateCompanyInformationCmd.Flags().StringVar(&OperatorCreateCompanyInformationCmdAddressLine2, "address-line2", "", TRAPI(""))

	OperatorCreateCompanyInformationCmd.Flags().StringVar(&OperatorCreateCompanyInformationCmdBuilding, "building", "", TRAPI(""))

	OperatorCreateCompanyInformationCmd.Flags().StringVar(&OperatorCreateCompanyInformationCmdCity, "city", "", TRAPI(""))

	OperatorCreateCompanyInformationCmd.Flags().StringVar(&OperatorCreateCompanyInformationCmdCompanyName, "company-name", "", TRAPI(""))

	OperatorCreateCompanyInformationCmd.Flags().StringVar(&OperatorCreateCompanyInformationCmdContactPersonName, "contact-person-name", "", TRAPI(""))

	OperatorCreateCompanyInformationCmd.Flags().StringVar(&OperatorCreateCompanyInformationCmdCountryCode, "country-code", "", TRAPI(""))

	OperatorCreateCompanyInformationCmd.Flags().StringVar(&OperatorCreateCompanyInformationCmdDepartment, "department", "", TRAPI(""))

	OperatorCreateCompanyInformationCmd.Flags().StringVar(&OperatorCreateCompanyInformationCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	OperatorCreateCompanyInformationCmd.Flags().StringVar(&OperatorCreateCompanyInformationCmdPhoneNumber, "phone-number", "", TRAPI(""))

	OperatorCreateCompanyInformationCmd.Flags().StringVar(&OperatorCreateCompanyInformationCmdState, "state", "", TRAPI(""))

	OperatorCreateCompanyInformationCmd.Flags().StringVar(&OperatorCreateCompanyInformationCmdVatIdentificationNumber, "vat-identification-number", "", TRAPI(""))

	OperatorCreateCompanyInformationCmd.Flags().StringVar(&OperatorCreateCompanyInformationCmdZipCode, "zip-code", "", TRAPI(""))

	OperatorCreateCompanyInformationCmd.Flags().StringVar(&OperatorCreateCompanyInformationCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	OperatorCmd.AddCommand(OperatorCreateCompanyInformationCmd)
}

// OperatorCreateCompanyInformationCmd defines 'create-company-information' subcommand
var OperatorCreateCompanyInformationCmd = &cobra.Command{
	Use:   "create-company-information",
	Short: TRAPI("/operators/{operator_id}/company_information:post:summary"),
	Long:  TRAPI(`/operators/{operator_id}/company_information:post:description`),
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

		param, err := collectOperatorCreateCompanyInformationCmdParams(ac)
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

		if queryString != "" {
			return processQuery(queryString, body)
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectOperatorCreateCompanyInformationCmdParams(ac *apiClient) (*apiParams, error) {
	if OperatorCreateCompanyInformationCmdOperatorId == "" {
		OperatorCreateCompanyInformationCmdOperatorId = ac.OperatorID
	}

	body, err := buildBodyForOperatorCreateCompanyInformationCmd()
	if err != nil {
		return nil, err
	}
	contentType := "application/json"

	if OperatorCreateCompanyInformationCmdCompanyName == "" {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "company-name")
		}

	}

	if OperatorCreateCompanyInformationCmdContactPersonName == "" {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "contact-person-name")
		}

	}

	if OperatorCreateCompanyInformationCmdCountryCode == "" {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "country-code")
		}

	}

	if OperatorCreateCompanyInformationCmdDepartment == "" {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "department")
		}

	}

	if OperatorCreateCompanyInformationCmdPhoneNumber == "" {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "phone-number")
		}

	}

	if OperatorCreateCompanyInformationCmdZipCode == "" {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "zip-code")
		}

	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForOperatorCreateCompanyInformationCmd("/operators/{operator_id}/company_information"),
		query:       buildQueryForOperatorCreateCompanyInformationCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForOperatorCreateCompanyInformationCmd(path string) string {

	escapedOperatorId := url.PathEscape(OperatorCreateCompanyInformationCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	return path
}

func buildQueryForOperatorCreateCompanyInformationCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForOperatorCreateCompanyInformationCmd() (string, error) {
	var result map[string]interface{}

	if OperatorCreateCompanyInformationCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(OperatorCreateCompanyInformationCmdBody, "@") {
			fname := strings.TrimPrefix(OperatorCreateCompanyInformationCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if OperatorCreateCompanyInformationCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(OperatorCreateCompanyInformationCmdBody)
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

	if OperatorCreateCompanyInformationCmdAddressLine1 != "" {
		result["addressLine1"] = OperatorCreateCompanyInformationCmdAddressLine1
	}

	if OperatorCreateCompanyInformationCmdAddressLine2 != "" {
		result["addressLine2"] = OperatorCreateCompanyInformationCmdAddressLine2
	}

	if OperatorCreateCompanyInformationCmdBuilding != "" {
		result["building"] = OperatorCreateCompanyInformationCmdBuilding
	}

	if OperatorCreateCompanyInformationCmdCity != "" {
		result["city"] = OperatorCreateCompanyInformationCmdCity
	}

	if OperatorCreateCompanyInformationCmdCompanyName != "" {
		result["companyName"] = OperatorCreateCompanyInformationCmdCompanyName
	}

	if OperatorCreateCompanyInformationCmdContactPersonName != "" {
		result["contactPersonName"] = OperatorCreateCompanyInformationCmdContactPersonName
	}

	if OperatorCreateCompanyInformationCmdCountryCode != "" {
		result["countryCode"] = OperatorCreateCompanyInformationCmdCountryCode
	}

	if OperatorCreateCompanyInformationCmdDepartment != "" {
		result["department"] = OperatorCreateCompanyInformationCmdDepartment
	}

	if OperatorCreateCompanyInformationCmdPhoneNumber != "" {
		result["phoneNumber"] = OperatorCreateCompanyInformationCmdPhoneNumber
	}

	if OperatorCreateCompanyInformationCmdState != "" {
		result["state"] = OperatorCreateCompanyInformationCmdState
	}

	if OperatorCreateCompanyInformationCmdVatIdentificationNumber != "" {
		result["vatIdentificationNumber"] = OperatorCreateCompanyInformationCmdVatIdentificationNumber
	}

	if OperatorCreateCompanyInformationCmdZipCode != "" {
		result["zipCode"] = OperatorCreateCompanyInformationCmdZipCode
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
