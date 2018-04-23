package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LoraNetworkSetsGetCmdNsId holds value of 'ns_id' option
var LoraNetworkSetsGetCmdNsId string

func init() {
	LoraNetworkSetsGetCmd.Flags().StringVar(&LoraNetworkSetsGetCmdNsId, "ns-id", "", TRAPI("ID of the target LoRa network set."))

	LoraNetworkSetsCmd.AddCommand(LoraNetworkSetsGetCmd)
}

// LoraNetworkSetsGetCmd defines 'get' subcommand
var LoraNetworkSetsGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/lora_network_sets/{ns_id}:get:summary"),
	Long:  TRAPI(`/lora_network_sets/{ns_id}:get:description`),
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

		param, err := collectLoraNetworkSetsGetCmdParams(ac)
		if err != nil {
			return err
		}

		result, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if result == "" {
			return nil
		}

		return prettyPrintStringAsJSON(result)
	},
}

func collectLoraNetworkSetsGetCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForLoraNetworkSetsGetCmd("/lora_network_sets/{ns_id}"),
		query:  buildQueryForLoraNetworkSetsGetCmd(),
	}, nil
}

func buildPathForLoraNetworkSetsGetCmd(path string) string {

	path = strings.Replace(path, "{"+"ns_id"+"}", LoraNetworkSetsGetCmdNsId, -1)

	return path
}

func buildQueryForLoraNetworkSetsGetCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
