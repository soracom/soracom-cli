package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var CredentialsDeleteCmdOperatorId string






func init() {
  CredentialsDeleteCmd.Flags().StringVar(&CredentialsDeleteCmdOperatorId, "operator-id", "", "Operator ID")




  CredentialsCmd.AddCommand(CredentialsDeleteCmd)
}

var CredentialsDeleteCmd = &cobra.Command{
  Use: "delete",
  Short: TR("Delete Credential"),
  Long: TR(`認証情報を削除する。`),
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
    
    param, err := collectCredentialsDeleteCmdParams()
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

func collectCredentialsDeleteCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "DELETE",
    path: buildPathForCredentialsDeleteCmd("/credentials/{credentials_id}"),
    query: buildQueryForCredentialsDeleteCmd(),
    
    
  }, nil
}

func buildPathForCredentialsDeleteCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "operator_id" + "}", CredentialsDeleteCmdOperatorId, -1)
  
  
  
  
  
  return path
}

func buildQueryForCredentialsDeleteCmd() string {
  result := []string{}
  
  
  

  

  

  

  return strings.Join(result, "&")
}


