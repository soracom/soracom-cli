package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// DevicesDeleteCmdDeviceId holds value of 'device_id' option
var DevicesDeleteCmdDeviceId string

func init() {
	DevicesDeleteCmd.Flags().StringVar(&DevicesDeleteCmdDeviceId, "device-id", "", TRAPI("Device to delete"))

	DevicesCmd.AddCommand(DevicesDeleteCmd)
}

// DevicesDeleteCmd defines 'delete' subcommand
var DevicesDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: TRAPI("/devices/{device_id}:delete:summary"),
	Long:  TRAPI(`/devices/{device_id}:delete:description`),
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

		param, err := collectDevicesDeleteCmdParams(ac)
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

func collectDevicesDeleteCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "DELETE",
		path:   buildPathForDevicesDeleteCmd("/devices/{device_id}"),
		query:  buildQueryForDevicesDeleteCmd(),
	}, nil
}

func buildPathForDevicesDeleteCmd(path string) string {

	path = strings.Replace(path, "{"+"device_id"+"}", DevicesDeleteCmdDeviceId, -1)

	return path
}

func buildQueryForDevicesDeleteCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
