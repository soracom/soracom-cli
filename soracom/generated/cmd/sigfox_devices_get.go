package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SigfoxDevicesGetCmdDeviceId holds value of 'device_id' option
var SigfoxDevicesGetCmdDeviceId string

func init() {
	SigfoxDevicesGetCmd.Flags().StringVar(&SigfoxDevicesGetCmdDeviceId, "device-id", "", TRAPI("Device ID of the target Sigfox device."))

	SigfoxDevicesCmd.AddCommand(SigfoxDevicesGetCmd)
}

// SigfoxDevicesGetCmd defines 'get' subcommand
var SigfoxDevicesGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/sigfox_devices/{device_id}:get:summary"),
	Long:  TRAPI(`/sigfox_devices/{device_id}:get:description`),
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

		param, err := collectSigfoxDevicesGetCmdParams(ac)
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

func collectSigfoxDevicesGetCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForSigfoxDevicesGetCmd("/sigfox_devices/{device_id}"),
		query:  buildQueryForSigfoxDevicesGetCmd(),
	}, nil
}

func buildPathForSigfoxDevicesGetCmd(path string) string {

	path = strings.Replace(path, "{"+"device_id"+"}", SigfoxDevicesGetCmdDeviceId, -1)

	return path
}

func buildQueryForSigfoxDevicesGetCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
