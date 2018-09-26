package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SubscribersSetImeiLockCmdImei holds value of 'imei' option
var SubscribersSetImeiLockCmdImei string

// SubscribersSetImeiLockCmdImsi holds value of 'imsi' option
var SubscribersSetImeiLockCmdImsi string

// SubscribersSetImeiLockCmdBody holds contents of request body to be sent
var SubscribersSetImeiLockCmdBody string

func init() {
	SubscribersSetImeiLockCmd.Flags().StringVar(&SubscribersSetImeiLockCmdImei, "imei", "", TRAPI(""))

	SubscribersSetImeiLockCmd.Flags().StringVar(&SubscribersSetImeiLockCmdImsi, "imsi", "", TRAPI("IMSI of the target subscriber."))

	SubscribersSetImeiLockCmd.Flags().StringVar(&SubscribersSetImeiLockCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	SubscribersCmd.AddCommand(SubscribersSetImeiLockCmd)
}

// SubscribersSetImeiLockCmd defines 'set-imei-lock' subcommand
var SubscribersSetImeiLockCmd = &cobra.Command{
	Use:   "set-imei-lock",
	Short: TRAPI("/subscribers/{imsi}/set_imei_lock:post:summary"),
	Long:  TRAPI(`/subscribers/{imsi}/set_imei_lock:post:description`),
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

		param, err := collectSubscribersSetImeiLockCmdParams(ac)
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

func collectSubscribersSetImeiLockCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForSubscribersSetImeiLockCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSubscribersSetImeiLockCmd("/subscribers/{imsi}/set_imei_lock"),
		query:       buildQueryForSubscribersSetImeiLockCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForSubscribersSetImeiLockCmd(path string) string {

	path = strings.Replace(path, "{"+"imsi"+"}", SubscribersSetImeiLockCmdImsi, -1)

	return path
}

func buildQueryForSubscribersSetImeiLockCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForSubscribersSetImeiLockCmd() (string, error) {
	var result map[string]interface{}

	if SubscribersSetImeiLockCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SubscribersSetImeiLockCmdBody, "@") {
			fname := strings.TrimPrefix(SubscribersSetImeiLockCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if SubscribersSetImeiLockCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(SubscribersSetImeiLockCmdBody)
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

	if SubscribersSetImeiLockCmdImei != "" {
		result["imei"] = SubscribersSetImeiLockCmdImei
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
