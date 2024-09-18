// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// GadgetsListCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var GadgetsListCmdLastEvaluatedKey string

// GadgetsListCmdProductId holds value of 'product_id' option
var GadgetsListCmdProductId string

// GadgetsListCmdTagName holds value of 'tag_name' option
var GadgetsListCmdTagName string

// GadgetsListCmdTagValue holds value of 'tag_value' option
var GadgetsListCmdTagValue string

// GadgetsListCmdTagValueMatchMode holds value of 'tag_value_match_mode' option
var GadgetsListCmdTagValueMatchMode string

// GadgetsListCmdLimit holds value of 'limit' option
var GadgetsListCmdLimit int64

// GadgetsListCmdPaginate indicates to do pagination or not
var GadgetsListCmdPaginate bool

// GadgetsListCmdOutputJSONL indicates to output with jsonl format
var GadgetsListCmdOutputJSONL bool

func InitGadgetsListCmd() {
	GadgetsListCmd.Flags().StringVar(&GadgetsListCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The ID ('{product_id}/{serial_number}') of the last gadget retrieved on the previous page in URL encoding (percent encoding).By specifying this parameter, you can continue to retrieve the list from the next Gadget API compatible device onward.The value of the 'last_evaluated_key' of 'rel=next' returned in the 'link' header when the API is called is expected to be specified in the next call, but any ID ('{productId}/{serialNumber}') can be specified."))

	GadgetsListCmd.Flags().StringVar(&GadgetsListCmdProductId, "product-id", "", TRAPI("Product ID of the target Gadget API compatible device for filtering the search.- 'button': Soracom LTE-M Button powered by AWS.- 'wimax': Soracom Cloud Camera Services Cellular Pack."))

	GadgetsListCmd.Flags().StringVar(&GadgetsListCmdTagName, "tag-name", "", TRAPI("Tag name for filtering the search. The search is always an exact match, regardless of the setting of 'tag_value_match_mode'."))

	GadgetsListCmd.Flags().StringVar(&GadgetsListCmdTagValue, "tag-value", "", TRAPI("Specifies a tag value to search for in a URL-encoded (percent-encoded) string. Required when 'tag_name' has been specified."))

	GadgetsListCmd.Flags().StringVar(&GadgetsListCmdTagValueMatchMode, "tag-value-match-mode", "exact", TRAPI("Tag match mode.- 'exact': exact match.- 'prefix': prefix match."))

	GadgetsListCmd.Flags().Int64Var(&GadgetsListCmdLimit, "limit", 0, TRAPI("Maximum number of Gadget API compatible devices data to retrieve."))

	GadgetsListCmd.Flags().BoolVar(&GadgetsListCmdPaginate, "fetch-all", false, TRCLI("cli.common_params.paginate.short_help"))

	GadgetsListCmd.Flags().BoolVar(&GadgetsListCmdOutputJSONL, "jsonl", false, TRCLI("cli.common_params.jsonl.short_help"))

	GadgetsListCmd.RunE = GadgetsListCmdRunE

	GadgetsCmd.AddCommand(GadgetsListCmd)
}

// GadgetsListCmd defines 'list' subcommand
var GadgetsListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/gadgets:get:summary"),
	Long:  TRAPI(`/gadgets:get:description`) + "\n\n" + createLinkToAPIReference("Gadget", "listGadgets"),
}

func GadgetsListCmdRunE(cmd *cobra.Command, args []string) error {

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
	err := ac.getAPICredentials()
	if err != nil {
		cmd.SilenceUsage = true
		return err
	}

	param, err := collectGadgetsListCmdParams(ac)
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
		if GadgetsListCmdOutputJSONL {
			return printStringAsJSONL(body)
		}

		return prettyPrintStringAsJSON(body)
	}
	return err
}

func collectGadgetsListCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForGadgetsListCmd("/gadgets"),
		query:  buildQueryForGadgetsListCmd(),

		doPagination:                      GadgetsListCmdPaginate,
		paginationKeyHeaderInResponse:     "x-soracom-next-key",
		paginationRequestParameterInQuery: "last_evaluated_key",

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForGadgetsListCmd(path string) string {

	return path
}

func buildQueryForGadgetsListCmd() url.Values {
	result := url.Values{}

	if GadgetsListCmdLastEvaluatedKey != "" {
		result.Add("last_evaluated_key", GadgetsListCmdLastEvaluatedKey)
	}

	if GadgetsListCmdProductId != "" {
		result.Add("product_id", GadgetsListCmdProductId)
	}

	if GadgetsListCmdTagName != "" {
		result.Add("tag_name", GadgetsListCmdTagName)
	}

	if GadgetsListCmdTagValue != "" {
		result.Add("tag_value", GadgetsListCmdTagValue)
	}

	if GadgetsListCmdTagValueMatchMode != "exact" {
		result.Add("tag_value_match_mode", GadgetsListCmdTagValueMatchMode)
	}

	if GadgetsListCmdLimit != 0 {
		result.Add("limit", sprintf("%d", GadgetsListCmdLimit))
	}

	return result
}
