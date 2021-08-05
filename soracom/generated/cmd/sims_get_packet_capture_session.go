// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SimsGetPacketCaptureSessionCmdSessionId holds value of 'session_id' option
var SimsGetPacketCaptureSessionCmdSessionId string

// SimsGetPacketCaptureSessionCmdSimId holds value of 'sim_id' option
var SimsGetPacketCaptureSessionCmdSimId string

func init() {
	SimsGetPacketCaptureSessionCmd.Flags().StringVar(&SimsGetPacketCaptureSessionCmdSessionId, "session-id", "", TRAPI("Packet capture session ID"))

	SimsGetPacketCaptureSessionCmd.Flags().StringVar(&SimsGetPacketCaptureSessionCmdSimId, "sim-id", "", TRAPI("SIM ID"))
	SimsCmd.AddCommand(SimsGetPacketCaptureSessionCmd)
}

// SimsGetPacketCaptureSessionCmd defines 'get-packet-capture-session' subcommand
var SimsGetPacketCaptureSessionCmd = &cobra.Command{
	Use:   "get-packet-capture-session",
	Short: TRAPI("/sims/{sim_id}/packet_capture_sessions/{session_id}:get:summary"),
	Long:  TRAPI(`/sims/{sim_id}/packet_capture_sessions/{session_id}:get:description`),
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

		param, err := collectSimsGetPacketCaptureSessionCmdParams(ac)
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

		if queryString != "" {
			return processQuery(queryString, body)
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectSimsGetPacketCaptureSessionCmdParams(ac *apiClient) (*apiParams, error) {
	if SimsGetPacketCaptureSessionCmdSessionId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "session-id")
	}

	if SimsGetPacketCaptureSessionCmdSimId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "sim-id")
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForSimsGetPacketCaptureSessionCmd("/sims/{sim_id}/packet_capture_sessions/{session_id}"),
		query:  buildQueryForSimsGetPacketCaptureSessionCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSimsGetPacketCaptureSessionCmd(path string) string {

	escapedSessionId := url.PathEscape(SimsGetPacketCaptureSessionCmdSessionId)

	path = strReplace(path, "{"+"session_id"+"}", escapedSessionId, -1)

	escapedSimId := url.PathEscape(SimsGetPacketCaptureSessionCmdSimId)

	path = strReplace(path, "{"+"sim_id"+"}", escapedSimId, -1)

	return path
}

func buildQueryForSimsGetPacketCaptureSessionCmd() url.Values {
	result := url.Values{}

	return result
}
