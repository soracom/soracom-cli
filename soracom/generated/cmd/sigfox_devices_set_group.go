package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SigfoxDevicesSetGroupCmdDeviceId holds value of 'device_id' option
var SigfoxDevicesSetGroupCmdDeviceId string

// SigfoxDevicesSetGroupCmdGroupId holds value of 'groupId' option
var SigfoxDevicesSetGroupCmdGroupId string

// SigfoxDevicesSetGroupCmdOperatorId holds value of 'operatorId' option
var SigfoxDevicesSetGroupCmdOperatorId string

// SigfoxDevicesSetGroupCmdCreatedTime holds value of 'createdTime' option
var SigfoxDevicesSetGroupCmdCreatedTime int64

// SigfoxDevicesSetGroupCmdLastModifiedTime holds value of 'lastModifiedTime' option
var SigfoxDevicesSetGroupCmdLastModifiedTime int64

// SigfoxDevicesSetGroupCmdBody holds contents of request body to be sent
var SigfoxDevicesSetGroupCmdBody string

func init() {
	SigfoxDevicesSetGroupCmd.Flags().StringVar(&SigfoxDevicesSetGroupCmdDeviceId, "device-id", "", TRAPI("Device ID of the target Sigfox device."))

	SigfoxDevicesSetGroupCmd.Flags().StringVar(&SigfoxDevicesSetGroupCmdGroupId, "group-id", "", TRAPI(""))

	SigfoxDevicesSetGroupCmd.Flags().StringVar(&SigfoxDevicesSetGroupCmdOperatorId, "operator-id", "", TRAPI(""))

	SigfoxDevicesSetGroupCmd.Flags().Int64Var(&SigfoxDevicesSetGroupCmdCreatedTime, "created-time", 0, TRAPI(""))

	SigfoxDevicesSetGroupCmd.Flags().Int64Var(&SigfoxDevicesSetGroupCmdLastModifiedTime, "last-modified-time", 0, TRAPI(""))

	SigfoxDevicesSetGroupCmd.Flags().StringVar(&SigfoxDevicesSetGroupCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	SigfoxDevicesCmd.AddCommand(SigfoxDevicesSetGroupCmd)
}

// SigfoxDevicesSetGroupCmd defines 'set-group' subcommand
var SigfoxDevicesSetGroupCmd = &cobra.Command{
	Use:   "set-group",
	Short: TRAPI("/sigfox_devices/{device_id}/set_group:post:summary"),
	Long:  TRAPI(`/sigfox_devices/{device_id}/set_group:post:description`),
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

		param, err := collectSigfoxDevicesSetGroupCmdParams()
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

func collectSigfoxDevicesSetGroupCmdParams() (*apiParams, error) {

	body, err := buildBodyForSigfoxDevicesSetGroupCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSigfoxDevicesSetGroupCmd("/sigfox_devices/{device_id}/set_group"),
		query:       buildQueryForSigfoxDevicesSetGroupCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForSigfoxDevicesSetGroupCmd(path string) string {

	path = strings.Replace(path, "{"+"device_id"+"}", SigfoxDevicesSetGroupCmdDeviceId, -1)

	return path
}

func buildQueryForSigfoxDevicesSetGroupCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForSigfoxDevicesSetGroupCmd() (string, error) {
	if SigfoxDevicesSetGroupCmdBody != "" {
		if strings.HasPrefix(SigfoxDevicesSetGroupCmdBody, "@") {
			fname := strings.TrimPrefix(SigfoxDevicesSetGroupCmdBody, "@")
			// #nosec
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if SigfoxDevicesSetGroupCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return SigfoxDevicesSetGroupCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if SigfoxDevicesSetGroupCmdGroupId != "" {
		result["groupId"] = SigfoxDevicesSetGroupCmdGroupId
	}

	if SigfoxDevicesSetGroupCmdOperatorId != "" {
		result["operatorId"] = SigfoxDevicesSetGroupCmdOperatorId
	}

	if SigfoxDevicesSetGroupCmdCreatedTime != 0 {
		result["createdTime"] = SigfoxDevicesSetGroupCmdCreatedTime
	}

	if SigfoxDevicesSetGroupCmdLastModifiedTime != 0 {
		result["lastModifiedTime"] = SigfoxDevicesSetGroupCmdLastModifiedTime
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
