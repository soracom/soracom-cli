// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SimsDeletePacketCaptureSessionCmdSessionId holds value of 'session_id' option
var SimsDeletePacketCaptureSessionCmdSessionId string

// SimsDeletePacketCaptureSessionCmdSimId holds value of 'sim_id' option
var SimsDeletePacketCaptureSessionCmdSimId string

func InitSimsDeletePacketCaptureSessionCmd() {
	SimsDeletePacketCaptureSessionCmd.Flags().StringVar(&SimsDeletePacketCaptureSessionCmdSessionId, "session-id", "", TRAPI("Packet capture session ID"))

	SimsDeletePacketCaptureSessionCmd.Flags().StringVar(&SimsDeletePacketCaptureSessionCmdSimId, "sim-id", "", TRAPI("SIM ID of the target SIM."))

	SimsDeletePacketCaptureSessionCmd.RunE = SimsDeletePacketCaptureSessionCmdRunE

	SimsCmd.AddCommand(SimsDeletePacketCaptureSessionCmd)
}

// SimsDeletePacketCaptureSessionCmd defines 'delete-packet-capture-session' subcommand
var SimsDeletePacketCaptureSessionCmd = &cobra.Command{
	Use:   "delete-packet-capture-session",
	Short: TRAPI("/sims/{sim_id}/packet_capture_sessions/{session_id}:delete:summary"),
	Long:  TRAPI(`/sims/{sim_id}/packet_capture_sessions/{session_id}:delete:description`) + "\n\n" + createLinkToAPIReference("Sim", "deleteSimPacketCaptureSession"),
}

func SimsDeletePacketCaptureSessionCmdRunE(cmd *cobra.Command, args []string) error {

	if len(args) > 0 {
		return fmt.Errorf("unexpected arguments passed => %v", args)
	}

	opt := &apiClientOptions{
		BasePath: "/v1",
		Language: getSelectedLanguage(),
		Profile:  getProfileIfExists(),
	}

	ac := newAPIClient(opt)
	if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
		ac.SetVerbose(true)
	}
	err := ac.getAPICredentials()
	if err != nil {
		cmd.SilenceUsage = true
		return err
	}

	param, err := collectSimsDeletePacketCaptureSessionCmdParams(ac)
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
}

func collectSimsDeletePacketCaptureSessionCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("session_id", "session-id", "path", parsedBody, SimsDeletePacketCaptureSessionCmdSessionId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("sim_id", "sim-id", "path", parsedBody, SimsDeletePacketCaptureSessionCmdSimId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForSimsDeletePacketCaptureSessionCmd("/sims/{sim_id}/packet_capture_sessions/{session_id}"),
		query:  buildQueryForSimsDeletePacketCaptureSessionCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSimsDeletePacketCaptureSessionCmd(path string) string {

	escapedSessionId := url.PathEscape(SimsDeletePacketCaptureSessionCmdSessionId)

	path = strReplace(path, "{"+"session_id"+"}", escapedSessionId, -1)

	escapedSimId := url.PathEscape(SimsDeletePacketCaptureSessionCmdSimId)

	path = strReplace(path, "{"+"sim_id"+"}", escapedSimId, -1)

	return path
}

func buildQueryForSimsDeletePacketCaptureSessionCmd() url.Values {
	result := url.Values{}

	return result
}
