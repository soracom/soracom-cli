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

// SoraletsTestCmdContentType holds value of 'contentType' option
var SoraletsTestCmdContentType string

// SoraletsTestCmdDirection holds value of 'direction' option
var SoraletsTestCmdDirection string

// SoraletsTestCmdEncodingType holds value of 'encodingType' option
var SoraletsTestCmdEncodingType string

// SoraletsTestCmdPayload holds value of 'payload' option
var SoraletsTestCmdPayload string

// SoraletsTestCmdSoraletId holds value of 'soralet_id' option
var SoraletsTestCmdSoraletId string

// SoraletsTestCmdUserdata holds value of 'userdata' option
var SoraletsTestCmdUserdata string

// SoraletsTestCmdVersion holds value of 'version' option
var SoraletsTestCmdVersion string

// SoraletsTestCmdBody holds contents of request body to be sent
var SoraletsTestCmdBody string

func init() {
	SoraletsTestCmd.Flags().StringVar(&SoraletsTestCmdContentType, "content-type", "", TRAPI(""))

	SoraletsTestCmd.Flags().StringVar(&SoraletsTestCmdDirection, "direction", "", TRAPI(""))

	SoraletsTestCmd.Flags().StringVar(&SoraletsTestCmdEncodingType, "encoding-type", "", TRAPI(""))

	SoraletsTestCmd.Flags().StringVar(&SoraletsTestCmdPayload, "payload", "", TRAPI(""))

	SoraletsTestCmd.Flags().StringVar(&SoraletsTestCmdSoraletId, "soralet-id", "", TRAPI("The identifier of Soralet."))

	SoraletsTestCmd.Flags().StringVar(&SoraletsTestCmdUserdata, "userdata", "", TRAPI(""))

	SoraletsTestCmd.Flags().StringVar(&SoraletsTestCmdVersion, "version", "", TRAPI(""))

	SoraletsTestCmd.Flags().StringVar(&SoraletsTestCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	SoraletsCmd.AddCommand(SoraletsTestCmd)
}

// SoraletsTestCmd defines 'test' subcommand
var SoraletsTestCmd = &cobra.Command{
	Use:   "test",
	Short: TRAPI("/soralets/{soralet_id}/test:post:summary"),
	Long:  TRAPI(`/soralets/{soralet_id}/test:post:description`) + "\n\n" + createLinkToAPIReference("Soralet", "testSoralet"),
	RunE: func(cmd *cobra.Command, args []string) error {

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

		param, err := collectSoraletsTestCmdParams(ac)
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
	},
}

func collectSoraletsTestCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForSoraletsTestCmd()
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

	err = checkIfRequiredStringParameterIsSupplied("contentType", "content-type", "body", parsedBody, SoraletsTestCmdContentType)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("direction", "direction", "body", parsedBody, SoraletsTestCmdDirection)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("payload", "payload", "body", parsedBody, SoraletsTestCmdPayload)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("soralet_id", "soralet-id", "path", parsedBody, SoraletsTestCmdSoraletId)
	if err != nil {
		return nil, err
	}

	err = checkIfRequiredStringParameterIsSupplied("version", "version", "body", parsedBody, SoraletsTestCmdVersion)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSoraletsTestCmd("/soralets/{soralet_id}/test"),
		query:       buildQueryForSoraletsTestCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSoraletsTestCmd(path string) string {

	escapedSoraletId := url.PathEscape(SoraletsTestCmdSoraletId)

	path = strReplace(path, "{"+"soralet_id"+"}", escapedSoraletId, -1)

	return path
}

func buildQueryForSoraletsTestCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForSoraletsTestCmd() (string, error) {
	var result map[string]interface{}

	if SoraletsTestCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SoraletsTestCmdBody, "@") {
			fname := strings.TrimPrefix(SoraletsTestCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if SoraletsTestCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(SoraletsTestCmdBody)
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

	if SoraletsTestCmdContentType != "" {
		result["contentType"] = SoraletsTestCmdContentType
	}

	if SoraletsTestCmdDirection != "" {
		result["direction"] = SoraletsTestCmdDirection
	}

	if SoraletsTestCmdEncodingType != "" {
		result["encodingType"] = SoraletsTestCmdEncodingType
	}

	if SoraletsTestCmdPayload != "" {
		result["payload"] = SoraletsTestCmdPayload
	}

	if SoraletsTestCmdUserdata != "" {
		result["userdata"] = SoraletsTestCmdUserdata
	}

	if SoraletsTestCmdVersion != "" {
		result["version"] = SoraletsTestCmdVersion
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
