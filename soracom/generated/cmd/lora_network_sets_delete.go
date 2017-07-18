package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LoraNetworkSetsDeleteCmdNsId holds value of 'ns_id' option
var LoraNetworkSetsDeleteCmdNsId string

func init() {
	LoraNetworkSetsDeleteCmd.Flags().StringVar(&LoraNetworkSetsDeleteCmdNsId, "ns-id", "", TRAPI("ID of the target LoRa network set."))

	LoraNetworkSetsCmd.AddCommand(LoraNetworkSetsDeleteCmd)
}

// LoraNetworkSetsDeleteCmd defines 'delete' subcommand
var LoraNetworkSetsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: TRAPI("/lora_network_sets/{ns_id}:delete:summary"),
	Long:  TRAPI(`/lora_network_sets/{ns_id}:delete:description`),
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

		param, err := collectLoraNetworkSetsDeleteCmdParams()
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

func collectLoraNetworkSetsDeleteCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "DELETE",
		path:   buildPathForLoraNetworkSetsDeleteCmd("/lora_network_sets/{ns_id}"),
		query:  buildQueryForLoraNetworkSetsDeleteCmd(),
	}, nil
}

func buildPathForLoraNetworkSetsDeleteCmd(path string) string {

	path = strings.Replace(path, "{"+"ns_id"+"}", LoraNetworkSetsDeleteCmdNsId, -1)

	return path
}

func buildQueryForLoraNetworkSetsDeleteCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
