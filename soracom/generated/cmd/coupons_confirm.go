// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// CouponsConfirmCmdOrderId holds value of 'order_id' option
var CouponsConfirmCmdOrderId string

func init() {
	CouponsConfirmCmd.Flags().StringVar(&CouponsConfirmCmdOrderId, "order-id", "", TRAPI("order_id"))
	CouponsCmd.AddCommand(CouponsConfirmCmd)
}

// CouponsConfirmCmd defines 'confirm' subcommand
var CouponsConfirmCmd = &cobra.Command{
	Use:   "confirm",
	Short: TRAPI("/coupons/{order_id}/confirm:put:summary"),
	Long:  TRAPI(`/coupons/{order_id}/confirm:put:description`) + "\n\n" + createLinkToAPIReference("Order", "confirmCouponOrder"),
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

		param, err := collectCouponsConfirmCmdParams(ac)
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

func collectCouponsConfirmCmdParams(ac *apiClient) (*apiParams, error) {
	var parsedBody interface{}
	var err error
	err = checkIfRequiredStringParameterIsSupplied("order_id", "order-id", "path", parsedBody, CouponsConfirmCmdOrderId)
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method: "PUT",
		path:   buildPathForCouponsConfirmCmd("/coupons/{order_id}/confirm"),
		query:  buildQueryForCouponsConfirmCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForCouponsConfirmCmd(path string) string {

	escapedOrderId := url.PathEscape(CouponsConfirmCmdOrderId)

	path = strReplace(path, "{"+"order_id"+"}", escapedOrderId, -1)

	return path
}

func buildQueryForCouponsConfirmCmd() url.Values {
	result := url.Values{}

	return result
}
