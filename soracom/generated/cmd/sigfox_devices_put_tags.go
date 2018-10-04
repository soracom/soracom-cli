package cmd

import (
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SigfoxDevicesPutTagsCmdDeviceId holds value of 'device_id' option
var SigfoxDevicesPutTagsCmdDeviceId string

// SigfoxDevicesPutTagsCmdBody holds contents of request body to be sent
var SigfoxDevicesPutTagsCmdBody string

func init() {
	SigfoxDevicesPutTagsCmd.Flags().StringVar(&SigfoxDevicesPutTagsCmdDeviceId, "device-id", "", TRAPI("Device ID of the target Sigfox device."))

	SigfoxDevicesPutTagsCmd.Flags().StringVar(&SigfoxDevicesPutTagsCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	SigfoxDevicesCmd.AddCommand(SigfoxDevicesPutTagsCmd)
}

// SigfoxDevicesPutTagsCmd defines 'put-tags' subcommand
var SigfoxDevicesPutTagsCmd = &cobra.Command{
	Use:   "put-tags",
	Short: TRAPI("/sigfox_devices/{device_id}/tags:put:summary"),
	Long:  TRAPI(`/sigfox_devices/{device_id}/tags:put:description`),
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

		param, err := collectSigfoxDevicesPutTagsCmdParams(ac)
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

func collectSigfoxDevicesPutTagsCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForSigfoxDevicesPutTagsCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForSigfoxDevicesPutTagsCmd("/sigfox_devices/{device_id}/tags"),
		query:       buildQueryForSigfoxDevicesPutTagsCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForSigfoxDevicesPutTagsCmd(path string) string {

	path = strings.Replace(path, "{"+"device_id"+"}", SigfoxDevicesPutTagsCmdDeviceId, -1)

	return path
}

func buildQueryForSigfoxDevicesPutTagsCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForSigfoxDevicesPutTagsCmd() (string, error) {
	var b []byte
	var err error

	if SigfoxDevicesPutTagsCmdBody != "" {
		if strings.HasPrefix(SigfoxDevicesPutTagsCmdBody, "@") {
			fname := strings.TrimPrefix(SigfoxDevicesPutTagsCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if SigfoxDevicesPutTagsCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(SigfoxDevicesPutTagsCmdBody)
		}

		if err != nil {
			return "", err
		}
	}

	if b == nil {
		b = []byte{}
	}

	return string(b), nil
}
