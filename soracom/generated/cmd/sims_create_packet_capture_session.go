// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"

	"fmt"

	"io/ioutil"

	"mime"
	"net/url"
	"os"

	"strings"

	"github.com/itchyny/gojq"
	"github.com/soracom/soracom-cli/generators/lib"
	"github.com/spf13/cobra"
)

// SimsCreatePacketCaptureSessionCmdPrefix holds value of 'prefix' option
var SimsCreatePacketCaptureSessionCmdPrefix string

// SimsCreatePacketCaptureSessionCmdSimId holds value of 'sim_id' option
var SimsCreatePacketCaptureSessionCmdSimId string

// SimsCreatePacketCaptureSessionCmdDuration holds value of 'duration' option
var SimsCreatePacketCaptureSessionCmdDuration int64

// SimsCreatePacketCaptureSessionCmdBody holds contents of request body to be sent
var SimsCreatePacketCaptureSessionCmdBody string

func init() {
	SimsCreatePacketCaptureSessionCmd.Flags().StringVar(&SimsCreatePacketCaptureSessionCmdPrefix, "prefix", "", TRAPI(""))

	SimsCreatePacketCaptureSessionCmd.Flags().StringVar(&SimsCreatePacketCaptureSessionCmdSimId, "sim-id", "", TRAPI("The SIM ID for which a packet capture session is created"))

	SimsCreatePacketCaptureSessionCmd.Flags().Int64Var(&SimsCreatePacketCaptureSessionCmdDuration, "duration", 0, TRAPI(""))

	SimsCreatePacketCaptureSessionCmd.Flags().StringVar(&SimsCreatePacketCaptureSessionCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	SimsCmd.AddCommand(SimsCreatePacketCaptureSessionCmd)
}

// SimsCreatePacketCaptureSessionCmd defines 'create-packet-capture-session' subcommand
var SimsCreatePacketCaptureSessionCmd = &cobra.Command{
	Use:   "create-packet-capture-session",
	Short: TRAPI("/sims/{sim_id}/packet_capture_sessions:post:summary"),
	Long:  TRAPI(`/sims/{sim_id}/packet_capture_sessions:post:description`),
	RunE: func(cmd *cobra.Command, args []string) error {

		var jq *gojq.Query
		if jqString != "" {
			var err error
			jq, err = gojq.Parse(jqString)
			if err != nil {
				return err
			}
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

		param, err := collectSimsCreatePacketCaptureSessionCmdParams(ac)
		if err != nil {
			return err
		}

		body, contentType, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		mediaType, _, err := mime.ParseMediaType(contentType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if jq != nil {
			if mediaType == "application/json" {
				return processJQ(jq, body)
			} else {
				lib.WarnfStderr(TRCLI("cli.tried-jq-on-non-json") + "\n")
			}
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectSimsCreatePacketCaptureSessionCmdParams(ac *apiClient) (*apiParams, error) {
	body, err := buildBodyForSimsCreatePacketCaptureSessionCmd()
	if err != nil {
		return nil, err
	}
	contentType := "application/json"

	if SimsCreatePacketCaptureSessionCmdSimId == "" {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "sim-id")
		}

	}

	if SimsCreatePacketCaptureSessionCmdDuration == 0 {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "duration")
		}

	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSimsCreatePacketCaptureSessionCmd("/sims/{sim_id}/packet_capture_sessions"),
		query:       buildQueryForSimsCreatePacketCaptureSessionCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSimsCreatePacketCaptureSessionCmd(path string) string {

	escapedSimId := url.PathEscape(SimsCreatePacketCaptureSessionCmdSimId)

	path = strReplace(path, "{"+"sim_id"+"}", escapedSimId, -1)

	return path
}

func buildQueryForSimsCreatePacketCaptureSessionCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSimsCreatePacketCaptureSessionCmd() (string, error) {
	var result map[string]interface{}

	if SimsCreatePacketCaptureSessionCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SimsCreatePacketCaptureSessionCmdBody, "@") {
			fname := strings.TrimPrefix(SimsCreatePacketCaptureSessionCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if SimsCreatePacketCaptureSessionCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(SimsCreatePacketCaptureSessionCmdBody)
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

	if SimsCreatePacketCaptureSessionCmdPrefix != "" {
		result["prefix"] = SimsCreatePacketCaptureSessionCmdPrefix
	}

	if SimsCreatePacketCaptureSessionCmdDuration != 0 {
		result["duration"] = SimsCreatePacketCaptureSessionCmdDuration
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
