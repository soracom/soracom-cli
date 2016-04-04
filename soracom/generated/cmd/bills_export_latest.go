package cmd

import (

  "fmt"

  "os"
  "strings"

  "github.com/spf13/cobra"
)










func init() {



  BillsCmd.AddCommand(BillsExportLatestCmd)
}

var BillsExportLatestCmd = &cobra.Command{
  Use: "export-latest",
  Short: TR("Output latest billing CSV file to S3"),
  Long: TR(`直近月の利用額明細を返します。この明細には、日ごと,Subscrierごと,課金項目ごとの利用額が含まれます。このAPIで取得した金額は請求確定前の金額になります。`),
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
    
    param, err := collectBillsExportLatestCmdParams()
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

func collectBillsExportLatestCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "POST",
    path: buildPathForBillsExportLatestCmd("/bills/latest/export"),
    query: buildQueryForBillsExportLatestCmd(),
    
    
  }, nil
}

func buildPathForBillsExportLatestCmd(path string) string {
  
  
  
  return path
}

func buildQueryForBillsExportLatestCmd() string {
  result := []string{}
  

  

  
  return strings.Join(result, "&")
}


