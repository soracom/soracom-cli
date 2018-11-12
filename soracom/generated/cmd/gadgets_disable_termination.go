package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// GadgetsDisableTerminationCmdProductId holds value of 'product_id' option
var GadgetsDisableTerminationCmdProductId string

// GadgetsDisableTerminationCmdSerialNumber holds value of 'serial_number' option
var GadgetsDisableTerminationCmdSerialNumber string

func init() {
	GadgetsDisableTerminationCmd.Flags().StringVar(&GadgetsDisableTerminationCmdProductId, "product-id", "", TRAPI("Product ID of the target gadget."))

	GadgetsDisableTerminationCmd.Flags().StringVar(&GadgetsDisableTerminationCmdSerialNumber, "serial-number", "", TRAPI("Serial Number of the target gadget."))

	GadgetsCmd.AddCommand(GadgetsDisableTerminationCmd)
}

// GadgetsDisableTerminationCmd defines 'disable-termination' subcommand
var GadgetsDisableTerminationCmd = &cobra.Command{
	Use:   "disable-termination",
	Short: TRAPI("/gadgets/{product_id}/{serial_number}/disable_termination:post:summary"),
	Long:  TRAPI(`/gadgets/{product_id}/{serial_number}/disable_termination:post:description`),
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

		param, err := collectGadgetsDisableTerminationCmdParams(ac)
		if err != nil {
			return err
		}

		_, body, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		return prettyPrintStringAsJSON(body)
	},
}

func collectGadgetsDisableTerminationCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForGadgetsDisableTerminationCmd("/gadgets/{product_id}/{serial_number}/disable_termination"),
		query:  buildQueryForGadgetsDisableTerminationCmd(),
	}, nil
}

func buildPathForGadgetsDisableTerminationCmd(path string) string {

	path = strings.Replace(path, "{"+"product_id"+"}", GadgetsDisableTerminationCmdProductId, -1)

	path = strings.Replace(path, "{"+"serial_number"+"}", GadgetsDisableTerminationCmdSerialNumber, -1)

	return path
}

func buildQueryForGadgetsDisableTerminationCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
