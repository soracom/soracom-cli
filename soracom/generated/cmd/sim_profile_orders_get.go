// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SimProfileOrdersGetCmdProfileOrderId holds value of 'profile_order_id' option
var SimProfileOrdersGetCmdProfileOrderId string

func InitSimProfileOrdersGetCmd() {
	SimProfileOrdersGetCmd.Flags().StringVar(&SimProfileOrdersGetCmdProfileOrderId, "profile-order-id", "", TRAPI("The ID of the eSIM profile order."))

	SimProfileOrdersGetCmd.RunE = SimProfileOrdersGetCmdRunE

	SimProfileOrdersCmd.AddCommand(SimProfileOrdersGetCmd)
}

// SimProfileOrdersGetCmd defines 'get' subcommand
var SimProfileOrdersGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/sim_profile_orders/{profile_order_id}:get:summary"),
	Long:  TRAPI(`/sim_profile_orders/{profile_order_id}:get:description`) + "\n\n" + createLinkToAPIReference("SimProfileOrder", "getProfileOrder"),
}

func SimProfileOrdersGetCmdRunE(cmd *cobra.Command, args []string) error {

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

	param, err := collectSimProfileOrdersGetCmdParams(ac)
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

func collectSimProfileOrdersGetCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("profile_order_id", "profile-order-id", "path", parsedBody, SimProfileOrdersGetCmdProfileOrderId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForSimProfileOrdersGetCmd("/sim_profile_orders/{profile_order_id}"),
		query:  buildQueryForSimProfileOrdersGetCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForSimProfileOrdersGetCmd(path string) string {

	escapedProfileOrderId := url.PathEscape(SimProfileOrdersGetCmdProfileOrderId)

	path = strReplace(path, "{"+"profile_order_id"+"}", escapedProfileOrderId, -1)

	return path
}

func buildQueryForSimProfileOrdersGetCmd() url.Values {
	result := url.Values{}

	return result
}
