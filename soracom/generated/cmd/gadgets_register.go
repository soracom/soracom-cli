package cmd

import (
	"encoding/json"

	"io/ioutil"

	"os"
	"strings"

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

		_, body, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		return prettyPrintStringAsJSON(body)
	},
}

func collectGadgetsRegisterCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForGadgetsRegisterCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForGadgetsRegisterCmd("/gadgets/{product_id}/{serial_number}/register"),
		query:       buildQueryForGadgetsRegisterCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForGadgetsRegisterCmd(path string) string {

	path = strings.Replace(path, "{"+"product_id"+"}", GadgetsRegisterCmdProductId, -1)

	path = strings.Replace(path, "{"+"serial_number"+"}", GadgetsRegisterCmdSerialNumber, -1)

	return path
}

func buildQueryForGadgetsRegisterCmd() string {
	result := []string{}

	return strings.Join(result, "&")
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
