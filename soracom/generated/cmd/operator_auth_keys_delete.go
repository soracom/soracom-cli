package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var OperatorAuthKeysDeleteCmdAuthKeyId string

var OperatorAuthKeysDeleteCmdOperatorId string





func init() {
  OperatorAuthKeysDeleteCmd.Flags().StringVar(&OperatorAuthKeysDeleteCmdAuthKeyId, "auth-key-id", "", "auth_key_id")

  OperatorAuthKeysDeleteCmd.Flags().StringVar(&OperatorAuthKeysDeleteCmdOperatorId, "operator-id", "", "operator_id")




  OperatorAuthKeysCmd.AddCommand(OperatorAuthKeysDeleteCmd)
}

var OperatorAuthKeysDeleteCmd = &cobra.Command{
  Use: "delete",
  Short: TR("Delete Operator AuthKey"),
  Long: TR(`OperatorのAuthKeyを削除する。`),
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
    
    param, err := collectOperatorAuthKeysDeleteCmdParams()
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

func collectOperatorAuthKeysDeleteCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "DELETE",
    path: buildPathForOperatorAuthKeysDeleteCmd("/operators/{operator_id}/auth_keys/{auth_key_id}"),
    query: buildQueryForOperatorAuthKeysDeleteCmd(),
    
    
  }, nil
}

func buildPathForOperatorAuthKeysDeleteCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "auth_key_id" + "}", OperatorAuthKeysDeleteCmdAuthKeyId, -1)
  
  
  
  path = strings.Replace(path, "{" + "operator_id" + "}", OperatorAuthKeysDeleteCmdOperatorId, -1)
  
  
  
  
  return path
}

func buildQueryForOperatorAuthKeysDeleteCmd() string {
  result := []string{}
  
  
  
  
  

  

  
  return strings.Join(result, "&")
}


