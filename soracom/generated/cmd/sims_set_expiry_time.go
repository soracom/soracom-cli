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

// SimsSetExpiryTimeCmdExpiryAction holds value of 'expiryAction' option
var SimsSetExpiryTimeCmdExpiryAction string

// SimsSetExpiryTimeCmdSimId holds value of 'sim_id' option
var SimsSetExpiryTimeCmdSimId string

// SimsSetExpiryTimeCmdExpiryTime holds value of 'expiryTime' option
var SimsSetExpiryTimeCmdExpiryTime int64

// SimsSetExpiryTimeCmdBody holds contents of request body to be sent
var SimsSetExpiryTimeCmdBody string

func init() {
	SimsSetExpiryTimeCmd.Flags().StringVar(&SimsSetExpiryTimeCmdExpiryAction, "expiry-action", "", TRAPI("Action at expiration. Specify one of the following Please refer to [Soracom Air Expiration Function](https://developers.soracom.io/en/docs/air/expiration/) for more detail. You have to disable termination protection if you want to specify `terminate` as an action.If omitted, a null value is set.- `doNothing` : do nothing- `deleteSession` : delete session of the SIM if any- `deactivate` : change the SIM status to Inactive- `suspend` : change the SIM status to Suspended- `terminate` : forcibly end any existing connections, and terminate the SIM- null value : not set (It works the same as `doNothing`)"))

	SimsSetExpiryTimeCmd.Flags().StringVar(&SimsSetExpiryTimeCmdSimId, "sim-id", "", TRAPI("SIM ID of the target SIM."))

	SimsSetExpiryTimeCmd.Flags().Int64Var(&SimsSetExpiryTimeCmdExpiryTime, "expiry-time", 0, TRAPI("Timestamp of date and time set using the Expiration function (UNIX time in milliseconds)"))

	SimsSetExpiryTimeCmd.Flags().StringVar(&SimsSetExpiryTimeCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	SimsCmd.AddCommand(SimsSetExpiryTimeCmd)
}

// SimsSetExpiryTimeCmd defines 'set-expiry-time' subcommand
var SimsSetExpiryTimeCmd = &cobra.Command{
	Use:   "set-expiry-time",
	Short: TRAPI("/sims/{sim_id}/set_expiry_time:post:summary"),
	Long:  TRAPI(`/sims/{sim_id}/set_expiry_time:post:description`) + "\n\n" + createLinkToAPIReference("Sim", "setSimExpiryTime"),
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
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForSimsSetExpiryTimeCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("sim_id", "sim-id", "path", parsedBody, SimsSetExpiryTimeCmdSimId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredIntegerParameterIsSupplied("expiryTime", "expiry-time", "body", parsedBody, SimsSetExpiryTimeCmdExpiryTime)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSimsSetExpiryTimeCmd("/sims/{sim_id}/set_expiry_time"),
		query:       buildQueryForSimsSetExpiryTimeCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
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
			b, err = os.ReadFile(fname)
		} else if SimsSetExpiryTimeCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
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

	result["expiryTime"] = SimsSetExpiryTimeCmdExpiryTime

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
