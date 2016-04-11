package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var BillsGetDailyCmdYyyyMM string






func init() {
  BillsGetDailyCmd.Flags().StringVar(&BillsGetDailyCmdYyyyMM, "yyyy-mm", "", "year and month")




  BillsCmd.AddCommand(BillsGetDailyCmd)
}

var BillsGetDailyCmd = &cobra.Command{
  Use: "get-daily",
  Short: TR("Get bill per day"),
  Long: TR(`指定した月の、日ごとの利用額明細を返します。このAPIは確定した利用額のみ返します。`),
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
    
    param, err := collectBillsGetDailyCmdParams()
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

func collectBillsGetDailyCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "GET",
    path: buildPathForBillsGetDailyCmd("/bills/{yyyyMM}/daily"),
    query: buildQueryForBillsGetDailyCmd(),
    
    
  }, nil
}

func buildPathForBillsGetDailyCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "yyyyMM" + "}", BillsGetDailyCmdYyyyMM, -1)
  
  
  
  
  
  return path
}

func buildQueryForBillsGetDailyCmd() string {
  result := []string{}
  
  
  

  

  

  

  return strings.Join(result, "&")
}


