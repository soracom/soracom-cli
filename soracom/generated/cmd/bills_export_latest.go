package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)











func init() {



  BillsCmd.AddCommand(BillsExportLatestCmd)
}

var BillsExportLatestCmd = &cobra.Command{
  Use: "export-latest",
  Short: TR("bills.export_latest_billing.post.summary"),
  Long: TR(`bills.export_latest_billing.post.description`),
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
    
    param, err := collectBillsExportLatestCmdParams()
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


