// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

func InitVolumeDiscountsListCmd() {

	VolumeDiscountsListCmd.RunE = VolumeDiscountsListCmdRunE

	VolumeDiscountsCmd.AddCommand(VolumeDiscountsListCmd)
}

// VolumeDiscountsListCmd defines 'list' subcommand
var VolumeDiscountsListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/volume_discounts:get:summary"),
	Long:  TRAPI(`/volume_discounts:get:description`) + "\n\n" + createLinkToAPIReference("Payment", "listVolumeDiscounts"),
}

func VolumeDiscountsListCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectVolumeDiscountsListCmdParams(ac)
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

func collectVolumeDiscountsListCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForVolumeDiscountsListCmd("/volume_discounts"),
		query:  buildQueryForVolumeDiscountsListCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForVolumeDiscountsListCmd(path string) string {

	return path
}

func buildQueryForVolumeDiscountsListCmd() url.Values {
	result := url.Values{}

	return result
}
