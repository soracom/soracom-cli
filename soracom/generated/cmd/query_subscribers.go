package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// QuerySubscribersCmdGroup holds value of 'group' option
var QuerySubscribersCmdGroup string

// QuerySubscribersCmdIccid holds value of 'iccid' option
var QuerySubscribersCmdIccid string

// QuerySubscribersCmdImsi holds value of 'imsi' option
var QuerySubscribersCmdImsi string

// QuerySubscribersCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var QuerySubscribersCmdLastEvaluatedKey string

// QuerySubscribersCmdMsisdn holds value of 'msisdn' option
var QuerySubscribersCmdMsisdn string

// QuerySubscribersCmdName holds value of 'name' option
var QuerySubscribersCmdName string

// QuerySubscribersCmdSearchType holds value of 'search_type' option
var QuerySubscribersCmdSearchType string

// QuerySubscribersCmdSerialNumber holds value of 'serial_number' option
var QuerySubscribersCmdSerialNumber string

// QuerySubscribersCmdTag holds value of 'tag' option
var QuerySubscribersCmdTag string

// QuerySubscribersCmdLimit holds value of 'limit' option
var QuerySubscribersCmdLimit int64

func init() {
	QuerySubscribersCmd.Flags().StringVar(&QuerySubscribersCmdGroup, "group", "", TRAPI("Group name to search"))

	QuerySubscribersCmd.Flags().StringVar(&QuerySubscribersCmdIccid, "iccid", "", TRAPI("ICCID to search"))

	QuerySubscribersCmd.Flags().StringVar(&QuerySubscribersCmdImsi, "imsi", "", TRAPI("IMSI to search"))

	QuerySubscribersCmd.Flags().StringVar(&QuerySubscribersCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The IMSI of the last subscriber retrieved on the current page. By specifying this parameter, you can continue to retrieve the list from the next subscriber onward."))

	QuerySubscribersCmd.Flags().StringVar(&QuerySubscribersCmdMsisdn, "msisdn", "", TRAPI("MSISDN to search"))

	QuerySubscribersCmd.Flags().StringVar(&QuerySubscribersCmdName, "name", "", TRAPI("Name to search"))

	QuerySubscribersCmd.Flags().StringVar(&QuerySubscribersCmdSearchType, "search-type", "", TRAPI("Type of the search ('AND searching' or 'OR searching')"))

	QuerySubscribersCmd.Flags().StringVar(&QuerySubscribersCmdSerialNumber, "serial-number", "", TRAPI("Serial number to search"))

	QuerySubscribersCmd.Flags().StringVar(&QuerySubscribersCmdTag, "tag", "", TRAPI("String of tag values to search"))

	QuerySubscribersCmd.Flags().Int64Var(&QuerySubscribersCmdLimit, "limit", 0, TRAPI("The maximum number of item to retrieve"))

	QueryCmd.AddCommand(QuerySubscribersCmd)
}

// QuerySubscribersCmd defines 'subscribers' subcommand
var QuerySubscribersCmd = &cobra.Command{
	Use:   "subscribers",
	Short: TRAPI("/query/subscribers:get:summary"),
	Long:  TRAPI(`/query/subscribers:get:description`),
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

		param, err := collectQuerySubscribersCmdParams()
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

func collectQuerySubscribersCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForQuerySubscribersCmd("/query/subscribers"),
		query:  buildQueryForQuerySubscribersCmd(),
	}, nil
}

func buildPathForQuerySubscribersCmd(path string) string {

	return path
}

func buildQueryForQuerySubscribersCmd() string {
	result := []string{}

	if QuerySubscribersCmdGroup != "" {
		result = append(result, sprintf("%s=%s", "group", QuerySubscribersCmdGroup))
	}

	if QuerySubscribersCmdIccid != "" {
		result = append(result, sprintf("%s=%s", "iccid", QuerySubscribersCmdIccid))
	}

	if QuerySubscribersCmdImsi != "" {
		result = append(result, sprintf("%s=%s", "imsi", QuerySubscribersCmdImsi))
	}

	if QuerySubscribersCmdLastEvaluatedKey != "" {
		result = append(result, sprintf("%s=%s", "last_evaluated_key", QuerySubscribersCmdLastEvaluatedKey))
	}

	if QuerySubscribersCmdMsisdn != "" {
		result = append(result, sprintf("%s=%s", "msisdn", QuerySubscribersCmdMsisdn))
	}

	if QuerySubscribersCmdName != "" {
		result = append(result, sprintf("%s=%s", "name", QuerySubscribersCmdName))
	}

	if QuerySubscribersCmdSearchType != "" {
		result = append(result, sprintf("%s=%s", "search_type", QuerySubscribersCmdSearchType))
	}

	if QuerySubscribersCmdSerialNumber != "" {
		result = append(result, sprintf("%s=%s", "serial_number", QuerySubscribersCmdSerialNumber))
	}

	if QuerySubscribersCmdTag != "" {
		result = append(result, sprintf("%s=%s", "tag", QuerySubscribersCmdTag))
	}

	if QuerySubscribersCmdLimit != 0 {
		result = append(result, sprintf("%s=%d", "limit", QuerySubscribersCmdLimit))
	}

	return strings.Join(result, "&")
}
