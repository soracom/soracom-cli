package cmd

import (

  "fmt"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var BillsExportCmdYyyyMM string





func init() {
  BillsExportCmd.Flags().StringVar(&BillsExportCmdYyyyMM, "yyyy-mm", "", "yyyyMM")




  BillsCmd.AddCommand(BillsExportCmd)
}

var BillsExportCmd = &cobra.Command{
  Use: "export",
  Short: TR("Output billing CSV file to S3"),
  Long: TR(`指定月の利用額明細を返します。この明細には、日ごと,Subscrierごと,課金項目ごとの利用額が含まれます。`),
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
    
    param, err := collectBillsExportCmdParams()
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

func collectBillsExportCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "POST",
    path: buildPathForBillsExportCmd("/bills/{yyyyMM}/export"),
    query: buildQueryForBillsExportCmd(),
    
    
  }, nil
}

func buildPathForBillsExportCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "yyyyMM" + "}", BillsExportCmdYyyyMM, -1)
  
  
  
  
  return path
}

func buildQueryForBillsExportCmd() string {
  result := []string{}
  
  
  

  

  
  return strings.Join(result, "&")
}


