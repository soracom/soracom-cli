// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// GadgetsUnsetGroupCmdProductId holds value of 'product_id' option
var GadgetsUnsetGroupCmdProductId string

// GadgetsUnsetGroupCmdSerialNumber holds value of 'serial_number' option
var GadgetsUnsetGroupCmdSerialNumber string

func init() {
	GadgetsUnsetGroupCmd.Flags().StringVar(&GadgetsUnsetGroupCmdProductId, "product-id", "", TRAPI("Product ID of the target gadget."))

	GadgetsUnsetGroupCmd.Flags().StringVar(&GadgetsUnsetGroupCmdSerialNumber, "serial-number", "", TRAPI("Serial Number of the target gadget."))
	GadgetsCmd.AddCommand(GadgetsUnsetGroupCmd)
}

// GadgetsUnsetGroupCmd defines 'unset-group' subcommand
var GadgetsUnsetGroupCmd = &cobra.Command{
	Use:   "unset-group",
	Short: TRAPI("/gadgets/{product_id}/{serial_number}/unset_group:post:summary"),
	Long:  TRAPI(`/gadgets/{product_id}/{serial_number}/unset_group:post:description`) + "\n\n" + createLinkToAPIReference("Gadget", "unsetGadgetGroup"),
	RunE: func(cmd *cobra.Command, args []string) error {

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
		err := authHelper(ac, cmd, args)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		param, err := collectGadgetsUnsetGroupCmdParams(ac)
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
	},
}

func collectGadgetsUnsetGroupCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("product_id", "product-id", "path", parsedBody, GadgetsUnsetGroupCmdProductId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("serial_number", "serial-number", "path", parsedBody, GadgetsUnsetGroupCmdSerialNumber)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForGadgetsUnsetGroupCmd("/gadgets/{product_id}/{serial_number}/unset_group"),
		query:  buildQueryForGadgetsUnsetGroupCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForGadgetsUnsetGroupCmd(path string) string {

	escapedProductId := url.PathEscape(GadgetsUnsetGroupCmdProductId)

	path = strReplace(path, "{"+"product_id"+"}", escapedProductId, -1)

	escapedSerialNumber := url.PathEscape(GadgetsUnsetGroupCmdSerialNumber)

	path = strReplace(path, "{"+"serial_number"+"}", escapedSerialNumber, -1)

	return path
}

func buildQueryForGadgetsUnsetGroupCmd() url.Values {
	result := url.Values{}

	return result
}
