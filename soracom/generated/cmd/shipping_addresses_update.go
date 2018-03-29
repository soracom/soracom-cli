package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// ShippingAddressesUpdateCmdAddressLine1 holds value of 'addressLine1' option
var ShippingAddressesUpdateCmdAddressLine1 string

// ShippingAddressesUpdateCmdAddressLine2 holds value of 'addressLine2' option
var ShippingAddressesUpdateCmdAddressLine2 string

// ShippingAddressesUpdateCmdBuilding holds value of 'building' option
var ShippingAddressesUpdateCmdBuilding string

// ShippingAddressesUpdateCmdCity holds value of 'city' option
var ShippingAddressesUpdateCmdCity string

// ShippingAddressesUpdateCmdCompanyName holds value of 'companyName' option
var ShippingAddressesUpdateCmdCompanyName string

// ShippingAddressesUpdateCmdDepartment holds value of 'department' option
var ShippingAddressesUpdateCmdDepartment string

// ShippingAddressesUpdateCmdFullName holds value of 'fullName' option
var ShippingAddressesUpdateCmdFullName string

// ShippingAddressesUpdateCmdOperatorId holds value of 'operator_id' option
var ShippingAddressesUpdateCmdOperatorId string

// ShippingAddressesUpdateCmdPhoneNumber holds value of 'phoneNumber' option
var ShippingAddressesUpdateCmdPhoneNumber string

// ShippingAddressesUpdateCmdShippingAddressId holds value of 'shipping_address_id' option
var ShippingAddressesUpdateCmdShippingAddressId string

// ShippingAddressesUpdateCmdState holds value of 'state' option
var ShippingAddressesUpdateCmdState string

// ShippingAddressesUpdateCmdZipCode holds value of 'zipCode' option
var ShippingAddressesUpdateCmdZipCode string

// ShippingAddressesUpdateCmdBody holds contents of request body to be sent
var ShippingAddressesUpdateCmdBody string

func init() {
	ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdAddressLine1, "address-line1", "", TRAPI(""))

	ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdAddressLine2, "address-line2", "", TRAPI(""))

	ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdBuilding, "building", "", TRAPI(""))

	ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdCity, "city", "", TRAPI(""))

	ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdCompanyName, "company-name", "", TRAPI(""))

	ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdDepartment, "department", "", TRAPI(""))

	ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdFullName, "full-name", "", TRAPI(""))

	ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdOperatorId, "operator-id", "", TRAPI("Operator ID"))

	ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdPhoneNumber, "phone-number", "", TRAPI(""))

	ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdShippingAddressId, "shipping-address-id", "", TRAPI("shipping_address_id"))

	ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdState, "state", "", TRAPI(""))

	ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdZipCode, "zip-code", "", TRAPI(""))

	ShippingAddressesUpdateCmd.Flags().StringVar(&ShippingAddressesUpdateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	ShippingAddressesCmd.AddCommand(ShippingAddressesUpdateCmd)
}

// ShippingAddressesUpdateCmd defines 'update' subcommand
var ShippingAddressesUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: TRAPI("/operators/{operator_id}/shipping_addresses/{shipping_address_id}:put:summary"),
	Long:  TRAPI(`/operators/{operator_id}/shipping_addresses/{shipping_address_id}:put:description`),
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

		param, err := collectShippingAddressesUpdateCmdParams()
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

func collectShippingAddressesUpdateCmdParams() (*apiParams, error) {

	body, err := buildBodyForShippingAddressesUpdateCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForShippingAddressesUpdateCmd("/operators/{operator_id}/shipping_addresses/{shipping_address_id}"),
		query:       buildQueryForShippingAddressesUpdateCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForShippingAddressesUpdateCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", ShippingAddressesUpdateCmdOperatorId, -1)

	path = strings.Replace(path, "{"+"shipping_address_id"+"}", ShippingAddressesUpdateCmdShippingAddressId, -1)

	return path
}

func buildQueryForShippingAddressesUpdateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForShippingAddressesUpdateCmd() (string, error) {
	if ShippingAddressesUpdateCmdBody != "" {
		if strings.HasPrefix(ShippingAddressesUpdateCmdBody, "@") {
			fname := strings.TrimPrefix(ShippingAddressesUpdateCmdBody, "@")
			// #nosec
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if ShippingAddressesUpdateCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return ShippingAddressesUpdateCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if ShippingAddressesUpdateCmdAddressLine1 != "" {
		result["addressLine1"] = ShippingAddressesUpdateCmdAddressLine1
	}

	if ShippingAddressesUpdateCmdAddressLine2 != "" {
		result["addressLine2"] = ShippingAddressesUpdateCmdAddressLine2
	}

	if ShippingAddressesUpdateCmdBuilding != "" {
		result["building"] = ShippingAddressesUpdateCmdBuilding
	}

	if ShippingAddressesUpdateCmdCity != "" {
		result["city"] = ShippingAddressesUpdateCmdCity
	}

	if ShippingAddressesUpdateCmdCompanyName != "" {
		result["companyName"] = ShippingAddressesUpdateCmdCompanyName
	}

	if ShippingAddressesUpdateCmdDepartment != "" {
		result["department"] = ShippingAddressesUpdateCmdDepartment
	}

	if ShippingAddressesUpdateCmdFullName != "" {
		result["fullName"] = ShippingAddressesUpdateCmdFullName
	}

	if ShippingAddressesUpdateCmdPhoneNumber != "" {
		result["phoneNumber"] = ShippingAddressesUpdateCmdPhoneNumber
	}

	if ShippingAddressesUpdateCmdState != "" {
		result["state"] = ShippingAddressesUpdateCmdState
	}

	if ShippingAddressesUpdateCmdZipCode != "" {
		result["zipCode"] = ShippingAddressesUpdateCmdZipCode
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
