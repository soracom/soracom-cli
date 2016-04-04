package cmd

import (

  "fmt"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var OperatorAuthKeysGenerateCmdOperatorId string





func init() {
  OperatorAuthKeysGenerateCmd.Flags().StringVar(&OperatorAuthKeysGenerateCmdOperatorId, "operator-id", "", "operator_id")




  OperatorAuthKeysCmd.AddCommand(OperatorAuthKeysGenerateCmd)
}

var OperatorAuthKeysGenerateCmd = &cobra.Command{
  Use: "generate",
  Short: TR("Generate Operator AuthKey"),
  Long: TR(`OperatorのAuthKeyを生成する。`),
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
    
    param, err := collectOperatorAuthKeysGenerateCmdParams()
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

func collectOperatorAuthKeysGenerateCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "POST",
    path: buildPathForOperatorAuthKeysGenerateCmd("/operators/{operator_id}/auth_keys"),
    query: buildQueryForOperatorAuthKeysGenerateCmd(),
    
    
  }, nil
}

func buildPathForOperatorAuthKeysGenerateCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "operator_id" + "}", OperatorAuthKeysGenerateCmdOperatorId, -1)
  
  
  
  
  return path
}

func buildQueryForOperatorAuthKeysGenerateCmd() string {
  result := []string{}
  
  
  

  

  
  return strings.Join(result, "&")
}


