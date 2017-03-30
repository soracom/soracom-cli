package cmd

import (
	"encoding/json"
	"io/ioutil"

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

// LoraNetworkSetsCreateCmdBody holds contents of request body to be sent
var LoraNetworkSetsCreateCmdBody string

func init() {
	LoraNetworkSetsCreateCmd.Flags().StringVar(&LoraNetworkSetsCreateCmdCreatedTime, "created-time", "", TR(""))

	LoraNetworkSetsCreateCmd.Flags().StringVar(&LoraNetworkSetsCreateCmdLastModifiedTime, "last-modified-time", "", TR(""))

	LoraNetworkSetsCreateCmd.Flags().StringVar(&LoraNetworkSetsCreateCmdNetworkSetId, "network-set-id", "", TR(""))

	LoraNetworkSetsCreateCmd.Flags().StringVar(&LoraNetworkSetsCreateCmdOperatorId, "operator-id", "", TR(""))

	LoraNetworkSetsCreateCmd.Flags().StringVar(&LoraNetworkSetsCreateCmdBody, "body", "", TR("cli.common_params.body.short_help"))

	LoraNetworkSetsCmd.AddCommand(LoraNetworkSetsCreateCmd)
}

// LoraNetworkSetsCreateCmd defines 'create' subcommand
var LoraNetworkSetsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: TR("lora_network_sets.create.summary"),
	Long:  TR(`lora_network_sets.create.description`),
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

		param, err := collectLoraNetworkSetsCreateCmdParams()
		if err != nil {
			return err
		}

		result, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if result == "" {
			return nil
		}

		return prettyPrintStringAsJSON(result)
	},
}

func collectLoraNetworkSetsCreateCmdParams() (*apiParams, error) {

	body, err := buildBodyForLoraNetworkSetsCreateCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForLoraNetworkSetsCreateCmd("/lora_network_sets"),
		query:       buildQueryForLoraNetworkSetsCreateCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForLoraNetworkSetsCreateCmd(path string) string {

	return path
}

func buildQueryForLoraNetworkSetsCreateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForLoraNetworkSetsCreateCmd() (string, error) {
	if LoraNetworkSetsCreateCmdBody != "" {
		if strings.HasPrefix(LoraNetworkSetsCreateCmdBody, "@") {
			fname := strings.TrimPrefix(LoraNetworkSetsCreateCmdBody, "@")
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if LoraNetworkSetsCreateCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return LoraNetworkSetsCreateCmdBody, nil
		}
	}

	result := map[string]interface{}{}

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

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
