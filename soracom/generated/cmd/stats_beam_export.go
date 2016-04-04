package cmd

import (

  "encoding/json"

  "fmt"

  "io/ioutil"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var StatsBeamExportCmdOperatorId string

var StatsBeamExportCmdPeriod string




var StatsBeamExportCmdBody string


func init() {
  StatsBeamExportCmd.Flags().StringVar(&StatsBeamExportCmdOperatorId, "operator-id", "", "operator ID")

  StatsBeamExportCmd.Flags().StringVar(&StatsBeamExportCmdPeriod, "period", "", "")



  StatsBeamExportCmd.Flags().StringVar(&StatsBeamExportCmdBody, "body", "", TR("cli.common_params.body.short_help"))


  StatsBeamCmd.AddCommand(StatsBeamExportCmd)
}

var StatsBeamExportCmd = &cobra.Command{
  Use: "export",
  Short: TR("Export Beam Usage Report of All Subscribers"),
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
    
    param, err := collectStatsBeamExportCmdParams()
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

func collectStatsBeamExportCmdParams() (*apiParams, error) {
  
  body, err := buildBodyForStatsBeamExportCmd()
  if err != nil {
    return nil, err
  }
  

  return &apiParams{
    method: "POST",
    path: buildPathForStatsBeamExportCmd("/stats/beam/operators/{operator_id}/export"),
    query: buildQueryForStatsBeamExportCmd(),
    contentType: "application/json",
    body: body,
  }, nil
}

func buildPathForStatsBeamExportCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "operator_id" + "}", StatsBeamExportCmdOperatorId, -1)
  
  
  
  
  
  
  return path
}

func buildQueryForStatsBeamExportCmd() string {
  result := []string{}
  
  
  
  
  

  

  
  return strings.Join(result, "&")
}


func buildBodyForStatsBeamExportCmd() (string, error) {
  if StatsBeamExportCmdBody != "" {
    if strings.HasPrefix(StatsBeamExportCmdBody, "@") {
      fname := strings.TrimPrefix(StatsBeamExportCmdBody, "@")
      bytes, err := ioutil.ReadFile(fname)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else if StatsBeamExportCmdBody == "-" {
      bytes, err := ioutil.ReadAll(os.Stdin)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else {
      return StatsBeamExportCmdBody, nil
    }
  }

  result := map[string]interface{}{}
  
  if StatsBeamExportCmdOperatorId != "" {
    result["operator_id"] = StatsBeamExportCmdOperatorId
  }
  
  if StatsBeamExportCmdPeriod != "" {
    result["period"] = StatsBeamExportCmdPeriod
  }
  
  
  

  resultBytes, err := json.Marshal(result)
  if err != nil {
    return "", err
  }
  return string(resultBytes), nil
}

