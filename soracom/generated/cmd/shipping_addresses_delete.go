package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// ShippingAddressesDeleteCmdOperatorId holds value of 'operator_id' option
var ShippingAddressesDeleteCmdOperatorId string

// ShippingAddressesDeleteCmdShippingAddressId holds value of 'shipping_address_id' option
var ShippingAddressesDeleteCmdShippingAddressId string

func init() {
	ShippingAddressesDeleteCmd.Flags().StringVar(&ShippingAddressesDeleteCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	ShippingAddressesDeleteCmd.Flags().StringVar(&ShippingAddressesDeleteCmdShippingAddressId, "shipping-address-id", "", TRAPI("shipping_address_id"))

	ShippingAddressesCmd.AddCommand(ShippingAddressesDeleteCmd)
}

// ShippingAddressesDeleteCmd defines 'delete' subcommand
var ShippingAddressesDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: TRAPI("/operators/{operator_id}/shipping_addresses/{shipping_address_id}:delete:summary"),
	Long:  TRAPI(`/operators/{operator_id}/shipping_addresses/{shipping_address_id}:delete:description`),
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

		param, err := collectShippingAddressesDeleteCmdParams(ac)
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

func collectShippingAddressesDeleteCmdParams(ac *apiClient) (*apiParams, error) {

	if ShippingAddressesDeleteCmdOperatorId == "" {
		ShippingAddressesDeleteCmdOperatorId = ac.OperatorID
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForShippingAddressesDeleteCmd("/operators/{operator_id}/shipping_addresses/{shipping_address_id}"),
		query:  buildQueryForShippingAddressesDeleteCmd(),
	}, nil
}

func buildPathForShippingAddressesDeleteCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", ShippingAddressesDeleteCmdOperatorId, -1)

	path = strings.Replace(path, "{"+"shipping_address_id"+"}", ShippingAddressesDeleteCmdShippingAddressId, -1)

	return path
}

func buildQueryForShippingAddressesDeleteCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
