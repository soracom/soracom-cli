// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// GadgetsEnableTerminationCmdProductId holds value of 'product_id' option
var GadgetsEnableTerminationCmdProductId string

// GadgetsEnableTerminationCmdSerialNumber holds value of 'serial_number' option
var GadgetsEnableTerminationCmdSerialNumber string

func init() {
	GadgetsEnableTerminationCmd.Flags().StringVar(&GadgetsEnableTerminationCmdProductId, "product-id", "", TRAPI("Product ID of the target gadget."))

	GadgetsEnableTerminationCmd.Flags().StringVar(&GadgetsEnableTerminationCmdSerialNumber, "serial-number", "", TRAPI("Serial Number of the target gadget."))
	GadgetsCmd.AddCommand(GadgetsEnableTerminationCmd)
}

// GadgetsEnableTerminationCmd defines 'enable-termination' subcommand
var GadgetsEnableTerminationCmd = &cobra.Command{
	Use:   "enable-termination",
	Short: TRAPI("/gadgets/{product_id}/{serial_number}/enable_termination:post:summary"),
	Long:  TRAPI(`/gadgets/{product_id}/{serial_number}/enable_termination:post:description`),
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

		param, err := collectGadgetsEnableTerminationCmdParams(ac)
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

func collectGadgetsEnableTerminationCmdParams(ac *apiClient) (*apiParams, error) {
	if GadgetsEnableTerminationCmdProductId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "product-id")
	}

	if GadgetsEnableTerminationCmdSerialNumber == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "serial-number")
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForGadgetsEnableTerminationCmd("/gadgets/{product_id}/{serial_number}/enable_termination"),
		query:  buildQueryForGadgetsEnableTerminationCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForGadgetsEnableTerminationCmd(path string) string {

	escapedProductId := url.PathEscape(GadgetsEnableTerminationCmdProductId)

	path = strReplace(path, "{"+"product_id"+"}", escapedProductId, -1)

	escapedSerialNumber := url.PathEscape(GadgetsEnableTerminationCmdSerialNumber)

	path = strReplace(path, "{"+"serial_number"+"}", escapedSerialNumber, -1)

	return path
}

func buildQueryForGadgetsEnableTerminationCmd() url.Values {
	result := url.Values{}

	return result
}
