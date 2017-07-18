package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LoraGatewaysPutTagsCmdGatewayId holds value of 'gateway_id' option
var LoraGatewaysPutTagsCmdGatewayId string

// LoraGatewaysPutTagsCmdBody holds contents of request body to be sent
var LoraGatewaysPutTagsCmdBody string

func init() {
	LoraGatewaysPutTagsCmd.Flags().StringVar(&LoraGatewaysPutTagsCmdGatewayId, "gateway-id", "", TRAPI("ID of the target LoRa gateway."))

	LoraGatewaysPutTagsCmd.Flags().StringVar(&LoraGatewaysPutTagsCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	LoraGatewaysCmd.AddCommand(LoraGatewaysPutTagsCmd)
}

// LoraGatewaysPutTagsCmd defines 'put-tags' subcommand
var LoraGatewaysPutTagsCmd = &cobra.Command{
	Use:   "put-tags",
	Short: TRAPI("/lora_gateways/{gateway_id}/tags:put:summary"),
	Long:  TRAPI(`/lora_gateways/{gateway_id}/tags:put:description`),
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

		param, err := collectLoraGatewaysPutTagsCmdParams()
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

func collectLoraGatewaysPutTagsCmdParams() (*apiParams, error) {

	body, err := buildBodyForLoraGatewaysPutTagsCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForLoraGatewaysPutTagsCmd("/lora_gateways/{gateway_id}/tags"),
		query:       buildQueryForLoraGatewaysPutTagsCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForLoraGatewaysPutTagsCmd(path string) string {

	path = strings.Replace(path, "{"+"gateway_id"+"}", LoraGatewaysPutTagsCmdGatewayId, -1)

	return path
}

func buildQueryForLoraGatewaysPutTagsCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForLoraGatewaysPutTagsCmd() (string, error) {
	if LoraGatewaysPutTagsCmdBody != "" {
		if strings.HasPrefix(LoraGatewaysPutTagsCmdBody, "@") {
			fname := strings.TrimPrefix(LoraGatewaysPutTagsCmdBody, "@")
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if LoraGatewaysPutTagsCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return LoraGatewaysPutTagsCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
