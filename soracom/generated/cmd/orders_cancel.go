package cmd

import (

  "fmt"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var OrdersCancelCmdOrderId string





func init() {
  OrdersCancelCmd.Flags().StringVar(&OrdersCancelCmdOrderId, "order-id", "", "order_id")




  OrdersCmd.AddCommand(OrdersCancelCmd)
}

var OrdersCancelCmd = &cobra.Command{
  Use: "cancel",
  Short: TR("Cancel order."),
  Long: TR(`発注をキャンセルします。既に発送済みの場合はエラーが返されます。`),
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
    
    param, err := collectOrdersCancelCmdParams()
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

func collectOrdersCancelCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "PUT",
    path: buildPathForOrdersCancelCmd("/orders/{order_id}/cancel"),
    query: buildQueryForOrdersCancelCmd(),
    
    
  }, nil
}

func buildPathForOrdersCancelCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "order_id" + "}", OrdersCancelCmdOrderId, -1)
  
  
  
  
  return path
}

func buildQueryForOrdersCancelCmd() string {
  result := []string{}
  
  
  

  

  
  return strings.Join(result, "&")
}

