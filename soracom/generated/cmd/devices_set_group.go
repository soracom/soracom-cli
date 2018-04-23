package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// DevicesSetGroupCmdDeviceId holds value of 'deviceId' option
var DevicesSetGroupCmdDeviceId string

// DevicesSetGroupCmdBody holds contents of request body to be sent
var DevicesSetGroupCmdBody string

func init() {
	DevicesSetGroupCmd.Flags().StringVar(&DevicesSetGroupCmdDeviceId, "device-id", "", TRAPI("Device to update"))

	DevicesSetGroupCmd.Flags().StringVar(&DevicesSetGroupCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	DevicesCmd.AddCommand(DevicesSetGroupCmd)
}

// DevicesSetGroupCmd defines 'set-group' subcommand
var DevicesSetGroupCmd = &cobra.Command{
	Use:   "set-group",
	Short: TRAPI("/devices/{deviceId}/set_group:post:summary"),
	Long:  TRAPI(`/devices/{deviceId}/set_group:post:description`),
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

		param, err := collectDevicesSetGroupCmdParams(ac)
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

func collectDevicesSetGroupCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForDevicesSetGroupCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForDevicesSetGroupCmd("/devices/{deviceId}/set_group"),
		query:       buildQueryForDevicesSetGroupCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForDevicesSetGroupCmd(path string) string {

	path = strings.Replace(path, "{"+"deviceId"+"}", DevicesSetGroupCmdDeviceId, -1)

	return path
}

func buildQueryForDevicesSetGroupCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForDevicesSetGroupCmd() (string, error) {
	if DevicesSetGroupCmdBody != "" {
		if strings.HasPrefix(DevicesSetGroupCmdBody, "@") {
			fname := strings.TrimPrefix(DevicesSetGroupCmdBody, "@")
			// #nosec
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if DevicesSetGroupCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return DevicesSetGroupCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
