package cmd

import (

  "fmt"

  "os"
  "strings"

  "github.com/spf13/cobra"
)










func init() {



  PaymentMethodsCmd.AddCommand(PaymentMethodsGetCurrentCmd)
}

var PaymentMethodsGetCurrentCmd = &cobra.Command{
  Use: "get-current",
  Short: TR("Get payment method information"),
  Long: TR(`現在の支払い方法を返します。propertiesに詳細情報が入っています`),
  RunE: func(cmd *cobra.Command, args []string) error {
    opt := &apiClientOptions{
      Endpoint: getSpecifiedEndpoint(),
      BasePath: "/v1",
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
    
    param, err := collectPaymentMethodsGetCurrentCmdParams()
    if err != nil {
      return err
    }

    result, err := ac.callAPI(param)
    if err != nil {
      cmd.SilenceUsage = true
      return err
    }

    fmt.Println(result)
    return nil
  },
}

func collectPaymentMethodsGetCurrentCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "GET",
    path: buildPathForPaymentMethodsGetCurrentCmd("/payment_methods/current"),
    query: buildQueryForPaymentMethodsGetCurrentCmd(),
    
    
  }, nil
}

func buildPathForPaymentMethodsGetCurrentCmd(path string) string {
  
  
  
  return path
}

func buildQueryForPaymentMethodsGetCurrentCmd() string {
  result := []string{}
  

  

  
  return strings.Join(result, "&")
}


