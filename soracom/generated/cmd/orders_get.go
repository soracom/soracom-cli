package cmd

import (

  "fmt"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var OrdersGetCmdOrderId string





func init() {
  OrdersGetCmd.Flags().StringVar(&OrdersGetCmdOrderId, "order-id", "", "order_id")




  OrdersCmd.AddCommand(OrdersGetCmd)
}

var OrdersGetCmd = &cobra.Command{
  Use: "get",
  Short: TR("Get confirmed order."),
  Long: TR(`発注確定済みの発注を返します。`),
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
    
    param, err := collectOrdersGetCmdParams()
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

func collectOrdersGetCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "GET",
    path: buildPathForOrdersGetCmd("/orders/{order_id}"),
    query: buildQueryForOrdersGetCmd(),
    
    
  }, nil
}

func buildPathForOrdersGetCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "order_id" + "}", OrdersGetCmdOrderId, -1)
  
  
  
  
  return path
}

func buildQueryForOrdersGetCmd() string {
  result := []string{}
  
  
  

  

  
  return strings.Join(result, "&")
}


