// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"mime"
	"net/url"
	"os"

	"github.com/itchyny/gojq"
	"github.com/soracom/soracom-cli/generators/lib"
	"github.com/spf13/cobra"
)

// DevicesListCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var DevicesListCmdLastEvaluatedKey string

// DevicesListCmdTagName holds value of 'tag_name' option
var DevicesListCmdTagName string

// DevicesListCmdTagValue holds value of 'tag_value' option
var DevicesListCmdTagValue string

// DevicesListCmdTagValueMatchMode holds value of 'tag_value_match_mode' option
var DevicesListCmdTagValueMatchMode string

// DevicesListCmdLimit holds value of 'limit' option
var DevicesListCmdLimit int64

// DevicesListCmdPaginate indicates to do pagination or not
var DevicesListCmdPaginate bool

func init() {
	DevicesListCmd.Flags().StringVar(&DevicesListCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("ID of the last Device in the previous page"))

	DevicesListCmd.Flags().StringVar(&DevicesListCmdTagName, "tag-name", "", TRAPI("Tag name"))

	DevicesListCmd.Flags().StringVar(&DevicesListCmdTagValue, "tag-value", "", TRAPI("Tag value"))

	DevicesListCmd.Flags().StringVar(&DevicesListCmdTagValueMatchMode, "tag-value-match-mode", "", TRAPI("Tag value match mode (exact | prefix)"))

	DevicesListCmd.Flags().Int64Var(&DevicesListCmdLimit, "limit", -1, TRAPI("Max number of Devices in a response"))

	DevicesListCmd.Flags().BoolVar(&DevicesListCmdPaginate, "fetch-all", false, TRCLI("cli.common_params.paginate.short_help"))
	DevicesCmd.AddCommand(DevicesListCmd)
}

// DevicesListCmd defines 'list' subcommand
var DevicesListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/devices:get:summary"),
	Long:  TRAPI(`/devices:get:description`),
	RunE: func(cmd *cobra.Command, args []string) error {

		var jq *gojq.Query
		if jqString != "" {
			var err error
			jq, err = gojq.Parse(jqString)
			if err != nil {
				return err
			}
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

		param, err := collectDevicesListCmdParams(ac)
		if err != nil {
			return err
		}

		body, contentType, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		mediaType, _, err := mime.ParseMediaType(contentType)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if jq != nil {
			if mediaType == "application/json" {
				return processJQ(jq, body)
			} else {
				lib.WarnfStderr(TRCLI("cli.tried-jq-on-non-json") + "\n")
			}
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectDevicesListCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForDevicesListCmd("/devices"),
		query:  buildQueryForDevicesListCmd(),

		doPagination:                      DevicesListCmdPaginate,
		paginationKeyHeaderInResponse:     "x-soracom-next-key",
		paginationRequestParameterInQuery: "last_evaluated_key",

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForDevicesListCmd(path string) string {

	return path
}

func buildQueryForDevicesListCmd() url.Values {
	result := url.Values{}

	if DevicesListCmdLastEvaluatedKey != "" {
		result.Add("last_evaluated_key", DevicesListCmdLastEvaluatedKey)
	}

	if DevicesListCmdTagName != "" {
		result.Add("tag_name", DevicesListCmdTagName)
	}

	if DevicesListCmdTagValue != "" {
		result.Add("tag_value", DevicesListCmdTagValue)
	}

	if DevicesListCmdTagValueMatchMode != "" {
		result.Add("tag_value_match_mode", DevicesListCmdTagValueMatchMode)
	}

	if DevicesListCmdLimit != -1 {
		result.Add("limit", sprintf("%d", DevicesListCmdLimit))
	}

	return result
}
