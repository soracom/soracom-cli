// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SoraCamDevicesAtomCamListFirmwareUpdatesCmdOutputJSONL indicates to output with jsonl format
var SoraCamDevicesAtomCamListFirmwareUpdatesCmdOutputJSONL bool

func InitSoraCamDevicesAtomCamListFirmwareUpdatesCmd() {
	SoraCamDevicesAtomCamListFirmwareUpdatesCmd.Flags().BoolVar(&SoraCamDevicesAtomCamListFirmwareUpdatesCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))

	SoraCamDevicesAtomCamListFirmwareUpdatesCmd.RunE = SoraCamDevicesAtomCamListFirmwareUpdatesCmdRunE

	SoraCamDevicesAtomCamCmd.AddCommand(SoraCamDevicesAtomCamListFirmwareUpdatesCmd)
}

// SoraCamDevicesAtomCamListFirmwareUpdatesCmd defines 'list-firmware-updates' subcommand
var SoraCamDevicesAtomCamListFirmwareUpdatesCmd = &cobra.Command{
	Use:   "list-firmware-updates",
	Short: TRAPI("/sora_cam/devices/atom_cam/firmware/updates:get:summary"),
	Long:  TRAPI(`/sora_cam/devices/atom_cam/firmware/updates:get:description`) + "\n\n" + createLinkToAPIReference("SoraCam", "listSoraCamDeviceAtomCamFirmwareUpdates"),
}

func SoraCamDevicesAtomCamListFirmwareUpdatesCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSoraCamDevicesAtomCamListFirmwareUpdatesCmdParams(ac)
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
		if SoraCamDevicesAtomCamListFirmwareUpdatesCmdOutputJSONL {
			return printStringAsJSONL(body)
		}

		return prettyPrintStringAsJSON(body)
	}
	return err
}

func collectSoraCamDevicesAtomCamListFirmwareUpdatesCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForSoraCamDevicesAtomCamListFirmwareUpdatesCmd("/sora_cam/devices/atom_cam/firmware/updates"),
		query:  buildQueryForSoraCamDevicesAtomCamListFirmwareUpdatesCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSoraCamDevicesAtomCamListFirmwareUpdatesCmd(path string) string {

	return path
}

func buildQueryForSoraCamDevicesAtomCamListFirmwareUpdatesCmd() url.Values {
	result := url.Values{}

	return result
}