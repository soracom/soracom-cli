package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// PaymentStatementsExportCmdExportMode holds value of 'export_mode' option
var PaymentStatementsExportCmdExportMode string

// PaymentStatementsExportCmdPaymentStatementId holds value of 'payment_statement_id' option
var PaymentStatementsExportCmdPaymentStatementId string

func init() {
	PaymentStatementsExportCmd.Flags().StringVar(&PaymentStatementsExportCmdExportMode, "export-mode", "", TRAPI("export_mode (async, sync)"))

	PaymentStatementsExportCmd.Flags().StringVar(&PaymentStatementsExportCmdPaymentStatementId, "payment-statement-id", "", TRAPI("payment_statement_id"))

	PaymentStatementsCmd.AddCommand(PaymentStatementsExportCmd)
}

// PaymentStatementsExportCmd defines 'export' subcommand
var PaymentStatementsExportCmd = &cobra.Command{
	Use:   "export",
	Short: TRAPI("/payment_statements/{payment_statement_id}/export:post:summary"),
	Long:  TRAPI(`/payment_statements/{payment_statement_id}/export:post:description`),
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

		param, err := collectPaymentStatementsExportCmdParams()
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

func collectPaymentStatementsExportCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "POST",
		path:   buildPathForPaymentStatementsExportCmd("/payment_statements/{payment_statement_id}/export"),
		query:  buildQueryForPaymentStatementsExportCmd(),
	}, nil
}

func buildPathForPaymentStatementsExportCmd(path string) string {

	path = strings.Replace(path, "{"+"payment_statement_id"+"}", PaymentStatementsExportCmdPaymentStatementId, -1)

	return path
}

func buildQueryForPaymentStatementsExportCmd() string {
	result := []string{}

	if PaymentStatementsExportCmdExportMode != "" {
		result = append(result, sprintf("%s=%s", "export_mode", PaymentStatementsExportCmdExportMode))
	}

	return strings.Join(result, "&")
}
