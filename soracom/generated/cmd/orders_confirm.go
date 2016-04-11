package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var OrdersConfirmCmdOrderId string






func init() {
  OrdersConfirmCmd.Flags().StringVar(&OrdersConfirmCmdOrderId, "order-id", "", "order_id")




  OrdersCmd.AddCommand(OrdersConfirmCmd)
}

var OrdersConfirmCmd = &cobra.Command{
  Use: "confirm",
  Short: TR("Confirm order."),
  Long: TR(`与信を実施し、問題ない場合に発注を確定します。`),
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
    
    param, err := collectOrdersConfirmCmdParams()
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

func collectOrdersConfirmCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "PUT",
    path: buildPathForOrdersConfirmCmd("/orders/{order_id}/confirm"),
    query: buildQueryForOrdersConfirmCmd(),
    
    
  }, nil
}

func buildPathForOrdersConfirmCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "order_id" + "}", OrdersConfirmCmdOrderId, -1)
  
  
  
  
  
  return path
}

func buildQueryForOrdersConfirmCmd() string {
  result := []string{}
  
  
  

  

  

  

  return strings.Join(result, "&")
}


