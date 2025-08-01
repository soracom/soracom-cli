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

// LoraNetworkSetsCreateCmdCreatedTime holds value of 'createdTime' option
var LoraNetworkSetsCreateCmdCreatedTime string

// LoraNetworkSetsCreateCmdLastModifiedTime holds value of 'lastModifiedTime' option
var LoraNetworkSetsCreateCmdLastModifiedTime string

// LoraNetworkSetsCreateCmdNetworkSetId holds value of 'networkSetId' option
var LoraNetworkSetsCreateCmdNetworkSetId string

// LoraNetworkSetsCreateCmdOperatorId holds value of 'operatorId' option
var LoraNetworkSetsCreateCmdOperatorId string

// LoraNetworkSetsCreateCmdAllowedOperators holds multiple values of 'allowedOperators' option
var LoraNetworkSetsCreateCmdAllowedOperators []string

// LoraNetworkSetsCreateCmdBody holds contents of request body to be sent
var LoraNetworkSetsCreateCmdBody string

func InitLoraNetworkSetsCreateCmd() {
	LoraNetworkSetsCreateCmd.Flags().StringVar(&LoraNetworkSetsCreateCmdCreatedTime, "created-time", "", TRAPI(""))

	LoraNetworkSetsCreateCmd.Flags().StringVar(&LoraNetworkSetsCreateCmdLastModifiedTime, "last-modified-time", "", TRAPI(""))

	LoraNetworkSetsCreateCmd.Flags().StringVar(&LoraNetworkSetsCreateCmdNetworkSetId, "network-set-id", "", TRAPI(""))

	LoraNetworkSetsCreateCmd.Flags().StringVar(&LoraNetworkSetsCreateCmdOperatorId, "operator-id", "", TRAPI(""))

	LoraNetworkSetsCreateCmd.Flags().StringSliceVar(&LoraNetworkSetsCreateCmdAllowedOperators, "allowed-operators", []string{}, TRAPI(""))

	LoraNetworkSetsCreateCmd.Flags().StringVar(&LoraNetworkSetsCreateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	LoraNetworkSetsCreateCmd.RunE = LoraNetworkSetsCreateCmdRunE

	LoraNetworkSetsCmd.AddCommand(LoraNetworkSetsCreateCmd)
}

// LoraNetworkSetsCreateCmd defines 'create' subcommand
var LoraNetworkSetsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: TRAPI("/lora_network_sets:post:summary"),
	Long:  TRAPI(`/lora_network_sets:post:description`) + "\n\n" + createLinkToAPIReference("LoraNetworkSet", "createLoraNetworkSet"),
}

func LoraNetworkSetsCreateCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectLoraNetworkSetsCreateCmdParams(ac)
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

func collectLoraNetworkSetsCreateCmdParams(ac *apiClient) (*apiParams, error) {
	var body string
	var parsedBody interface{}
	var err error
	body, err = buildBodyForLoraNetworkSetsCreateCmd()
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

	return &apiParams{
		method:      "POST",
		path:        buildPathForLoraNetworkSetsCreateCmd("/lora_network_sets"),
		query:       buildQueryForLoraNetworkSetsCreateCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLoraNetworkSetsCreateCmd(path string) string {

	return path
}

func buildQueryForLoraNetworkSetsCreateCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForLoraNetworkSetsCreateCmd() (string, error) {
	var result map[string]interface{}

	if LoraNetworkSetsCreateCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(LoraNetworkSetsCreateCmdBody, "@") {
			fname := strings.TrimPrefix(LoraNetworkSetsCreateCmdBody, "@")
			// #nosec
			b, err = os.ReadFile(fname)
		} else if LoraNetworkSetsCreateCmdBody == "-" {
			b, err = io.ReadAll(os.Stdin)
		} else {
			b = []byte(LoraNetworkSetsCreateCmdBody)
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

	if LoraNetworkSetsCreateCmdCreatedTime != "" {
		result["createdTime"] = LoraNetworkSetsCreateCmdCreatedTime
	}

	if LoraNetworkSetsCreateCmdLastModifiedTime != "" {
		result["lastModifiedTime"] = LoraNetworkSetsCreateCmdLastModifiedTime
	}

	if LoraNetworkSetsCreateCmdNetworkSetId != "" {
		result["networkSetId"] = LoraNetworkSetsCreateCmdNetworkSetId
	}

	if LoraNetworkSetsCreateCmdOperatorId != "" {
		result["operatorId"] = LoraNetworkSetsCreateCmdOperatorId
	}

	if len(LoraNetworkSetsCreateCmdAllowedOperators) != 0 {
		result["allowedOperators"] = LoraNetworkSetsCreateCmdAllowedOperators
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
