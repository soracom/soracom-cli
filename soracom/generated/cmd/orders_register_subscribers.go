package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var OrdersRegisterSubscribersCmdOrderId string





func init() {
  OrdersRegisterSubscribersCmd.Flags().StringVar(&OrdersRegisterSubscribersCmdOrderId, "order-id", "", "order_id")




  OrdersCmd.AddCommand(OrdersRegisterSubscribersCmd)
}

var OrdersRegisterSubscribersCmd = &cobra.Command{
  Use: "register-subscribers",
  Short: TR("Register subscribers for operator."),
  Long: TR(`発注したSIMをオペレーターに登録します。`),
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
    
    param, err := collectOrdersRegisterSubscribersCmdParams()
    if err != nil {
      return err
    }

    result, err := ac.callAPI(param)
    if err != nil {
      cmd.SilenceUsage = true
      return err
    }

    return prettyPrintJSON(result)
  },
}

func collectOrdersRegisterSubscribersCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "POST",
    path: buildPathForOrdersRegisterSubscribersCmd("/orders/{order_id}/subscribers/register"),
    query: buildQueryForOrdersRegisterSubscribersCmd(),
    
    
  }, nil
}

func buildPathForOrdersRegisterSubscribersCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "order_id" + "}", OrdersRegisterSubscribersCmdOrderId, -1)
  
  
  
  
  return path
}

func buildQueryForOrdersRegisterSubscribersCmd() string {
  result := []string{}
  
  
  

  

  
  return strings.Join(result, "&")
}


