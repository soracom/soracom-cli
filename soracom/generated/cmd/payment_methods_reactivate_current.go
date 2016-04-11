package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)











func init() {



  PaymentMethodsCmd.AddCommand(PaymentMethodsReactivateCurrentCmd)
}

var PaymentMethodsReactivateCurrentCmd = &cobra.Command{
  Use: "reactivate-current",
  Short: TR("Activate payment method"),
  Long: TR(`エラーのある現在の支払い方法を有効化します。`),
  RunE: func(cmd *cobra.Command, args []string) error {
    opt := &apiClientOptions{
      Endpoint: getSpecifiedEndpoint(),
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
    
    param, err := collectPaymentMethodsReactivateCurrentCmdParams()
    if err != nil {
      return err
    }

    result, err := ac.callAPI(param)
    if err != nil {
      cmd.SilenceUsage = true
      return err
    }

    if result != "" {
      return prettyPrintStringAsJSON(result)
    } else {
      return nil
    }
  },
}

func collectPaymentMethodsReactivateCurrentCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "POST",
    path: buildPathForPaymentMethodsReactivateCurrentCmd("/payment_methods/current/activate"),
    query: buildQueryForPaymentMethodsReactivateCurrentCmd(),
    
    
  }, nil
}

func buildPathForPaymentMethodsReactivateCurrentCmd(path string) string {
  
  
  
  
  return path
}

func buildQueryForPaymentMethodsReactivateCurrentCmd() string {
  result := []string{}
  

  

  

  

  return strings.Join(result, "&")
}


