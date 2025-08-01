// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// GadgetsTerminateCmdProductId holds value of 'product_id' option
var GadgetsTerminateCmdProductId string

// GadgetsTerminateCmdSerialNumber holds value of 'serial_number' option
var GadgetsTerminateCmdSerialNumber string

func InitGadgetsTerminateCmd() {
	GadgetsTerminateCmd.Flags().StringVar(&GadgetsTerminateCmdProductId, "product-id", "", TRAPI("Product ID of the target Gadget API compatible device.- 'wimax': Soracom Cloud Camera Services Cellular Pack."))

	GadgetsTerminateCmd.Flags().StringVar(&GadgetsTerminateCmdSerialNumber, "serial-number", "", TRAPI("Serial Number of the target Gadget API compatible device."))

	GadgetsTerminateCmd.RunE = GadgetsTerminateCmdRunE

	GadgetsCmd.AddCommand(GadgetsTerminateCmd)
}

// GadgetsTerminateCmd defines 'terminate' subcommand
var GadgetsTerminateCmd = &cobra.Command{
	Use:   "terminate",
	Short: TRAPI("/gadgets/{product_id}/{serial_number}/terminate:post:summary"),
	Long:  TRAPI(`/gadgets/{product_id}/{serial_number}/terminate:post:description`) + "\n\n" + createLinkToAPIReference("Gadget", "terminateGadget"),
}

func GadgetsTerminateCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectGadgetsTerminateCmdParams(ac)
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

func collectGadgetsTerminateCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("product_id", "product-id", "path", parsedBody, GadgetsTerminateCmdProductId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("serial_number", "serial-number", "path", parsedBody, GadgetsTerminateCmdSerialNumber)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForGadgetsTerminateCmd("/gadgets/{product_id}/{serial_number}/terminate"),
		query:  buildQueryForGadgetsTerminateCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForGadgetsTerminateCmd(path string) string {

	escapedProductId := url.PathEscape(GadgetsTerminateCmdProductId)

	path = strReplace(path, "{"+"product_id"+"}", escapedProductId, -1)

	escapedSerialNumber := url.PathEscape(GadgetsTerminateCmdSerialNumber)

	path = strReplace(path, "{"+"serial_number"+"}", escapedSerialNumber, -1)

	return path
}

func buildQueryForGadgetsTerminateCmd() url.Values {
	result := url.Values{}

	return result
}
