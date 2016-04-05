package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var BillsGetCmdYyyyMM string





func init() {
  BillsGetCmd.Flags().StringVar(&BillsGetCmdYyyyMM, "yyyy-mm", "", "year and month")




  BillsCmd.AddCommand(BillsGetCmd)
}

var BillsGetCmd = &cobra.Command{
  Use: "get",
  Short: TR("Get bill"),
  Long: TR(`指定した月の利用額履歴（無料利用枠などの割引適用後、税込）を返します。このAPIは月末締めをして確定した利用額のみ返します。`),
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
    
    param, err := collectBillsGetCmdParams()
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

func collectBillsGetCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "GET",
    path: buildPathForBillsGetCmd("/bills/{yyyyMM}"),
    query: buildQueryForBillsGetCmd(),
    
    
  }, nil
}

func buildPathForBillsGetCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "yyyyMM" + "}", BillsGetCmdYyyyMM, -1)
  
  
  
  
  return path
}

func buildQueryForBillsGetCmd() string {
  result := []string{}
  
  
  

  

  
  return strings.Join(result, "&")
}


