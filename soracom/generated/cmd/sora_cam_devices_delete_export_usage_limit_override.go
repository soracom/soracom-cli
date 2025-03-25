// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SoraCamDevicesDeleteExportUsageLimitOverrideCmdDeviceId holds value of 'device_id' option
var SoraCamDevicesDeleteExportUsageLimitOverrideCmdDeviceId string

func InitSoraCamDevicesDeleteExportUsageLimitOverrideCmd() {
	SoraCamDevicesDeleteExportUsageLimitOverrideCmd.Flags().StringVar(&SoraCamDevicesDeleteExportUsageLimitOverrideCmdDeviceId, "device-id", "", TRAPI("Device ID of the target SoraCam compatible camera device."))

	SoraCamDevicesDeleteExportUsageLimitOverrideCmd.RunE = SoraCamDevicesDeleteExportUsageLimitOverrideCmdRunE

	SoraCamDevicesCmd.AddCommand(SoraCamDevicesDeleteExportUsageLimitOverrideCmd)
}

// SoraCamDevicesDeleteExportUsageLimitOverrideCmd defines 'delete-export-usage-limit-override' subcommand
var SoraCamDevicesDeleteExportUsageLimitOverrideCmd = &cobra.Command{
	Use:   "delete-export-usage-limit-override",
	Short: TRAPI("/sora_cam/devices/{device_id}/exports/usage/limit_override:delete:summary"),
	Long:  TRAPI(`/sora_cam/devices/{device_id}/exports/usage/limit_override:delete:description`) + "\n\n" + createLinkToAPIReference("SoraCam", "deleteSoraCamDeviceExportUsageLimitOverride"),
}

func SoraCamDevicesDeleteExportUsageLimitOverrideCmdRunE(cmd *cobra.Command, args []string) error {

	if len(args) > 0 {
		return fmt.Errorf("unexpected arguments passed => %v", args)
	}

	opt := &apiClientOptions{
		BasePath: "/v1",
		Language: getSelectedLanguage(),
	}

	ac := newAPIClient(opt)
	if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
		ac.SetVerbose(true)
	}
	err := ac.getAPICredentials()
	if err != nil {
		cmd.SilenceUsage = true
		return err
	}

	param, err := collectSoraCamDevicesDeleteExportUsageLimitOverrideCmdParams(ac)
	if err != nil {
		return err
	}

	body, err := ac.callAPI(param)
	if err != nil {
		cmd.SilenceUsage = true
		return err
	}

	if body == "" {
		return nil
	}

	if rawOutput {
		_, err = os.Stdout.Write([]byte(body))
	} else {
		return prettyPrintStringAsJSON(body)
	}
	return err
}

func collectSoraCamDevicesDeleteExportUsageLimitOverrideCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, SoraCamDevicesDeleteExportUsageLimitOverrideCmdDeviceId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForSoraCamDevicesDeleteExportUsageLimitOverrideCmd("/sora_cam/devices/{device_id}/exports/usage/limit_override"),
		query:  buildQueryForSoraCamDevicesDeleteExportUsageLimitOverrideCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSoraCamDevicesDeleteExportUsageLimitOverrideCmd(path string) string {

	escapedDeviceId := url.PathEscape(SoraCamDevicesDeleteExportUsageLimitOverrideCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForSoraCamDevicesDeleteExportUsageLimitOverrideCmd() url.Values {
	result := url.Values{}

	return result
}
