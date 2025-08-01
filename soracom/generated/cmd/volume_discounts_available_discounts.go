// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

func InitVolumeDiscountsAvailableDiscountsCmd() {

	VolumeDiscountsAvailableDiscountsCmd.RunE = VolumeDiscountsAvailableDiscountsCmdRunE

	VolumeDiscountsCmd.AddCommand(VolumeDiscountsAvailableDiscountsCmd)
}

// VolumeDiscountsAvailableDiscountsCmd defines 'available-discounts' subcommand
var VolumeDiscountsAvailableDiscountsCmd = &cobra.Command{
	Use:   "available-discounts",
	Short: TRAPI("/volume_discounts/available_discounts:get:summary"),
	Long:  TRAPI(`/volume_discounts/available_discounts:get:description`) + "\n\n" + createLinkToAPIReference("Order", "listAvailableDiscounts"),
}

func VolumeDiscountsAvailableDiscountsCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectVolumeDiscountsAvailableDiscountsCmdParams(ac)
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

func collectVolumeDiscountsAvailableDiscountsCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForVolumeDiscountsAvailableDiscountsCmd("/volume_discounts/available_discounts"),
		query:  buildQueryForVolumeDiscountsAvailableDiscountsCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForVolumeDiscountsAvailableDiscountsCmd(path string) string {

	return path
}

func buildQueryForVolumeDiscountsAvailableDiscountsCmd() url.Values {
	result := url.Values{}

	return result
}
