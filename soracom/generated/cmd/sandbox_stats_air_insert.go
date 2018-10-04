package cmd

import (
	"encoding/json"

	"io/ioutil"

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

func init() {
	SandboxStatsAirInsertCmd.Flags().StringVar(&SandboxStatsAirInsertCmdImsi, "imsi", "", TRAPI("IMSI"))

	SandboxStatsAirInsertCmd.Flags().Int64Var(&SandboxStatsAirInsertCmdUnixtime, "unixtime", 0, TRAPI(""))

	SandboxStatsAirInsertCmd.Flags().StringVar(&SandboxStatsAirInsertCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	SandboxStatsAirCmd.AddCommand(SandboxStatsAirInsertCmd)
}

// SandboxStatsAirInsertCmd defines 'insert' subcommand
var SandboxStatsAirInsertCmd = &cobra.Command{
	Use:   "insert",
	Short: TRAPI("/sandbox/stats/air/subscribers/{imsi}:post:summary"),
	Long:  TRAPI(`/sandbox/stats/air/subscribers/{imsi}:post:description`),
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

		param, err := collectSandboxStatsAirInsertCmdParams(ac)
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

func collectSandboxStatsAirInsertCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForSandboxStatsAirInsertCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSandboxStatsAirInsertCmd("/sandbox/stats/air/subscribers/{imsi}"),
		query:       buildQueryForSandboxStatsAirInsertCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForSandboxStatsAirInsertCmd(path string) string {

	path = strings.Replace(path, "{"+"imsi"+"}", SandboxStatsAirInsertCmdImsi, -1)

	return path
}

func buildQueryForSandboxStatsAirInsertCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForSandboxStatsAirInsertCmd() (string, error) {
	var result map[string]interface{}

	if SandboxStatsAirInsertCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SandboxStatsAirInsertCmdBody, "@") {
			fname := strings.TrimPrefix(SandboxStatsAirInsertCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if SandboxStatsAirInsertCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
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

	if SandboxStatsAirInsertCmdUnixtime != 0 {
		result["unixtime"] = SandboxStatsAirInsertCmdUnixtime
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
