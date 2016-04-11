package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var PaymentHistoryGetCmdPaymentTransactionId string






func init() {
  PaymentHistoryGetCmd.Flags().StringVar(&PaymentHistoryGetCmdPaymentTransactionId, "payment-transaction-id", "", "payment_transaction_id")




  PaymentHistoryCmd.AddCommand(PaymentHistoryGetCmd)
}

var PaymentHistoryGetCmd = &cobra.Command{
  Use: "get",
  Short: TR("Get payment transaction result"),
  Long: TR(`課金処理の結果を返します`),
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
    
    param, err := collectPaymentHistoryGetCmdParams()
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

func collectPaymentHistoryGetCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "GET",
    path: buildPathForPaymentHistoryGetCmd("/payment_history/transactions/{payment_transaction_id}"),
    query: buildQueryForPaymentHistoryGetCmd(),
    
    
  }, nil
}

func buildPathForPaymentHistoryGetCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "payment_transaction_id" + "}", PaymentHistoryGetCmdPaymentTransactionId, -1)
  
  
  
  
  
  return path
}

func buildQueryForPaymentHistoryGetCmd() string {
  result := []string{}
  
  
  

  

  

  

  return strings.Join(result, "&")
}


