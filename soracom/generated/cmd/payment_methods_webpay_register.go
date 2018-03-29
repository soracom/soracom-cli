package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// PaymentMethodsWebpayRegisterCmdCvc holds value of 'cvc' option
var PaymentMethodsWebpayRegisterCmdCvc string

// PaymentMethodsWebpayRegisterCmdName holds value of 'name' option
var PaymentMethodsWebpayRegisterCmdName string

// PaymentMethodsWebpayRegisterCmdNumber holds value of 'number' option
var PaymentMethodsWebpayRegisterCmdNumber string

// PaymentMethodsWebpayRegisterCmdExpireMonth holds value of 'expireMonth' option
var PaymentMethodsWebpayRegisterCmdExpireMonth int64

// PaymentMethodsWebpayRegisterCmdExpireYear holds value of 'expireYear' option
var PaymentMethodsWebpayRegisterCmdExpireYear int64

// PaymentMethodsWebpayRegisterCmdBody holds contents of request body to be sent
var PaymentMethodsWebpayRegisterCmdBody string

func init() {
	PaymentMethodsWebpayRegisterCmd.Flags().StringVar(&PaymentMethodsWebpayRegisterCmdCvc, "cvc", "", TRAPI(""))

	PaymentMethodsWebpayRegisterCmd.Flags().StringVar(&PaymentMethodsWebpayRegisterCmdName, "name", "", TRAPI(""))

	PaymentMethodsWebpayRegisterCmd.Flags().StringVar(&PaymentMethodsWebpayRegisterCmdNumber, "number", "", TRAPI(""))

	PaymentMethodsWebpayRegisterCmd.Flags().Int64Var(&PaymentMethodsWebpayRegisterCmdExpireMonth, "expire-month", 0, TRAPI(""))

	PaymentMethodsWebpayRegisterCmd.Flags().Int64Var(&PaymentMethodsWebpayRegisterCmdExpireYear, "expire-year", 0, TRAPI(""))

	PaymentMethodsWebpayRegisterCmd.Flags().StringVar(&PaymentMethodsWebpayRegisterCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	PaymentMethodsWebpayCmd.AddCommand(PaymentMethodsWebpayRegisterCmd)
}

// PaymentMethodsWebpayRegisterCmd defines 'register' subcommand
var PaymentMethodsWebpayRegisterCmd = &cobra.Command{
	Use:   "register",
	Short: TRAPI("/payment_methods/webpay:post:summary"),
	Long:  TRAPI(`/payment_methods/webpay:post:description`),
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

		param, err := collectPaymentMethodsWebpayRegisterCmdParams()
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

func collectPaymentMethodsWebpayRegisterCmdParams() (*apiParams, error) {

	body, err := buildBodyForPaymentMethodsWebpayRegisterCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForPaymentMethodsWebpayRegisterCmd("/payment_methods/webpay"),
		query:       buildQueryForPaymentMethodsWebpayRegisterCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForPaymentMethodsWebpayRegisterCmd(path string) string {

	return path
}

func buildQueryForPaymentMethodsWebpayRegisterCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForPaymentMethodsWebpayRegisterCmd() (string, error) {
	if PaymentMethodsWebpayRegisterCmdBody != "" {
		if strings.HasPrefix(PaymentMethodsWebpayRegisterCmdBody, "@") {
			fname := strings.TrimPrefix(PaymentMethodsWebpayRegisterCmdBody, "@")
			// #nosec
			bytes, err := ioutil.ReadFile(fname)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else if PaymentMethodsWebpayRegisterCmdBody == "-" {
			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		} else {
			return PaymentMethodsWebpayRegisterCmdBody, nil
		}
	}

	result := map[string]interface{}{}

	if PaymentMethodsWebpayRegisterCmdCvc != "" {
		result["cvc"] = PaymentMethodsWebpayRegisterCmdCvc
	}

	if PaymentMethodsWebpayRegisterCmdName != "" {
		result["name"] = PaymentMethodsWebpayRegisterCmdName
	}

	if PaymentMethodsWebpayRegisterCmdNumber != "" {
		result["number"] = PaymentMethodsWebpayRegisterCmdNumber
	}

	if PaymentMethodsWebpayRegisterCmdExpireMonth != 0 {
		result["expireMonth"] = PaymentMethodsWebpayRegisterCmdExpireMonth
	}

	if PaymentMethodsWebpayRegisterCmdExpireYear != 0 {
		result["expireYear"] = PaymentMethodsWebpayRegisterCmdExpireYear
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
