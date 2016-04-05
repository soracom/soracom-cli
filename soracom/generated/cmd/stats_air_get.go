package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var StatsAirGetCmdImsi string

var StatsAirGetCmdPeriod string


var StatsAirGetCmdFrom int64

var StatsAirGetCmdTo int64




func init() {
  StatsAirGetCmd.Flags().StringVar(&StatsAirGetCmdImsi, "imsi", "", "imsi")

  StatsAirGetCmd.Flags().StringVar(&StatsAirGetCmdPeriod, "period", "", "集計単位。minutesは5分おき")

  StatsAirGetCmd.Flags().Int64Var(&StatsAirGetCmdFrom, "from", 0, "集計対象時刻の始まりをunixtimeで与える")

  StatsAirGetCmd.Flags().Int64Var(&StatsAirGetCmdTo, "to", 0, "集計対象時刻の終わりをunixtimeで与える")




  StatsAirCmd.AddCommand(StatsAirGetCmd)
}

var StatsAirGetCmd = &cobra.Command{
  Use: "get",
  Short: TR("Get Air Usage Report of Subscriber"),
  Long: TR(`IMSI で指定した Subscriber の通信量履歴を取得する。`),
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
    
    param, err := collectStatsAirGetCmdParams()
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

func collectStatsAirGetCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "GET",
    path: buildPathForStatsAirGetCmd("/stats/air/subscribers/{imsi}"),
    query: buildQueryForStatsAirGetCmd(),
    
    
  }, nil
}

func buildPathForStatsAirGetCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "imsi" + "}", StatsAirGetCmdImsi, -1)
  
  
  
  
  
  
  
  
  
  
  return path
}

func buildQueryForStatsAirGetCmd() string {
  result := []string{}
  
  
  
  
  if StatsAirGetCmdPeriod != "" {
    result = append(result, sprintf("%s=%s", "period", StatsAirGetCmdPeriod))
  }
  
  

  
  
  if StatsAirGetCmdFrom != 0 {
    result = append(result, sprintf("%s=%d", "from", StatsAirGetCmdFrom))
  }
  
  
  
  if StatsAirGetCmdTo != 0 {
    result = append(result, sprintf("%s=%d", "to", StatsAirGetCmdTo))
  }
  
  

  
  return strings.Join(result, "&")
}


