package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var StatsBeamGetCmdImsi string

var StatsBeamGetCmdPeriod string


var StatsBeamGetCmdFrom int64

var StatsBeamGetCmdTo int64




func init() {
  StatsBeamGetCmd.Flags().StringVar(&StatsBeamGetCmdImsi, "imsi", "", "imsi")

  StatsBeamGetCmd.Flags().StringVar(&StatsBeamGetCmdPeriod, "period", "", "集計単位。minutesは5分おき")

  StatsBeamGetCmd.Flags().Int64Var(&StatsBeamGetCmdFrom, "from", 0, "集計対象時刻の始まりをunixtimeで与える")

  StatsBeamGetCmd.Flags().Int64Var(&StatsBeamGetCmdTo, "to", 0, "集計対象時刻の終わりをunixtimeで与える")




  StatsBeamCmd.AddCommand(StatsBeamGetCmd)
}

var StatsBeamGetCmd = &cobra.Command{
  Use: "get",
  Short: TR("Get Beam Usage Report of Subscriber"),
  Long: TR(`IMSI で指定した Subscriber のSoracom Beam利用量履歴を取得する。`),
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
    
    param, err := collectStatsBeamGetCmdParams()
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

func collectStatsBeamGetCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "GET",
    path: buildPathForStatsBeamGetCmd("/stats/beam/subscribers/{imsi}"),
    query: buildQueryForStatsBeamGetCmd(),
    
    
  }, nil
}

func buildPathForStatsBeamGetCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "imsi" + "}", StatsBeamGetCmdImsi, -1)
  
  
  
  
  
  
  
  
  
  
  return path
}

func buildQueryForStatsBeamGetCmd() string {
  result := []string{}
  
  
  
  
  if StatsBeamGetCmdPeriod != "" {
    result = append(result, sprintf("%s=%s", "period", StatsBeamGetCmdPeriod))
  }
  
  

  
  
  if StatsBeamGetCmdFrom != 0 {
    result = append(result, sprintf("%s=%d", "from", StatsBeamGetCmdFrom))
  }
  
  
  
  if StatsBeamGetCmdTo != 0 {
    result = append(result, sprintf("%s=%d", "to", StatsBeamGetCmdTo))
  }
  
  

  
  return strings.Join(result, "&")
}


