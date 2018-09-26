package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SigfoxDevicesSendDataCmdData holds value of 'data' option
var SigfoxDevicesSendDataCmdData string

// SigfoxDevicesSendDataCmdDeviceId holds value of 'device_id' option
var SigfoxDevicesSendDataCmdDeviceId string

// SigfoxDevicesSendDataCmdBody holds contents of request body to be sent
var SigfoxDevicesSendDataCmdBody string

func init() {
	SigfoxDevicesSendDataCmd.Flags().StringVar(&SigfoxDevicesSendDataCmdData, "data", "", TRAPI(""))

	SigfoxDevicesSendDataCmd.Flags().StringVar(&SigfoxDevicesSendDataCmdDeviceId, "device-id", "", TRAPI("ID of the recipient device."))

	SigfoxDevicesSendDataCmd.Flags().StringVar(&SigfoxDevicesSendDataCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	SigfoxDevicesCmd.AddCommand(SigfoxDevicesSendDataCmd)
}

// SigfoxDevicesSendDataCmd defines 'send-data' subcommand
var SigfoxDevicesSendDataCmd = &cobra.Command{
	Use:   "send-data",
	Short: TRAPI("/sigfox_devices/{device_id}/data:post:summary"),
	Long:  TRAPI(`/sigfox_devices/{device_id}/data:post:description`),
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

		param, err := collectSigfoxDevicesSendDataCmdParams(ac)
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

func collectSigfoxDevicesSendDataCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForSigfoxDevicesSendDataCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSigfoxDevicesSendDataCmd("/sigfox_devices/{device_id}/data"),
		query:       buildQueryForSigfoxDevicesSendDataCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForSigfoxDevicesSendDataCmd(path string) string {

	path = strings.Replace(path, "{"+"device_id"+"}", SigfoxDevicesSendDataCmdDeviceId, -1)

	return path
}

func buildQueryForSigfoxDevicesSendDataCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForSigfoxDevicesSendDataCmd() (string, error) {
	var result map[string]interface{}

	if SigfoxDevicesSendDataCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SigfoxDevicesSendDataCmdBody, "@") {
			fname := strings.TrimPrefix(SigfoxDevicesSendDataCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if SigfoxDevicesSendDataCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(SigfoxDevicesSendDataCmdBody)
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

	if SigfoxDevicesSendDataCmdData != "" {
		result["data"] = SigfoxDevicesSendDataCmdData
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
