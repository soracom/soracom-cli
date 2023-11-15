// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// LoraGatewaysDeleteTagCmdGatewayId holds value of 'gateway_id' option
var LoraGatewaysDeleteTagCmdGatewayId string

// LoraGatewaysDeleteTagCmdTagName holds value of 'tag_name' option
var LoraGatewaysDeleteTagCmdTagName string

func InitLoraGatewaysDeleteTagCmd() {
	LoraGatewaysDeleteTagCmd.Flags().StringVar(&LoraGatewaysDeleteTagCmdGatewayId, "gateway-id", "", TRAPI("ID of the target LoRaWAN gateway."))

	LoraGatewaysDeleteTagCmd.Flags().StringVar(&LoraGatewaysDeleteTagCmdTagName, "tag-name", "", TRAPI("Tag name to be deleted. (This will be part of a URL path, so it needs to be percent-encoded. In JavaScript, specify the name after it has been encoded using encodeURIComponent().)"))

	LoraGatewaysDeleteTagCmd.RunE = LoraGatewaysDeleteTagCmdRunE

	LoraGatewaysCmd.AddCommand(LoraGatewaysDeleteTagCmd)
}

// LoraGatewaysDeleteTagCmd defines 'delete-tag' subcommand
var LoraGatewaysDeleteTagCmd = &cobra.Command{
	Use:   "delete-tag",
	Short: TRAPI("/lora_gateways/{gateway_id}/tags/{tag_name}:delete:summary"),
	Long:  TRAPI(`/lora_gateways/{gateway_id}/tags/{tag_name}:delete:description`) + "\n\n" + createLinkToAPIReference("LoraGateway", "deleteLoraGatewayTag"),
}

func LoraGatewaysDeleteTagCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectLoraGatewaysDeleteTagCmdParams(ac)
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

func collectLoraGatewaysDeleteTagCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("gateway_id", "gateway-id", "path", parsedBody, LoraGatewaysDeleteTagCmdGatewayId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("tag_name", "tag-name", "path", parsedBody, LoraGatewaysDeleteTagCmdTagName)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForLoraGatewaysDeleteTagCmd("/lora_gateways/{gateway_id}/tags/{tag_name}"),
		query:  buildQueryForLoraGatewaysDeleteTagCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLoraGatewaysDeleteTagCmd(path string) string {

	escapedGatewayId := url.PathEscape(LoraGatewaysDeleteTagCmdGatewayId)

	path = strReplace(path, "{"+"gateway_id"+"}", escapedGatewayId, -1)

	escapedTagName := url.PathEscape(LoraGatewaysDeleteTagCmdTagName)

	path = strReplace(path, "{"+"tag_name"+"}", escapedTagName, -1)

	return path
}

func buildQueryForLoraGatewaysDeleteTagCmd() url.Values {
	result := url.Values{}

	return result
}
