package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SigfoxDevicesTerminateCmdDeviceId holds value of 'device_id' option
var SigfoxDevicesTerminateCmdDeviceId string

func init() {
	SigfoxDevicesTerminateCmd.Flags().StringVar(&SigfoxDevicesTerminateCmdDeviceId, "device-id", "", TRAPI("Device ID of the target Sigfox device."))

	SigfoxDevicesCmd.AddCommand(SigfoxDevicesTerminateCmd)
}

// SigfoxDevicesTerminateCmd defines 'terminate' subcommand
var SigfoxDevicesTerminateCmd = &cobra.Command{
	Use:   "terminate",
	Short: TRAPI("/sigfox_devices/{device_id}/terminate:post:summary"),
	Long:  TRAPI(`/sigfox_devices/{device_id}/terminate:post:description`),
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

		param, err := collectSigfoxDevicesTerminateCmdParams(ac)
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

func collectSigfoxDevicesTerminateCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForSigfoxDevicesTerminateCmd("/sigfox_devices/{device_id}/terminate"),
		query:  buildQueryForSigfoxDevicesTerminateCmd(),
	}, nil
}

func buildPathForSigfoxDevicesTerminateCmd(path string) string {

	path = strings.Replace(path, "{"+"device_id"+"}", SigfoxDevicesTerminateCmdDeviceId, -1)

	return path
}

func buildQueryForSigfoxDevicesTerminateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
