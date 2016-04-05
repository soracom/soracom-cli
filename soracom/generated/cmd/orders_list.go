package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)










func init() {



  OrdersCmd.AddCommand(OrdersListCmd)
}

var OrdersListCmd = &cobra.Command{
  Use: "list",
  Short: TR("List confirmed orders."),
  Long: TR(`発注確定済みの発注一覧を返します。`),
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
    
    param, err := collectOrdersListCmdParams()
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

func collectOrdersListCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "GET",
    path: buildPathForOrdersListCmd("/orders"),
    query: buildQueryForOrdersListCmd(),
    
    
  }, nil
}

func buildPathForOrdersListCmd(path string) string {
  
  
  
  return path
}

func buildQueryForOrdersListCmd() string {
  result := []string{}
  

  

  
  return strings.Join(result, "&")
}


