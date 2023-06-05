// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// VpgDeletePacketCaptureSessionCmdSessionId holds value of 'session_id' option
var VpgDeletePacketCaptureSessionCmdSessionId string

// VpgDeletePacketCaptureSessionCmdVpgId holds value of 'vpg_id' option
var VpgDeletePacketCaptureSessionCmdVpgId string

func InitVpgDeletePacketCaptureSessionCmd() {
	VpgDeletePacketCaptureSessionCmd.Flags().StringVar(&VpgDeletePacketCaptureSessionCmdSessionId, "session-id", "", TRAPI("Packet capture session ID"))

	VpgDeletePacketCaptureSessionCmd.Flags().StringVar(&VpgDeletePacketCaptureSessionCmdVpgId, "vpg-id", "", TRAPI("VPG ID"))

	VpgDeletePacketCaptureSessionCmd.RunE = VpgDeletePacketCaptureSessionCmdRunE

	VpgCmd.AddCommand(VpgDeletePacketCaptureSessionCmd)
}

// VpgDeletePacketCaptureSessionCmd defines 'delete-packet-capture-session' subcommand
var VpgDeletePacketCaptureSessionCmd = &cobra.Command{
	Use:   "delete-packet-capture-session",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/packet_capture_sessions/{session_id}:delete:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/packet_capture_sessions/{session_id}:delete:description`) + "\n\n" + createLinkToAPIReference("VirtualPrivateGateway", "deletePacketCaptureSession"),
}

func VpgDeletePacketCaptureSessionCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectVpgDeletePacketCaptureSessionCmdParams(ac)
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

func collectVpgDeletePacketCaptureSessionCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("session_id", "session-id", "path", parsedBody, VpgDeletePacketCaptureSessionCmdSessionId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("vpg_id", "vpg-id", "path", parsedBody, VpgDeletePacketCaptureSessionCmdVpgId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForVpgDeletePacketCaptureSessionCmd("/virtual_private_gateways/{vpg_id}/packet_capture_sessions/{session_id}"),
		query:  buildQueryForVpgDeletePacketCaptureSessionCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForVpgDeletePacketCaptureSessionCmd(path string) string {

	escapedSessionId := url.PathEscape(VpgDeletePacketCaptureSessionCmdSessionId)

	path = strReplace(path, "{"+"session_id"+"}", escapedSessionId, -1)

	escapedVpgId := url.PathEscape(VpgDeletePacketCaptureSessionCmdVpgId)

	path = strReplace(path, "{"+"vpg_id"+"}", escapedVpgId, -1)

	return path
}

func buildQueryForVpgDeletePacketCaptureSessionCmd() url.Values {
	result := url.Values{}

	return result
}
