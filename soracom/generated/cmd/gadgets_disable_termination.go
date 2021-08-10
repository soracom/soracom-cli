// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"mime"
	"net/url"
	"os"

	"github.com/itchyny/gojq"
	"github.com/soracom/soracom-cli/generators/lib"
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

		param, err := collectGadgetsDisableTerminationCmdParams(ac)
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

func collectGadgetsDisableTerminationCmdParams(ac *apiClient) (*apiParams, error) {
	if GadgetsDisableTerminationCmdProductId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "product-id")
	}

	if GadgetsDisableTerminationCmdSerialNumber == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "serial-number")
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForGadgetsDisableTerminationCmd("/gadgets/{product_id}/{serial_number}/disable_termination"),
		query:  buildQueryForGadgetsDisableTerminationCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForGadgetsDisableTerminationCmd(path string) string {

	escapedProductId := url.PathEscape(GadgetsDisableTerminationCmdProductId)

	path = strReplace(path, "{"+"product_id"+"}", escapedProductId, -1)

	escapedSerialNumber := url.PathEscape(GadgetsDisableTerminationCmdSerialNumber)

	path = strReplace(path, "{"+"serial_number"+"}", escapedSerialNumber, -1)

	return path
}

func buildQueryForGadgetsDisableTerminationCmd() url.Values {
	result := url.Values{}

	return result
}
