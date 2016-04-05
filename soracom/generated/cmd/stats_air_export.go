package cmd

import (

  "encoding/json"
  "io/ioutil"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var StatsAirExportCmdOperatorId string

var StatsAirExportCmdPeriod string


var StatsAirExportCmdFrom int64

var StatsAirExportCmdTo int64



var StatsAirExportCmdBody string


func init() {
  StatsAirExportCmd.Flags().StringVar(&StatsAirExportCmdOperatorId, "operator-id", "", "operator ID")

  StatsAirExportCmd.Flags().StringVar(&StatsAirExportCmdPeriod, "period", "", "")

  StatsAirExportCmd.Flags().Int64Var(&StatsAirExportCmdFrom, "from", 0, "")

  StatsAirExportCmd.Flags().Int64Var(&StatsAirExportCmdTo, "to", 0, "")



  StatsAirExportCmd.Flags().StringVar(&StatsAirExportCmdBody, "body", "", TR("cli.common_params.body.short_help"))


  StatsAirCmd.AddCommand(StatsAirExportCmd)
}

var StatsAirExportCmd = &cobra.Command{
  Use: "export",
  Short: TR("Export Air Usage Report of All Subscribers"),
  Long: TR(`Operator が保有する全 Subscriber の通信量をファイルで取得する。
取得対象の期間は from, to のunixtimeで指定する。
履歴の詳細度は月単位。
ファイルの出力先は AWS S3。
ファイルの出力形式は CSV。
`),
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
    
    param, err := collectStatsAirExportCmdParams()
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

func collectStatsAirExportCmdParams() (*apiParams, error) {
  
  body, err := buildBodyForStatsAirExportCmd()
  if err != nil {
    return nil, err
  }
  

  return &apiParams{
    method: "POST",
    path: buildPathForStatsAirExportCmd("/stats/air/operators/{operator_id}/export"),
    query: buildQueryForStatsAirExportCmd(),
    contentType: "application/json",
    body: body,
  }, nil
}

func buildPathForStatsAirExportCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "operator_id" + "}", StatsAirExportCmdOperatorId, -1)
  
  
  
  
  
  
  
  
  
  
  return path
}

func buildQueryForStatsAirExportCmd() string {
  result := []string{}
  
  
  
  
  

  
  
  
  
  

  
  return strings.Join(result, "&")
}


func buildBodyForStatsAirExportCmd() (string, error) {
  if StatsAirExportCmdBody != "" {
    if strings.HasPrefix(StatsAirExportCmdBody, "@") {
      fname := strings.TrimPrefix(StatsAirExportCmdBody, "@")
      bytes, err := ioutil.ReadFile(fname)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else if StatsAirExportCmdBody == "-" {
      bytes, err := ioutil.ReadAll(os.Stdin)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else {
      return StatsAirExportCmdBody, nil
    }
  }

  result := map[string]interface{}{}
  
  if StatsAirExportCmdOperatorId != "" {
    result["operator_id"] = StatsAirExportCmdOperatorId
  }
  
  if StatsAirExportCmdPeriod != "" {
    result["period"] = StatsAirExportCmdPeriod
  }
  
  
  if StatsAirExportCmdFrom != 0 {
    result["from"] = StatsAirExportCmdFrom
  }
  
  if StatsAirExportCmdTo != 0 {
    result["to"] = StatsAirExportCmdTo
  }
  
  

  resultBytes, err := json.Marshal(result)
  if err != nil {
    return "", err
  }
  return string(resultBytes), nil
}

