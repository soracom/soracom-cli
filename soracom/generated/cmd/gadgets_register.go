// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"

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

// GadgetsRegisterCmdProductId holds value of 'product_id' option
var GadgetsRegisterCmdProductId string

// GadgetsRegisterCmdSerialNumber holds value of 'serial_number' option
var GadgetsRegisterCmdSerialNumber string

// GadgetsRegisterCmdBody holds contents of request body to be sent
var GadgetsRegisterCmdBody string

func init() {
	GadgetsRegisterCmd.Flags().StringVar(&GadgetsRegisterCmdProductId, "product-id", "", TRAPI("Product ID of the target gadget."))

	GadgetsRegisterCmd.Flags().StringVar(&GadgetsRegisterCmdSerialNumber, "serial-number", "", TRAPI("Serial Number of the target gadget."))

	GadgetsRegisterCmd.Flags().StringVar(&GadgetsRegisterCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	GadgetsCmd.AddCommand(GadgetsRegisterCmd)
}

// GadgetsRegisterCmd defines 'register' subcommand
var GadgetsRegisterCmd = &cobra.Command{
	Use:   "register",
	Short: TRAPI("/gadgets/{product_id}/{serial_number}/register:post:summary"),
	Long:  TRAPI(`/gadgets/{product_id}/{serial_number}/register:post:description`),
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

		param, err := collectGadgetsRegisterCmdParams(ac)
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

func collectGadgetsRegisterCmdParams(ac *apiClient) (*apiParams, error) {
	body, err := buildBodyForGadgetsRegisterCmd()
	if err != nil {
		return nil, err
	}
	contentType := "application/json"

	if GadgetsRegisterCmdProductId == "" {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "product-id")
		}

	}

	if GadgetsRegisterCmdSerialNumber == "" {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "serial-number")
		}

	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForGadgetsRegisterCmd("/gadgets/{product_id}/{serial_number}/register"),
		query:       buildQueryForGadgetsRegisterCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForGadgetsRegisterCmd(path string) string {

	escapedProductId := url.PathEscape(GadgetsRegisterCmdProductId)

	path = strReplace(path, "{"+"product_id"+"}", escapedProductId, -1)

	escapedSerialNumber := url.PathEscape(GadgetsRegisterCmdSerialNumber)

	path = strReplace(path, "{"+"serial_number"+"}", escapedSerialNumber, -1)

	return path
}

func buildQueryForGadgetsRegisterCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForGadgetsRegisterCmd() (string, error) {
	var result map[string]interface{}

	if GadgetsRegisterCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(GadgetsRegisterCmdBody, "@") {
			fname := strings.TrimPrefix(GadgetsRegisterCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if GadgetsRegisterCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(GadgetsRegisterCmdBody)
		}

		if err != nil {
			return "", err
		}

		err = json.Unmarshal(b, &result)
		if err != nil {
			return "", err
		}
	}

	if result == nil {
		result = make(map[string]interface{})
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
