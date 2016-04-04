package cmd

import (

  "fmt"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var OperatorGetCmdOperatorId string





func init() {
  OperatorGetCmd.Flags().StringVar(&OperatorGetCmdOperatorId, "operator-id", "", "operator ID")




  OperatorCmd.AddCommand(OperatorGetCmd)
}

var OperatorGetCmd = &cobra.Command{
  Use: "get",
  Short: TR("Get Operator"),
  Long: TR(`Operatorの情報を返す。
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
    
    param, err := collectOperatorGetCmdParams()
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

func collectOperatorGetCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "GET",
    path: buildPathForOperatorGetCmd("/operators/{operator_id}"),
    query: buildQueryForOperatorGetCmd(),
    
    
  }, nil
}

func buildPathForOperatorGetCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "operator_id" + "}", OperatorGetCmdOperatorId, -1)
  
  
  
  
  return path
}

func buildQueryForOperatorGetCmd() string {
  result := []string{}
  
  
  

  

  
  return strings.Join(result, "&")
}


