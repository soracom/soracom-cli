// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// QuerySigfoxDevicesCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var QuerySigfoxDevicesCmdLastEvaluatedKey string

// QuerySigfoxDevicesCmdRegistration holds value of 'registration' option
var QuerySigfoxDevicesCmdRegistration string

// QuerySigfoxDevicesCmdSearchType holds value of 'search_type' option
var QuerySigfoxDevicesCmdSearchType string

// QuerySigfoxDevicesCmdStatus holds value of 'status' option
var QuerySigfoxDevicesCmdStatus string

// QuerySigfoxDevicesCmdDeviceId holds multiple values of 'deviceId' option
var QuerySigfoxDevicesCmdDeviceId []string

// QuerySigfoxDevicesCmdGroup holds multiple values of 'group' option
var QuerySigfoxDevicesCmdGroup []string

// QuerySigfoxDevicesCmdName holds multiple values of 'name' option
var QuerySigfoxDevicesCmdName []string

// QuerySigfoxDevicesCmdTag holds multiple values of 'tag' option
var QuerySigfoxDevicesCmdTag []string

// QuerySigfoxDevicesCmdLimit holds value of 'limit' option
var QuerySigfoxDevicesCmdLimit int64

// QuerySigfoxDevicesCmdPaginate indicates to do pagination or not
var QuerySigfoxDevicesCmdPaginate bool

func init() {
	QuerySigfoxDevicesCmd.Flags().StringVar(&QuerySigfoxDevicesCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The Sigfox device ID of the last Sigfox device retrieved on the current page. By specifying this parameter, you can continue to retrieve the list from the next Sigfox device onward."))

	QuerySigfoxDevicesCmd.Flags().StringVar(&QuerySigfoxDevicesCmdRegistration, "registration", "and", TRAPI("Registration status of sigfox devices"))

	QuerySigfoxDevicesCmd.Flags().StringVar(&QuerySigfoxDevicesCmdSearchType, "search-type", "and", TRAPI("Type of the search ('AND searching' or 'OR searching')"))

	QuerySigfoxDevicesCmd.Flags().StringVar(&QuerySigfoxDevicesCmdStatus, "status", "and", TRAPI("Status of sigfox devices"))

	QuerySigfoxDevicesCmd.Flags().StringSliceVar(&QuerySigfoxDevicesCmdDeviceId, "device-id", []string{}, TRAPI("Sigfox device ID to search"))

	QuerySigfoxDevicesCmd.Flags().StringSliceVar(&QuerySigfoxDevicesCmdGroup, "group", []string{}, TRAPI("Group name to search"))

	QuerySigfoxDevicesCmd.Flags().StringSliceVar(&QuerySigfoxDevicesCmdName, "name", []string{}, TRAPI("Name to search"))

	QuerySigfoxDevicesCmd.Flags().StringSliceVar(&QuerySigfoxDevicesCmdTag, "tag", []string{}, TRAPI("String of tag values to search"))

	QuerySigfoxDevicesCmd.Flags().Int64Var(&QuerySigfoxDevicesCmdLimit, "limit", 10, TRAPI("The maximum number of item to retrieve"))

	QuerySigfoxDevicesCmd.Flags().BoolVar(&QuerySigfoxDevicesCmdPaginate, "fetch-all", false, TRCLI("cli.common_params.paginate.short_help"))
	QueryCmd.AddCommand(QuerySigfoxDevicesCmd)
}

// QuerySigfoxDevicesCmd defines 'sigfox-devices' subcommand
var QuerySigfoxDevicesCmd = &cobra.Command{
	Use:   "sigfox-devices",
	Short: TRAPI("/query/sigfox_devices:get:summary"),
	Long:  TRAPI(`/query/sigfox_devices:get:description`),
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

		param, err := collectQuerySigfoxDevicesCmdParams(ac)
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

		if queryString != "" {
			return processQuery(queryString, body)
		}

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectQuerySigfoxDevicesCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForQuerySigfoxDevicesCmd("/query/sigfox_devices"),
		query:  buildQueryForQuerySigfoxDevicesCmd(),

		doPagination:                      QuerySigfoxDevicesCmdPaginate,
		paginationKeyHeaderInResponse:     "x-soracom-next-key",
		paginationRequestParameterInQuery: "last_evaluated_key",

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForQuerySigfoxDevicesCmd(path string) string {

	return path
}

func buildQueryForQuerySigfoxDevicesCmd() url.Values {
	result := url.Values{}

	if QuerySigfoxDevicesCmdLastEvaluatedKey != "" {
		result.Add("last_evaluated_key", QuerySigfoxDevicesCmdLastEvaluatedKey)
	}

	if QuerySigfoxDevicesCmdRegistration != "and" {
		result.Add("registration", QuerySigfoxDevicesCmdRegistration)
	}

	if QuerySigfoxDevicesCmdSearchType != "and" {
		result.Add("search_type", QuerySigfoxDevicesCmdSearchType)
	}

	if QuerySigfoxDevicesCmdStatus != "and" {
		result.Add("status", QuerySigfoxDevicesCmdStatus)
	}

	for _, s := range QuerySigfoxDevicesCmdDeviceId {
		if s != "" {
			result.Add("deviceId", s)
		}
	}

	for _, s := range QuerySigfoxDevicesCmdGroup {
		if s != "" {
			result.Add("group", s)
		}
	}

	for _, s := range QuerySigfoxDevicesCmdName {
		if s != "" {
			result.Add("name", s)
		}
	}

	for _, s := range QuerySigfoxDevicesCmdTag {
		if s != "" {
			result.Add("tag", s)
		}
	}

	if QuerySigfoxDevicesCmdLimit != 10 {
		result.Add("limit", sprintf("%d", QuerySigfoxDevicesCmdLimit))
	}

	return result
}
