// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"

	"fmt"

	"io/ioutil"

	"net/url"
	"os"

	"strings"

	"github.com/spf13/cobra"
)

// SimsSetExpiryTimeCmdExpiryAction holds value of 'expiryAction' option
var SimsSetExpiryTimeCmdExpiryAction string

// SimsSetExpiryTimeCmdSimId holds value of 'sim_id' option
var SimsSetExpiryTimeCmdSimId string

// SimsSetExpiryTimeCmdExpiryTime holds value of 'expiryTime' option
var SimsSetExpiryTimeCmdExpiryTime int64

// SimsSetExpiryTimeCmdBody holds contents of request body to be sent
var SimsSetExpiryTimeCmdBody string

func init() {
	SimsSetExpiryTimeCmd.Flags().StringVar(&SimsSetExpiryTimeCmdExpiryAction, "expiry-action", "", TRAPI(""))

	SimsSetExpiryTimeCmd.Flags().StringVar(&SimsSetExpiryTimeCmdSimId, "sim-id", "", TRAPI("SIM ID of the target SIM."))

	SimsSetExpiryTimeCmd.Flags().Int64Var(&SimsSetExpiryTimeCmdExpiryTime, "expiry-time", 0, TRAPI(""))

	SimsSetExpiryTimeCmd.Flags().StringVar(&SimsSetExpiryTimeCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	SimsCmd.AddCommand(SimsSetExpiryTimeCmd)
}

// SimsSetExpiryTimeCmd defines 'set-expiry-time' subcommand
var SimsSetExpiryTimeCmd = &cobra.Command{
	Use:   "set-expiry-time",
	Short: TRAPI("/sims/{sim_id}/set_expiry_time:post:summary"),
	Long:  TRAPI(`/sims/{sim_id}/set_expiry_time:post:description`),
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

		param, err := collectSimsSetExpiryTimeCmdParams(ac)
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

func collectSimsSetExpiryTimeCmdParams(ac *apiClient) (*apiParams, error) {
	body, err := buildBodyForSimsSetExpiryTimeCmd()
	if err != nil {
		return nil, err
	}
	contentType := "application/json"

	if SimsSetExpiryTimeCmdSimId == "" {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "sim-id")
		}

	}

	if SimsSetExpiryTimeCmdExpiryTime == 0 {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "expiry-time")
		}

	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSimsSetExpiryTimeCmd("/sims/{sim_id}/set_expiry_time"),
		query:       buildQueryForSimsSetExpiryTimeCmd(),
		contentType: contentType,
		body:        body,
	}, nil
}

func buildPathForSimsSetExpiryTimeCmd(path string) string {

	escapedSimId := url.PathEscape(SimsSetExpiryTimeCmdSimId)

	path = strReplace(path, "{"+"sim_id"+"}", escapedSimId, -1)

	return path
}

func buildQueryForSimsSetExpiryTimeCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSimsSetExpiryTimeCmd() (string, error) {
	var result map[string]interface{}

	if SimsSetExpiryTimeCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SimsSetExpiryTimeCmdBody, "@") {
			fname := strings.TrimPrefix(SimsSetExpiryTimeCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if SimsSetExpiryTimeCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(SimsSetExpiryTimeCmdBody)
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

	if SimsSetExpiryTimeCmdExpiryAction != "" {
		result["expiryAction"] = SimsSetExpiryTimeCmdExpiryAction
	}

	if SimsSetExpiryTimeCmdExpiryTime != 0 {
		result["expiryTime"] = SimsSetExpiryTimeCmdExpiryTime
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
