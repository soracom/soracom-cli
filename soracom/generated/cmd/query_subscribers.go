// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/soracom/soracom-cli/generators/lib"

	"github.com/spf13/cobra"
)

// QuerySubscribersCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var QuerySubscribersCmdLastEvaluatedKey string

// QuerySubscribersCmdSearchType holds value of 'search_type' option
var QuerySubscribersCmdSearchType string

// QuerySubscribersCmdGroup holds multiple values of 'group' option
var QuerySubscribersCmdGroup []string

// QuerySubscribersCmdIccid holds multiple values of 'iccid' option
var QuerySubscribersCmdIccid []string

// QuerySubscribersCmdImsi holds multiple values of 'imsi' option
var QuerySubscribersCmdImsi []string

// QuerySubscribersCmdModuleType holds multiple values of 'module_type' option
var QuerySubscribersCmdModuleType []string

// QuerySubscribersCmdMsisdn holds multiple values of 'msisdn' option
var QuerySubscribersCmdMsisdn []string

// QuerySubscribersCmdName holds multiple values of 'name' option
var QuerySubscribersCmdName []string

// QuerySubscribersCmdSerialNumber holds multiple values of 'serial_number' option
var QuerySubscribersCmdSerialNumber []string

// QuerySubscribersCmdSubscription holds multiple values of 'subscription' option
var QuerySubscribersCmdSubscription []string

// QuerySubscribersCmdTag holds multiple values of 'tag' option
var QuerySubscribersCmdTag []string

// QuerySubscribersCmdLimit holds value of 'limit' option
var QuerySubscribersCmdLimit int64

// QuerySubscribersCmdPaginate indicates to do pagination or not
var QuerySubscribersCmdPaginate bool

// QuerySubscribersCmdOutputJSONL indicates to output with jsonl format
var QuerySubscribersCmdOutputJSONL bool

func init() {
	QuerySubscribersCmd.Flags().StringVar(&QuerySubscribersCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The IMSI of the last subscriber retrieved on the current page. By specifying this parameter, you can continue to retrieve the list from the next subscriber onward."))

	QuerySubscribersCmd.Flags().StringVar(&QuerySubscribersCmdSearchType, "search-type", "and", TRAPI("Type of the search ('AND searching' or 'OR searching')"))

	QuerySubscribersCmd.Flags().StringSliceVar(&QuerySubscribersCmdGroup, "group", []string{}, TRAPI("Group name to search"))

	QuerySubscribersCmd.Flags().StringSliceVar(&QuerySubscribersCmdIccid, "iccid", []string{}, TRAPI("ICCID to search"))

	QuerySubscribersCmd.Flags().StringSliceVar(&QuerySubscribersCmdImsi, "imsi", []string{}, TRAPI("IMSI to search"))

	QuerySubscribersCmd.Flags().StringSliceVar(&QuerySubscribersCmdModuleType, "module-type", []string{}, TRAPI("Module type (e.g. `mini`, `virtual`) to search"))

	QuerySubscribersCmd.Flags().StringSliceVar(&QuerySubscribersCmdMsisdn, "msisdn", []string{}, TRAPI("MSISDN to search"))

	QuerySubscribersCmd.Flags().StringSliceVar(&QuerySubscribersCmdName, "name", []string{}, TRAPI("Name to search"))

	QuerySubscribersCmd.Flags().StringSliceVar(&QuerySubscribersCmdSerialNumber, "serial-number", []string{}, TRAPI("Serial number to search"))

	QuerySubscribersCmd.Flags().StringSliceVar(&QuerySubscribersCmdSubscription, "subscription", []string{}, TRAPI("Subscription (e.g. `plan01s`) to search"))

	QuerySubscribersCmd.Flags().StringSliceVar(&QuerySubscribersCmdTag, "tag", []string{}, TRAPI("String of tag values to search"))

	QuerySubscribersCmd.Flags().Int64Var(&QuerySubscribersCmdLimit, "limit", 10, TRAPI("The maximum number of item to retrieve"))

	QuerySubscribersCmd.Flags().BoolVar(&QuerySubscribersCmdPaginate, "fetch-all", false, TRCLI("cli.common_params.paginate.short_help"))

	QuerySubscribersCmd.Flags().BoolVar(&QuerySubscribersCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))
	QueryCmd.AddCommand(QuerySubscribersCmd)
}

// QuerySubscribersCmd defines 'subscribers' subcommand
var QuerySubscribersCmd = &cobra.Command{
	Use:   "subscribers",
	Short: TRAPI("/query/subscribers:get:summary"),
	Long:  TRAPI(`/query/subscribers:get:description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		lib.WarnfStderr(TRCLI("cli.deprecated-api") + "\n")
		lib.WarnfStderr(TRCLI("cli.alternative-api-suggestion")+"\n", "query sims")

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

		param, err := collectQuerySubscribersCmdParams(ac)
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
			if QuerySubscribersCmdOutputJSONL {
				return printStringAsJSONL(body)
			}

			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectQuerySubscribersCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForQuerySubscribersCmd("/query/subscribers"),
		query:  buildQueryForQuerySubscribersCmd(),

		doPagination:                      QuerySubscribersCmdPaginate,
		paginationKeyHeaderInResponse:     "x-soracom-next-key",
		paginationRequestParameterInQuery: "last_evaluated_key",

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForQuerySubscribersCmd(path string) string {

	return path
}

func buildQueryForQuerySubscribersCmd() url.Values {
	result := url.Values{}

	if QuerySubscribersCmdLastEvaluatedKey != "" {
		result.Add("last_evaluated_key", QuerySubscribersCmdLastEvaluatedKey)
	}

	if QuerySubscribersCmdSearchType != "and" {
		result.Add("search_type", QuerySubscribersCmdSearchType)
	}

	for _, s := range QuerySubscribersCmdGroup {
		if s != "" {
			result.Add("group", s)
		}
	}

	for _, s := range QuerySubscribersCmdIccid {
		if s != "" {
			result.Add("iccid", s)
		}
	}

	for _, s := range QuerySubscribersCmdImsi {
		if s != "" {
			result.Add("imsi", s)
		}
	}

	for _, s := range QuerySubscribersCmdModuleType {
		if s != "" {
			result.Add("module_type", s)
		}
	}

	for _, s := range QuerySubscribersCmdMsisdn {
		if s != "" {
			result.Add("msisdn", s)
		}
	}

	for _, s := range QuerySubscribersCmdName {
		if s != "" {
			result.Add("name", s)
		}
	}

	for _, s := range QuerySubscribersCmdSerialNumber {
		if s != "" {
			result.Add("serial_number", s)
		}
	}

	for _, s := range QuerySubscribersCmdSubscription {
		if s != "" {
			result.Add("subscription", s)
		}
	}

	for _, s := range QuerySubscribersCmdTag {
		if s != "" {
			result.Add("tag", s)
		}
	}

	if QuerySubscribersCmdLimit != 10 {
		result.Add("limit", sprintf("%d", QuerySubscribersCmdLimit))
	}

	return result
}
