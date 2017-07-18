package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {

	CouponsCmd.AddCommand(CouponsListCmd)
}

// CouponsListCmd defines 'list' subcommand
var CouponsListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/coupons:get:summary"),
	Long:  TRAPI(`/coupons:get:description`),
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

		param, err := collectCouponsListCmdParams()
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

func collectCouponsListCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForCouponsListCmd("/coupons"),
		query:  buildQueryForCouponsListCmd(),
	}, nil
}

func buildPathForCouponsListCmd(path string) string {

	return path
}

func buildQueryForCouponsListCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
