// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// VolumeDiscountsGetCmdContractId holds value of 'contract_id' option
var VolumeDiscountsGetCmdContractId string

func InitVolumeDiscountsGetCmd() {
	VolumeDiscountsGetCmd.Flags().StringVar(&VolumeDiscountsGetCmdContractId, "contract-id", "", TRAPI("Contract ID."))

	VolumeDiscountsGetCmd.RunE = VolumeDiscountsGetCmdRunE

	VolumeDiscountsCmd.AddCommand(VolumeDiscountsGetCmd)
}

// VolumeDiscountsGetCmd defines 'get' subcommand
var VolumeDiscountsGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/volume_discounts/{contract_id}:get:summary"),
	Long:  TRAPI(`/volume_discounts/{contract_id}:get:description`) + "\n\n" + createLinkToAPIReference("Payment", "getVolumeDiscount"),
}

func VolumeDiscountsGetCmdRunE(cmd *cobra.Command, args []string) error {

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
	err := ac.getAPICredentials()
	if err != nil {
		cmd.SilenceUsage = true
		return err
	}

	param, err := collectVolumeDiscountsGetCmdParams(ac)
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

func collectVolumeDiscountsGetCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("contract_id", "contract-id", "path", parsedBody, VolumeDiscountsGetCmdContractId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForVolumeDiscountsGetCmd("/volume_discounts/{contract_id}"),
		query:  buildQueryForVolumeDiscountsGetCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForVolumeDiscountsGetCmd(path string) string {

	escapedContractId := url.PathEscape(VolumeDiscountsGetCmdContractId)

	path = strReplace(path, "{"+"contract_id"+"}", escapedContractId, -1)

	return path
}

func buildQueryForVolumeDiscountsGetCmd() url.Values {
	result := url.Values{}

	return result
}
