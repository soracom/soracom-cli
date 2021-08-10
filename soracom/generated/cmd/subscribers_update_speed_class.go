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

// SubscribersUpdateSpeedClassCmdImsi holds value of 'imsi' option
var SubscribersUpdateSpeedClassCmdImsi string

// SubscribersUpdateSpeedClassCmdSpeedClass holds value of 'speedClass' option
var SubscribersUpdateSpeedClassCmdSpeedClass string

// SubscribersUpdateSpeedClassCmdBody holds contents of request body to be sent
var SubscribersUpdateSpeedClassCmdBody string

func init() {
	SubscribersUpdateSpeedClassCmd.Flags().StringVar(&SubscribersUpdateSpeedClassCmdImsi, "imsi", "", TRAPI("IMSI of the target subscriber."))

	SubscribersUpdateSpeedClassCmd.Flags().StringVar(&SubscribersUpdateSpeedClassCmdSpeedClass, "speed-class", "", TRAPI(""))

	SubscribersUpdateSpeedClassCmd.Flags().StringVar(&SubscribersUpdateSpeedClassCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	SubscribersCmd.AddCommand(SubscribersUpdateSpeedClassCmd)
}

// SubscribersUpdateSpeedClassCmd defines 'update-speed-class' subcommand
var SubscribersUpdateSpeedClassCmd = &cobra.Command{
	Use:   "update-speed-class",
	Short: TRAPI("/subscribers/{imsi}/update_speed_class:post:summary"),
	Long:  TRAPI(`/subscribers/{imsi}/update_speed_class:post:description`),
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

		param, err := collectSubscribersUpdateSpeedClassCmdParams(ac)
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

func collectSubscribersUpdateSpeedClassCmdParams(ac *apiClient) (*apiParams, error) {
	body, err := buildBodyForSubscribersUpdateSpeedClassCmd()
	if err != nil {
		return nil, err
	}
	contentType := "application/json"

	if SubscribersUpdateSpeedClassCmdImsi == "" {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "imsi")
		}

	}

	if SubscribersUpdateSpeedClassCmdSpeedClass == "" {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "speed-class")
		}

	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSubscribersUpdateSpeedClassCmd("/subscribers/{imsi}/update_speed_class"),
		query:       buildQueryForSubscribersUpdateSpeedClassCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSubscribersUpdateSpeedClassCmd(path string) string {

	escapedImsi := url.PathEscape(SubscribersUpdateSpeedClassCmdImsi)

	path = strReplace(path, "{"+"imsi"+"}", escapedImsi, -1)

	return path
}

func buildQueryForSubscribersUpdateSpeedClassCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSubscribersUpdateSpeedClassCmd() (string, error) {
	var result map[string]interface{}

	if SubscribersUpdateSpeedClassCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SubscribersUpdateSpeedClassCmdBody, "@") {
			fname := strings.TrimPrefix(SubscribersUpdateSpeedClassCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if SubscribersUpdateSpeedClassCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(SubscribersUpdateSpeedClassCmdBody)
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

	if SubscribersUpdateSpeedClassCmdSpeedClass != "" {
		result["speedClass"] = SubscribersUpdateSpeedClassCmdSpeedClass
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
