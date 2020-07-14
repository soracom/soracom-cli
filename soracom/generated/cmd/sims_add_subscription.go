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

// SimsAddSubscriptionCmdIccid holds value of 'iccid' option
var SimsAddSubscriptionCmdIccid string

// SimsAddSubscriptionCmdSimId holds value of 'sim_id' option
var SimsAddSubscriptionCmdSimId string

// SimsAddSubscriptionCmdBody holds contents of request body to be sent
var SimsAddSubscriptionCmdBody string

func init() {
	SimsAddSubscriptionCmd.Flags().StringVar(&SimsAddSubscriptionCmdIccid, "iccid", "", TRAPI("Iccid of the target profile"))

	SimsAddSubscriptionCmd.Flags().StringVar(&SimsAddSubscriptionCmdSimId, "sim-id", "", TRAPI("Id of the target SIM"))

	SimsAddSubscriptionCmd.Flags().StringVar(&SimsAddSubscriptionCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	SimsCmd.AddCommand(SimsAddSubscriptionCmd)
}

// SimsAddSubscriptionCmd defines 'add-subscription' subcommand
var SimsAddSubscriptionCmd = &cobra.Command{
	Use:   "add-subscription",
	Short: TRAPI("/sims/{sim_id}/profiles/{iccid}/add_subscription:post:summary"),
	Long:  TRAPI(`/sims/{sim_id}/profiles/{iccid}/add_subscription:post:description`),
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

		param, err := collectSimsAddSubscriptionCmdParams(ac)
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

func collectSimsAddSubscriptionCmdParams(ac *apiClient) (*apiParams, error) {
	body, err := buildBodyForSimsAddSubscriptionCmd()
	if err != nil {
		return nil, err
	}
	contentType := "application/json"

	if SimsAddSubscriptionCmdIccid == "" {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "iccid")
		}

	}

	if SimsAddSubscriptionCmdSimId == "" {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "sim-id")
		}

	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSimsAddSubscriptionCmd("/sims/{sim_id}/profiles/{iccid}/add_subscription"),
		query:       buildQueryForSimsAddSubscriptionCmd(),
		contentType: contentType,
		body:        body,
	}, nil
}

func buildPathForSimsAddSubscriptionCmd(path string) string {

	escapedIccid := url.PathEscape(SimsAddSubscriptionCmdIccid)

	path = strReplace(path, "{"+"iccid"+"}", escapedIccid, -1)

	escapedSimId := url.PathEscape(SimsAddSubscriptionCmdSimId)

	path = strReplace(path, "{"+"sim_id"+"}", escapedSimId, -1)

	return path
}

func buildQueryForSimsAddSubscriptionCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSimsAddSubscriptionCmd() (string, error) {
	var result map[string]interface{}

	if SimsAddSubscriptionCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SimsAddSubscriptionCmdBody, "@") {
			fname := strings.TrimPrefix(SimsAddSubscriptionCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if SimsAddSubscriptionCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(SimsAddSubscriptionCmdBody)
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

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
