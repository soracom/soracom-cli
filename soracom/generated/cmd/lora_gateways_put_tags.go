// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"io/ioutil"

	"mime"
	"net/url"
	"os"

	"strings"

	"github.com/itchyny/gojq"
	"github.com/soracom/soracom-cli/generators/lib"
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

		param, err := collectLoraGatewaysPutTagsCmdParams(ac)
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

func collectLoraGatewaysPutTagsCmdParams(ac *apiClient) (*apiParams, error) {
	body, err := buildBodyForLoraGatewaysPutTagsCmd()
	if err != nil {
		return nil, err
	}
	contentType := "application/json"

	if LoraGatewaysPutTagsCmdGatewayId == "" {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "gateway-id")
		}

	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForLoraGatewaysPutTagsCmd("/lora_gateways/{gateway_id}/tags"),
		query:       buildQueryForLoraGatewaysPutTagsCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLoraGatewaysPutTagsCmd(path string) string {

	escapedGatewayId := url.PathEscape(LoraGatewaysPutTagsCmdGatewayId)

	path = strReplace(path, "{"+"gateway_id"+"}", escapedGatewayId, -1)

	return path
}

func buildQueryForLoraGatewaysPutTagsCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForLoraGatewaysPutTagsCmd() (string, error) {
	var b []byte
	var err error

	if LoraGatewaysPutTagsCmdBody != "" {
		if strings.HasPrefix(LoraGatewaysPutTagsCmdBody, "@") {
			fname := strings.TrimPrefix(LoraGatewaysPutTagsCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if LoraGatewaysPutTagsCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(LoraGatewaysPutTagsCmdBody)
		}

		if err != nil {
			return "", err
		}
	}

	if b == nil {
		b = []byte{}
	}

	return string(b), nil
}
