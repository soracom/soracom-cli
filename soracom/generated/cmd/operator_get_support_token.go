package cmd

import (

  "fmt"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var OperatorGetSupportTokenCmdOperatorId string





func init() {
  OperatorGetSupportTokenCmd.Flags().StringVar(&OperatorGetSupportTokenCmdOperatorId, "operator-id", "", "operator ID")




  OperatorCmd.AddCommand(OperatorGetSupportTokenCmd)
}

var OperatorGetSupportTokenCmd = &cobra.Command{
  Use: "get-support-token",
  Short: TR("Generate Token for Support Console"),
  Long: TR(`サポート画面アクセス用のトークンを返す。
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
    
    param, err := collectOperatorGetSupportTokenCmdParams()
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

func collectOperatorGetSupportTokenCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "POST",
    path: buildPathForOperatorGetSupportTokenCmd("/operators/{operator_id}/support/token"),
    query: buildQueryForOperatorGetSupportTokenCmd(),
    
    
  }, nil
}

func buildPathForOperatorGetSupportTokenCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "operator_id" + "}", OperatorGetSupportTokenCmdOperatorId, -1)
  
  
  
  
  return path
}

func buildQueryForOperatorGetSupportTokenCmd() string {
  result := []string{}
  
  
  

  

  
  return strings.Join(result, "&")
}

