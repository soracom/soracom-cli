package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// CouponsRegisterCmdCouponCode holds value of 'coupon_code' option
var CouponsRegisterCmdCouponCode string

func init() {
	CouponsRegisterCmd.Flags().StringVar(&CouponsRegisterCmdCouponCode, "coupon-code", "", TRAPI("coupon_code"))

	CouponsCmd.AddCommand(CouponsRegisterCmd)
}

// CouponsRegisterCmd defines 'register' subcommand
var CouponsRegisterCmd = &cobra.Command{
	Use:   "register",
	Short: TRAPI("/coupons/{coupon_code}/register:post:summary"),
	Long:  TRAPI(`/coupons/{coupon_code}/register:post:description`),
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

		param, err := collectCouponsRegisterCmdParams(ac)
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

func collectCouponsRegisterCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForCouponsRegisterCmd("/coupons/{coupon_code}/register"),
		query:  buildQueryForCouponsRegisterCmd(),
	}, nil
}

func buildPathForCouponsRegisterCmd(path string) string {

	path = strings.Replace(path, "{"+"coupon_code"+"}", CouponsRegisterCmdCouponCode, -1)

	return path
}

func buildQueryForCouponsRegisterCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
