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

func InitGadgetsEnableTerminationCmd() {
	GadgetsEnableTerminationCmd.Flags().StringVar(&GadgetsEnableTerminationCmdProductId, "product-id", "", TRAPI("Product ID of the target gadget."))

	GadgetsEnableTerminationCmd.Flags().StringVar(&GadgetsEnableTerminationCmdSerialNumber, "serial-number", "", TRAPI("Serial Number of the target gadget."))

	GadgetsEnableTerminationCmd.RunE = GadgetsEnableTerminationCmdRunE

	GadgetsCmd.AddCommand(GadgetsEnableTerminationCmd)
}

// GadgetsEnableTerminationCmd defines 'enable-termination' subcommand
var GadgetsEnableTerminationCmd = &cobra.Command{
	Use:   "enable-termination",
	Short: TRAPI("/gadgets/{product_id}/{serial_number}/enable_termination:post:summary"),
	Long:  TRAPI(`/gadgets/{product_id}/{serial_number}/enable_termination:post:description`) + "\n\n" + createLinkToAPIReference("Gadget", "enableTerminationOnGadget"),
}

func GadgetsEnableTerminationCmdRunE(cmd *cobra.Command, args []string) error {

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

	if rawOutput {
		_, err = os.Stdout.Write([]byte(body))
	} else {
		return prettyPrintStringAsJSON(body)
	}
	return err
}

func collectGadgetsEnableTerminationCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("product_id", "product-id", "path", parsedBody, GadgetsEnableTerminationCmdProductId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("serial_number", "serial-number", "path", parsedBody, GadgetsEnableTerminationCmdSerialNumber)
	if err != nil {
		return nil, err
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
