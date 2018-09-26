package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SandboxCouponsCreateCmdApplicableBillItemName holds value of 'applicableBillItemName' option
var SandboxCouponsCreateCmdApplicableBillItemName string

// SandboxCouponsCreateCmdExpiryYearMonth holds value of 'expiryYearMonth' option
var SandboxCouponsCreateCmdExpiryYearMonth string

// SandboxCouponsCreateCmdAmount holds value of 'amount' option
var SandboxCouponsCreateCmdAmount int64

// SandboxCouponsCreateCmdBody holds contents of request body to be sent
var SandboxCouponsCreateCmdBody string

func init() {
	SandboxCouponsCreateCmd.Flags().StringVar(&SandboxCouponsCreateCmdApplicableBillItemName, "applicable-bill-item-name", "", TRAPI(""))

	SandboxCouponsCreateCmd.Flags().StringVar(&SandboxCouponsCreateCmdExpiryYearMonth, "expiry-year-month", "", TRAPI(""))

	SandboxCouponsCreateCmd.Flags().Int64Var(&SandboxCouponsCreateCmdAmount, "amount", 0, TRAPI(""))

	SandboxCouponsCreateCmd.Flags().StringVar(&SandboxCouponsCreateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	SandboxCouponsCmd.AddCommand(SandboxCouponsCreateCmd)
}

// SandboxCouponsCreateCmd defines 'create' subcommand
var SandboxCouponsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: TRAPI("/sandbox/coupons/create:post:summary"),
	Long:  TRAPI(`/sandbox/coupons/create:post:description`),
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

		param, err := collectSandboxCouponsCreateCmdParams(ac)
		if err != nil {
			return err
		}

		_, body, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}

		return prettyPrintStringAsJSON(body)
	},
}

func collectSandboxCouponsCreateCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForSandboxCouponsCreateCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForSandboxCouponsCreateCmd("/sandbox/coupons/create"),
		query:       buildQueryForSandboxCouponsCreateCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForSandboxCouponsCreateCmd(path string) string {

	return path
}

func buildQueryForSandboxCouponsCreateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForSandboxCouponsCreateCmd() (string, error) {
	var result map[string]interface{}

	if SandboxCouponsCreateCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SandboxCouponsCreateCmdBody, "@") {
			fname := strings.TrimPrefix(SandboxCouponsCreateCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if SandboxCouponsCreateCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(SandboxCouponsCreateCmdBody)
		}

		if err != nil {
			return "", err
		}

		err = json.Unmarshal(b, &result)
		if err != nil {
			return "", err
		}
	}

	if result == nil {
		result = make(map[string]interface{})
	}

	if SandboxCouponsCreateCmdApplicableBillItemName != "" {
		result["applicableBillItemName"] = SandboxCouponsCreateCmdApplicableBillItemName
	}

	if SandboxCouponsCreateCmdExpiryYearMonth != "" {
		result["expiryYearMonth"] = SandboxCouponsCreateCmdExpiryYearMonth
	}

	if SandboxCouponsCreateCmdAmount != 0 {
		result["amount"] = SandboxCouponsCreateCmdAmount
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
