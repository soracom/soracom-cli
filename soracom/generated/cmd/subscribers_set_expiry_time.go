package cmd

import (
	"encoding/json"

	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SubscribersSetExpiryTimeCmdExpiryAction holds value of 'expiryAction' option
var SubscribersSetExpiryTimeCmdExpiryAction string

// SubscribersSetExpiryTimeCmdImsi holds value of 'imsi' option
var SubscribersSetExpiryTimeCmdImsi string

// SubscribersSetExpiryTimeCmdExpiryTime holds value of 'expiryTime' option
var SubscribersSetExpiryTimeCmdExpiryTime int64

// SubscribersSetExpiryTimeCmdBody holds contents of request body to be sent
var SubscribersSetExpiryTimeCmdBody string

func init() {
	SubscribersSetExpiryTimeCmd.Flags().StringVar(&SubscribersSetExpiryTimeCmdExpiryAction, "expiry-action", "", TRAPI(""))

	SubscribersSetExpiryTimeCmd.Flags().StringVar(&SubscribersSetExpiryTimeCmdImsi, "imsi", "", TRAPI("IMSI of the target subscriber."))

	SubscribersSetExpiryTimeCmd.Flags().Int64Var(&SubscribersSetExpiryTimeCmdExpiryTime, "expiry-time", 0, TRAPI(""))

	SubscribersSetExpiryTimeCmd.Flags().StringVar(&SubscribersSetExpiryTimeCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	SubscribersCmd.AddCommand(SubscribersSetExpiryTimeCmd)
}

// SubscribersSetExpiryTimeCmd defines 'set-expiry-time' subcommand
var SubscribersSetExpiryTimeCmd = &cobra.Command{
	Use:   "set-expiry-time",
	Short: TRAPI("/subscribers/{imsi}/set_expiry_time:post:summary"),
	Long:  TRAPI(`/subscribers/{imsi}/set_expiry_time:post:description`),
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

		param, err := collectSubscribersSetExpiryTimeCmdParams(ac)
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

func collectSubscribersSetExpiryTimeCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForSubscribersSetExpiryTimeCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSubscribersSetExpiryTimeCmd("/subscribers/{imsi}/set_expiry_time"),
		query:       buildQueryForSubscribersSetExpiryTimeCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForSubscribersSetExpiryTimeCmd(path string) string {

	path = strings.Replace(path, "{"+"imsi"+"}", SubscribersSetExpiryTimeCmdImsi, -1)

	return path
}

func buildQueryForSubscribersSetExpiryTimeCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForSubscribersSetExpiryTimeCmd() (string, error) {
	var result map[string]interface{}

	if SubscribersSetExpiryTimeCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SubscribersSetExpiryTimeCmdBody, "@") {
			fname := strings.TrimPrefix(SubscribersSetExpiryTimeCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if SubscribersSetExpiryTimeCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(SubscribersSetExpiryTimeCmdBody)
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

	if SubscribersSetExpiryTimeCmdExpiryAction != "" {
		result["expiryAction"] = SubscribersSetExpiryTimeCmdExpiryAction
	}

	if SubscribersSetExpiryTimeCmdExpiryTime != 0 {
		result["expiryTime"] = SubscribersSetExpiryTimeCmdExpiryTime
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
