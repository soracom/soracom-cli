package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LoraNetworkSetsDeleteTagCmdNsId holds value of 'ns_id' option
var LoraNetworkSetsDeleteTagCmdNsId string

// LoraNetworkSetsDeleteTagCmdTagName holds value of 'tag_name' option
var LoraNetworkSetsDeleteTagCmdTagName string

func init() {
	LoraNetworkSetsDeleteTagCmd.Flags().StringVar(&LoraNetworkSetsDeleteTagCmdNsId, "ns-id", "", TRAPI("ID of the target LoRa network set."))

	LoraNetworkSetsDeleteTagCmd.Flags().StringVar(&LoraNetworkSetsDeleteTagCmdTagName, "tag-name", "", TRAPI("Name of tag to be deleted. (This will be part of a URL path, so it needs to be percent-encoded. In JavaScript, specify the name after it has been encoded using encodeURIComponent().)"))

	LoraNetworkSetsCmd.AddCommand(LoraNetworkSetsDeleteTagCmd)
}

// LoraNetworkSetsDeleteTagCmd defines 'delete-tag' subcommand
var LoraNetworkSetsDeleteTagCmd = &cobra.Command{
	Use:   "delete-tag",
	Short: TRAPI("/lora_network_sets/{ns_id}/tags/{tag_name}:delete:summary"),
	Long:  TRAPI(`/lora_network_sets/{ns_id}/tags/{tag_name}:delete:description`),
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

		param, err := collectLoraNetworkSetsDeleteTagCmdParams(ac)
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

func collectLoraNetworkSetsDeleteTagCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "DELETE",
		path:   buildPathForLoraNetworkSetsDeleteTagCmd("/lora_network_sets/{ns_id}/tags/{tag_name}"),
		query:  buildQueryForLoraNetworkSetsDeleteTagCmd(),
	}, nil
}

func buildPathForLoraNetworkSetsDeleteTagCmd(path string) string {

	path = strings.Replace(path, "{"+"ns_id"+"}", LoraNetworkSetsDeleteTagCmdNsId, -1)

	path = strings.Replace(path, "{"+"tag_name"+"}", LoraNetworkSetsDeleteTagCmdTagName, -1)

	return path
}

func buildQueryForLoraNetworkSetsDeleteTagCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
