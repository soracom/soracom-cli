package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// DevicesCreateCmdDeviceId holds value of 'deviceId' option
var DevicesCreateCmdDeviceId string

// DevicesCreateCmdEndpoint holds value of 'endpoint' option
var DevicesCreateCmdEndpoint string

// DevicesCreateCmdFirmwareVersion holds value of 'firmwareVersion' option
var DevicesCreateCmdFirmwareVersion string

// DevicesCreateCmdGroupId holds value of 'groupId' option
var DevicesCreateCmdGroupId string

// DevicesCreateCmdImei holds value of 'imei' option
var DevicesCreateCmdImei string

// DevicesCreateCmdImsi holds value of 'imsi' option
var DevicesCreateCmdImsi string

// DevicesCreateCmdIpAddress holds value of 'ipAddress' option
var DevicesCreateCmdIpAddress string

// DevicesCreateCmdLastModifiedTime holds value of 'lastModifiedTime' option
var DevicesCreateCmdLastModifiedTime string

// DevicesCreateCmdLastRegistrationUpdate holds value of 'lastRegistrationUpdate' option
var DevicesCreateCmdLastRegistrationUpdate string

// DevicesCreateCmdManufacturer holds value of 'manufacturer' option
var DevicesCreateCmdManufacturer string

// DevicesCreateCmdModelNumber holds value of 'modelNumber' option
var DevicesCreateCmdModelNumber string

// DevicesCreateCmdOperatorId holds value of 'operatorId' option
var DevicesCreateCmdOperatorId string

// DevicesCreateCmdRegistrationId holds value of 'registrationId' option
var DevicesCreateCmdRegistrationId string

// DevicesCreateCmdSerialNumber holds value of 'serialNumber' option
var DevicesCreateCmdSerialNumber string

// DevicesCreateCmdRegistrationLifeTime holds value of 'registrationLifeTime' option
var DevicesCreateCmdRegistrationLifeTime int64

// DevicesCreateCmdOnline holds value of 'online' option
var DevicesCreateCmdOnline bool

// DevicesCreateCmdBody holds contents of request body to be sent
var DevicesCreateCmdBody string

func init() {
	DevicesCreateCmd.Flags().StringVar(&DevicesCreateCmdDeviceId, "device-id", "", TRAPI(""))

	DevicesCreateCmd.Flags().StringVar(&DevicesCreateCmdEndpoint, "endpoint", "", TRAPI(""))

	DevicesCreateCmd.Flags().StringVar(&DevicesCreateCmdFirmwareVersion, "firmware-version", "", TRAPI(""))

	DevicesCreateCmd.Flags().StringVar(&DevicesCreateCmdGroupId, "group-id", "", TRAPI(""))

	DevicesCreateCmd.Flags().StringVar(&DevicesCreateCmdImei, "imei", "", TRAPI(""))

	DevicesCreateCmd.Flags().StringVar(&DevicesCreateCmdImsi, "imsi", "", TRAPI(""))

	DevicesCreateCmd.Flags().StringVar(&DevicesCreateCmdIpAddress, "ip-address", "", TRAPI(""))

	DevicesCreateCmd.Flags().StringVar(&DevicesCreateCmdLastModifiedTime, "last-modified-time", "", TRAPI(""))

	DevicesCreateCmd.Flags().StringVar(&DevicesCreateCmdLastRegistrationUpdate, "last-registration-update", "", TRAPI(""))

	DevicesCreateCmd.Flags().StringVar(&DevicesCreateCmdManufacturer, "manufacturer", "", TRAPI(""))

	DevicesCreateCmd.Flags().StringVar(&DevicesCreateCmdModelNumber, "model-number", "", TRAPI(""))

	DevicesCreateCmd.Flags().StringVar(&DevicesCreateCmdOperatorId, "operator-id", "", TRAPI(""))

	DevicesCreateCmd.Flags().StringVar(&DevicesCreateCmdRegistrationId, "registration-id", "", TRAPI(""))

	DevicesCreateCmd.Flags().StringVar(&DevicesCreateCmdSerialNumber, "serial-number", "", TRAPI(""))

	DevicesCreateCmd.Flags().Int64Var(&DevicesCreateCmdRegistrationLifeTime, "registration-life-time", 0, TRAPI(""))

	DevicesCreateCmd.Flags().BoolVar(&DevicesCreateCmdOnline, "online", false, TRAPI(""))

	DevicesCreateCmd.Flags().StringVar(&DevicesCreateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	DevicesCmd.AddCommand(DevicesCreateCmd)
}

// DevicesCreateCmd defines 'create' subcommand
var DevicesCreateCmd = &cobra.Command{
	Use:   "create",
	Short: TRAPI("/devices:post:summary"),
	Long:  TRAPI(`/devices:post:description`),
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

		param, err := collectDevicesCreateCmdParams(ac)
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

func collectDevicesCreateCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForDevicesCreateCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForDevicesCreateCmd("/devices"),
		query:       buildQueryForDevicesCreateCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForDevicesCreateCmd(path string) string {

	return path
}

func buildQueryForDevicesCreateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForDevicesCreateCmd() (string, error) {
	if DevicesCreateCmdBody != "" {
		if strings.HasPrefix(DevicesCreateCmdBody, "@") {
			fname := strings.TrimPrefix(DevicesCreateCmdBody, "@")
			// #nosec
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if DevicesCreateCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return DevicesCreateCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if DevicesCreateCmdDeviceId != "" {
		result["deviceId"] = DevicesCreateCmdDeviceId
	}

	if DevicesCreateCmdEndpoint != "" {
		result["endpoint"] = DevicesCreateCmdEndpoint
	}

	if DevicesCreateCmdFirmwareVersion != "" {
		result["firmwareVersion"] = DevicesCreateCmdFirmwareVersion
	}

	if DevicesCreateCmdGroupId != "" {
		result["groupId"] = DevicesCreateCmdGroupId
	}

	if DevicesCreateCmdImei != "" {
		result["imei"] = DevicesCreateCmdImei
	}

	if DevicesCreateCmdImsi != "" {
		result["imsi"] = DevicesCreateCmdImsi
	}

	if DevicesCreateCmdIpAddress != "" {
		result["ipAddress"] = DevicesCreateCmdIpAddress
	}

	if DevicesCreateCmdLastModifiedTime != "" {
		result["lastModifiedTime"] = DevicesCreateCmdLastModifiedTime
	}

	if DevicesCreateCmdLastRegistrationUpdate != "" {
		result["lastRegistrationUpdate"] = DevicesCreateCmdLastRegistrationUpdate
	}

	if DevicesCreateCmdManufacturer != "" {
		result["manufacturer"] = DevicesCreateCmdManufacturer
	}

	if DevicesCreateCmdModelNumber != "" {
		result["modelNumber"] = DevicesCreateCmdModelNumber
	}

	if DevicesCreateCmdOperatorId != "" {
		result["operatorId"] = DevicesCreateCmdOperatorId
	}

	if DevicesCreateCmdRegistrationId != "" {
		result["registrationId"] = DevicesCreateCmdRegistrationId
	}

	if DevicesCreateCmdSerialNumber != "" {
		result["serialNumber"] = DevicesCreateCmdSerialNumber
	}

	if DevicesCreateCmdRegistrationLifeTime != 0 {
		result["registrationLifeTime"] = DevicesCreateCmdRegistrationLifeTime
	}

	if DevicesCreateCmdOnline != false {
		result["online"] = DevicesCreateCmdOnline
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
