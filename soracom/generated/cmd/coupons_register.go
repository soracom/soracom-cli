// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"mime"
	"net/url"
	"os"

	"github.com/itchyny/gojq"
	"github.com/soracom/soracom-cli/generators/lib"
	"github.com/spf13/cobra"
)

// CouponsRegisterCmdCouponCode holds value of 'coupon_code' option
var CouponsRegisterCmdCouponCode string

func init() {
	CouponsRegisterCmd.Flags().StringVar(&CouponsRegisterCmdCouponCode, "coupon-code", "", TRAPI("Coupon code"))
	CouponsCmd.AddCommand(CouponsRegisterCmd)
}

// CouponsRegisterCmd defines 'register' subcommand
var CouponsRegisterCmd = &cobra.Command{
	Use:   "register",
	Short: TRAPI("/coupons/{coupon_code}/register:post:summary"),
	Long:  TRAPI(`/coupons/{coupon_code}/register:post:description`),
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

		param, err := collectCouponsRegisterCmdParams(ac)
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

func collectCouponsRegisterCmdParams(ac *apiClient) (*apiParams, error) {
	if CouponsRegisterCmdCouponCode == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "coupon-code")
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForCouponsRegisterCmd("/coupons/{coupon_code}/register"),
		query:  buildQueryForCouponsRegisterCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForCouponsRegisterCmd(path string) string {

	escapedCouponCode := url.PathEscape(CouponsRegisterCmdCouponCode)

	path = strReplace(path, "{"+"coupon_code"+"}", escapedCouponCode, -1)

	return path
}

func buildQueryForCouponsRegisterCmd() url.Values {
	result := url.Values{}

	return result
}
