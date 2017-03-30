package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LoraGatewaysDeleteTagCmdGatewayId holds value of 'gateway_id' option
var LoraGatewaysDeleteTagCmdGatewayId string

// LoraGatewaysDeleteTagCmdTagName holds value of 'tag_name' option
var LoraGatewaysDeleteTagCmdTagName string

func init() {
	LoraGatewaysDeleteTagCmd.Flags().StringVar(&LoraGatewaysDeleteTagCmdGatewayId, "gateway-id", "", TR("lora_gateways.delete_tag.parameters.gateway_id.description"))

	LoraGatewaysDeleteTagCmd.Flags().StringVar(&LoraGatewaysDeleteTagCmdTagName, "tag-name", "", TR("lora_gateways.delete_tag.parameters.tag_name.description"))

	LoraGatewaysCmd.AddCommand(LoraGatewaysDeleteTagCmd)
}

// LoraGatewaysDeleteTagCmd defines 'delete-tag' subcommand
var LoraGatewaysDeleteTagCmd = &cobra.Command{
	Use:   "delete-tag",
	Short: TR("lora_gateways.delete_tag.summary"),
	Long:  TR(`lora_gateways.delete_tag.description`),
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

		param, err := collectLoraGatewaysDeleteTagCmdParams()
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

func collectLoraGatewaysDeleteTagCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "DELETE",
		path:   buildPathForLoraGatewaysDeleteTagCmd("/lora_gateways/{gateway_id}/tags/{tag_name}"),
		query:  buildQueryForLoraGatewaysDeleteTagCmd(),
	}, nil
}

func buildPathForLoraGatewaysDeleteTagCmd(path string) string {

	path = strings.Replace(path, "{"+"gateway_id"+"}", LoraGatewaysDeleteTagCmdGatewayId, -1)

	path = strings.Replace(path, "{"+"tag_name"+"}", LoraGatewaysDeleteTagCmdTagName, -1)

	return path
}

func buildQueryForLoraGatewaysDeleteTagCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
