// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SimsStopPacketCaptureSessionCmdSessionId holds value of 'session_id' option
var SimsStopPacketCaptureSessionCmdSessionId string

// SimsStopPacketCaptureSessionCmdSimId holds value of 'sim_id' option
var SimsStopPacketCaptureSessionCmdSimId string

func init() {
	SimsStopPacketCaptureSessionCmd.Flags().StringVar(&SimsStopPacketCaptureSessionCmdSessionId, "session-id", "", TRAPI("Packet capture session ID"))

	SimsStopPacketCaptureSessionCmd.Flags().StringVar(&SimsStopPacketCaptureSessionCmdSimId, "sim-id", "", TRAPI("SIM ID"))
	SimsCmd.AddCommand(SimsStopPacketCaptureSessionCmd)
}

// SimsStopPacketCaptureSessionCmd defines 'stop-packet-capture-session' subcommand
var SimsStopPacketCaptureSessionCmd = &cobra.Command{
	Use:   "stop-packet-capture-session",
	Short: TRAPI("/sims/{sim_id}/packet_capture_sessions/{session_id}/stop:post:summary"),
	Long:  TRAPI(`/sims/{sim_id}/packet_capture_sessions/{session_id}/stop:post:description`),
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

		param, err := collectSimsStopPacketCaptureSessionCmdParams(ac)
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

		if jqString != "" {
			return processJQ(jqString, body)
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectSimsStopPacketCaptureSessionCmdParams(ac *apiClient) (*apiParams, error) {
	if SimsStopPacketCaptureSessionCmdSessionId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "session-id")
	}

	if SimsStopPacketCaptureSessionCmdSimId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "sim-id")
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForSimsStopPacketCaptureSessionCmd("/sims/{sim_id}/packet_capture_sessions/{session_id}/stop"),
		query:  buildQueryForSimsStopPacketCaptureSessionCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSimsStopPacketCaptureSessionCmd(path string) string {

	escapedSessionId := url.PathEscape(SimsStopPacketCaptureSessionCmdSessionId)

	path = strReplace(path, "{"+"session_id"+"}", escapedSessionId, -1)

	escapedSimId := url.PathEscape(SimsStopPacketCaptureSessionCmdSimId)

	path = strReplace(path, "{"+"sim_id"+"}", escapedSimId, -1)

	return path
}

func buildQueryForSimsStopPacketCaptureSessionCmd() url.Values {
	result := url.Values{}

	return result
}
