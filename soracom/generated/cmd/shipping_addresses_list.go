// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"mime"
	"net/url"
	"os"

	"github.com/itchyny/gojq"
	"github.com/soracom/soracom-cli/generators/lib"
	"github.com/spf13/cobra"
)

// ShippingAddressesListCmdOperatorId holds value of 'operator_id' option
var ShippingAddressesListCmdOperatorId string

func init() {
	ShippingAddressesListCmd.Flags().StringVar(&ShippingAddressesListCmdOperatorId, "operator-id", "", TRAPI("operator_id"))
	ShippingAddressesCmd.AddCommand(ShippingAddressesListCmd)
}

// ShippingAddressesListCmd defines 'list' subcommand
var ShippingAddressesListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/operators/{operator_id}/shipping_addresses:get:summary"),
	Long:  TRAPI(`/operators/{operator_id}/shipping_addresses:get:description`),
	RunE: func(cmd *cobra.Command, args []string) error {

		var jq *gojq.Query
		if jqString != "" {
			var err error
			jq, err = gojq.Parse(jqString)
			if err != nil {
				return err
			}
		}

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

		param, err := collectShippingAddressesListCmdParams(ac)
		if err != nil {
			return err
		}

		body, contentType, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		mediaType, _, err := mime.ParseMediaType(contentType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if jq != nil {
			if mediaType == "application/json" {
				return processJQ(jq, body)
			} else {
				lib.WarnfStderr(TRCLI("cli.tried-jq-on-non-json") + "\n")
			}
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectShippingAddressesListCmdParams(ac *apiClient) (*apiParams, error) {
	if ShippingAddressesListCmdOperatorId == "" {
		ShippingAddressesListCmdOperatorId = ac.OperatorID
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForShippingAddressesListCmd("/operators/{operator_id}/shipping_addresses"),
		query:  buildQueryForShippingAddressesListCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForShippingAddressesListCmd(path string) string {

	escapedOperatorId := url.PathEscape(ShippingAddressesListCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	return path
}

func buildQueryForShippingAddressesListCmd() url.Values {
	result := url.Values{}

	return result
}
