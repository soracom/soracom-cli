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

// SandboxStatsAirInsertCmdImsi holds value of 'imsi' option
var SandboxStatsAirInsertCmdImsi string

// SandboxStatsAirInsertCmdUnixtime holds value of 'unixtime' option
var SandboxStatsAirInsertCmdUnixtime int64

// SandboxStatsAirInsertCmdBody holds contents of request body to be sent
var SandboxStatsAirInsertCmdBody string

func InitSandboxStatsAirInsertCmd() {
	SandboxStatsAirInsertCmd.Flags().StringVar(&SandboxStatsAirInsertCmdImsi, "imsi", "", TRAPI("IMSI"))

	SandboxStatsAirInsertCmd.Flags().Int64Var(&SandboxStatsAirInsertCmdUnixtime, "unixtime", 0, TRAPI("UNIX time (in milliseconds)"))

	SandboxStatsAirInsertCmd.Flags().StringVar(&SandboxStatsAirInsertCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	SandboxStatsAirInsertCmd.RunE = SandboxStatsAirInsertCmdRunE

	SandboxStatsAirCmd.AddCommand(SandboxStatsAirInsertCmd)
}

// SandboxStatsAirInsertCmd defines 'insert' subcommand
var SandboxStatsAirInsertCmd = &cobra.Command{
	Use:   "insert",
	Short: TRAPI("/sandbox/stats/air/subscribers/{imsi}:post:summary"),
	Long:  TRAPI(`/sandbox/stats/air/subscribers/{imsi}:post:description`) + "\n\n" + createLinkToAPIReference("Stats", "sandboxInsertAirStats"),
}

func SandboxStatsAirInsertCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSandboxStatsAirInsertCmdParams(ac)
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

func collectSandboxStatsAirInsertCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForSandboxStatsAirInsertCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("imsi", "imsi", "path", parsedBody, SandboxStatsAirInsertCmdImsi)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSandboxStatsAirInsertCmd("/sandbox/stats/air/subscribers/{imsi}"),
		query:       buildQueryForSandboxStatsAirInsertCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSandboxStatsAirInsertCmd(path string) string {

	escapedImsi := url.PathEscape(SandboxStatsAirInsertCmdImsi)

	path = strReplace(path, "{"+"imsi"+"}", escapedImsi, -1)

	return path
}

func buildQueryForSandboxStatsAirInsertCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSandboxStatsAirInsertCmd() (string, error) {
	var result map[string]interface{}

	if SandboxStatsAirInsertCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SandboxStatsAirInsertCmdBody, "@") {
			fname := strings.TrimPrefix(SandboxStatsAirInsertCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if SandboxStatsAirInsertCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(SandboxStatsAirInsertCmdBody)
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

	if SandboxStatsAirInsertCmd.Flags().Lookup("unixtime").Changed {
		result["unixtime"] = SandboxStatsAirInsertCmdUnixtime
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
