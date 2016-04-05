package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var UsersPasswordConfiguredCmdOperatorId string

var UsersPasswordConfiguredCmdUserName string





func init() {
  UsersPasswordConfiguredCmd.Flags().StringVar(&UsersPasswordConfiguredCmdOperatorId, "operator-id", "", "operator_id")

  UsersPasswordConfiguredCmd.Flags().StringVar(&UsersPasswordConfiguredCmdUserName, "user-name", "", "user_name")




  UsersPasswordCmd.AddCommand(UsersPasswordConfiguredCmd)
}

var UsersPasswordConfiguredCmd = &cobra.Command{
  Use: "configured",
  Short: TR("Has User Password"),
  Long: TR(`SAMユーザーのパスワードがセットされているかを取得する。`),
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
    
    param, err := collectUsersPasswordConfiguredCmdParams()
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

func collectUsersPasswordConfiguredCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "GET",
    path: buildPathForUsersPasswordConfiguredCmd("/operators/{operator_id}/users/{user_name}/password"),
    query: buildQueryForUsersPasswordConfiguredCmd(),
    
    
  }, nil
}

func buildPathForUsersPasswordConfiguredCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "operator_id" + "}", UsersPasswordConfiguredCmdOperatorId, -1)
  
  
  
  path = strings.Replace(path, "{" + "user_name" + "}", UsersPasswordConfiguredCmdUserName, -1)
  
  
  
  
  return path
}

func buildQueryForUsersPasswordConfiguredCmd() string {
  result := []string{}
  
  
  
  
  

  

  
  return strings.Join(result, "&")
}


