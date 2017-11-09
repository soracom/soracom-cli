package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SigfoxDevicesDeleteTagCmdDeviceId holds value of 'device_id' option
var SigfoxDevicesDeleteTagCmdDeviceId string

// SigfoxDevicesDeleteTagCmdTagName holds value of 'tag_name' option
var SigfoxDevicesDeleteTagCmdTagName string

func init() {
	SigfoxDevicesDeleteTagCmd.Flags().StringVar(&SigfoxDevicesDeleteTagCmdDeviceId, "device-id", "", TRAPI("device ID of the target Sigfox device."))

	SigfoxDevicesDeleteTagCmd.Flags().StringVar(&SigfoxDevicesDeleteTagCmdTagName, "tag-name", "", TRAPI("Tag name to be deleted. (This will be part of a URL path, so it needs to be percent-encoded. In JavaScript, specify the name after it has been encoded using encodeURIComponent().)"))

	SigfoxDevicesCmd.AddCommand(SigfoxDevicesDeleteTagCmd)
}

// SigfoxDevicesDeleteTagCmd defines 'delete-tag' subcommand
var SigfoxDevicesDeleteTagCmd = &cobra.Command{
	Use:   "delete-tag",
	Short: TRAPI("/sigfox_devices/{device_id}/tags/{tag_name}:delete:summary"),
	Long:  TRAPI(`/sigfox_devices/{device_id}/tags/{tag_name}:delete:description`),
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

		param, err := collectSigfoxDevicesDeleteTagCmdParams()
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

func collectSigfoxDevicesDeleteTagCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "DELETE",
		path:   buildPathForSigfoxDevicesDeleteTagCmd("/sigfox_devices/{device_id}/tags/{tag_name}"),
		query:  buildQueryForSigfoxDevicesDeleteTagCmd(),
	}, nil
}

func buildPathForSigfoxDevicesDeleteTagCmd(path string) string {

	path = strings.Replace(path, "{"+"device_id"+"}", SigfoxDevicesDeleteTagCmdDeviceId, -1)

	path = strings.Replace(path, "{"+"tag_name"+"}", SigfoxDevicesDeleteTagCmdTagName, -1)

	return path
}

func buildQueryForSigfoxDevicesDeleteTagCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
