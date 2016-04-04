package cmd

import (

  "fmt"

  "os"
  "strings"

  "github.com/spf13/cobra"
)










func init() {



  BillsCmd.AddCommand(BillsListCmd)
}

var BillsListCmd = &cobra.Command{
  Use: "list",
  Short: TR("Get billing history"),
  Long: TR(`過去の利用額履歴（無料利用枠などの割引適用後、税込）を返します。このAPIは月末締めをして確定した利用額のみ返します。`),
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
    
    param, err := collectBillsListCmdParams()
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

func collectBillsListCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "GET",
    path: buildPathForBillsListCmd("/bills"),
    query: buildQueryForBillsListCmd(),
    
    
  }, nil
}

func buildPathForBillsListCmd(path string) string {
  
  
  
  return path
}

func buildQueryForBillsListCmd() string {
  result := []string{}
  

  

  
  return strings.Join(result, "&")
}


