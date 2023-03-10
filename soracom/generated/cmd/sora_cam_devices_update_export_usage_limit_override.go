// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"

	"strings"

	"github.com/spf13/cobra"
)

// SoraCamDevicesUpdateExportUsageLimitOverrideCmdDeviceId holds value of 'device_id' option
var SoraCamDevicesUpdateExportUsageLimitOverrideCmdDeviceId string

// SoraCamDevicesUpdateExportUsageLimitOverrideCmdLimitHours holds value of 'limitHours' option
var SoraCamDevicesUpdateExportUsageLimitOverrideCmdLimitHours int64

// SoraCamDevicesUpdateExportUsageLimitOverrideCmdBody holds contents of request body to be sent
var SoraCamDevicesUpdateExportUsageLimitOverrideCmdBody string

func init() {
	SoraCamDevicesUpdateExportUsageLimitOverrideCmd.Flags().StringVar(&SoraCamDevicesUpdateExportUsageLimitOverrideCmdDeviceId, "device-id", "", TRAPI("Device ID of the target compatible camera device."))

	SoraCamDevicesUpdateExportUsageLimitOverrideCmd.Flags().Int64Var(&SoraCamDevicesUpdateExportUsageLimitOverrideCmdLimitHours, "limit-hours", 0, TRAPI("New value for the monthly limit for the amount of hours that recorded video can be exported."))

	SoraCamDevicesUpdateExportUsageLimitOverrideCmd.Flags().StringVar(&SoraCamDevicesUpdateExportUsageLimitOverrideCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	SoraCamDevicesCmd.AddCommand(SoraCamDevicesUpdateExportUsageLimitOverrideCmd)
}

// SoraCamDevicesUpdateExportUsageLimitOverrideCmd defines 'update-export-usage-limit-override' subcommand
var SoraCamDevicesUpdateExportUsageLimitOverrideCmd = &cobra.Command{
	Use:   "update-export-usage-limit-override",
	Short: TRAPI("/sora_cam/devices/{device_id}/exports/usage/limit_override:put:summary"),
	Long:  TRAPI(`/sora_cam/devices/{device_id}/exports/usage/limit_override:put:description`) + "\n\n" + createLinkToAPIReference("SoraCam", "updateSoraCamDeviceExportUsageLimitOverride"),
	RunE: func(cmd *cobra.Command, args []string) error {

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
		err := authHelper(ac, cmd, args)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		param, err := collectSoraCamDevicesUpdateExportUsageLimitOverrideCmdParams(ac)
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
	},
}

func collectSoraCamDevicesUpdateExportUsageLimitOverrideCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForSoraCamDevicesUpdateExportUsageLimitOverrideCmd()
	if err != nil {
		return nil, err
	}
	contentType := "application/json"

	if contentType == "application/json" {
		err = json.Unmarshal([]byte(body), &parsedBody)
		if err != nil {
			return nil, fmt.Errorf("invalid json format specified for `--body` parameter: %s", err)
		}
	}

	err = checkIfRequiredStringParameterIsSupplied("device_id", "device-id", "path", parsedBody, SoraCamDevicesUpdateExportUsageLimitOverrideCmdDeviceId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForSoraCamDevicesUpdateExportUsageLimitOverrideCmd("/sora_cam/devices/{device_id}/exports/usage/limit_override"),
		query:       buildQueryForSoraCamDevicesUpdateExportUsageLimitOverrideCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSoraCamDevicesUpdateExportUsageLimitOverrideCmd(path string) string {

	escapedDeviceId := url.PathEscape(SoraCamDevicesUpdateExportUsageLimitOverrideCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForSoraCamDevicesUpdateExportUsageLimitOverrideCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSoraCamDevicesUpdateExportUsageLimitOverrideCmd() (string, error) {
	var result map[string]interface{}

	if SoraCamDevicesUpdateExportUsageLimitOverrideCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SoraCamDevicesUpdateExportUsageLimitOverrideCmdBody, "@") {
			fname := strings.TrimPrefix(SoraCamDevicesUpdateExportUsageLimitOverrideCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if SoraCamDevicesUpdateExportUsageLimitOverrideCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(SoraCamDevicesUpdateExportUsageLimitOverrideCmdBody)
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

	if SoraCamDevicesUpdateExportUsageLimitOverrideCmdLimitHours != 0 {
		result["limitHours"] = SoraCamDevicesUpdateExportUsageLimitOverrideCmdLimitHours
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
