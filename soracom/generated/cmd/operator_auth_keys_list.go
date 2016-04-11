package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var OperatorAuthKeysListCmdOperatorId string






func init() {
  OperatorAuthKeysListCmd.Flags().StringVar(&OperatorAuthKeysListCmdOperatorId, "operator-id", "", "operator_id")




  OperatorAuthKeysCmd.AddCommand(OperatorAuthKeysListCmd)
}

var OperatorAuthKeysListCmd = &cobra.Command{
  Use: "list",
  Short: TR("List Operator AuthKeys"),
  Long: TR(`OperatorのAuthKey一覧を返す。`),
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
    
    param, err := collectOperatorAuthKeysListCmdParams()
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

func collectOperatorAuthKeysListCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "GET",
    path: buildPathForOperatorAuthKeysListCmd("/operators/{operator_id}/auth_keys"),
    query: buildQueryForOperatorAuthKeysListCmd(),
    
    
  }, nil
}

func buildPathForOperatorAuthKeysListCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "operator_id" + "}", OperatorAuthKeysListCmdOperatorId, -1)
  
  
  
  
  
  return path
}

func buildQueryForOperatorAuthKeysListCmd() string {
  result := []string{}
  
  
  

  

  

  

  return strings.Join(result, "&")
}


