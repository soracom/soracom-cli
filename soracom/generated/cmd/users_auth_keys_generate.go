package cmd

import (

  "fmt"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var UsersAuthKeysGenerateCmdOperatorId string

var UsersAuthKeysGenerateCmdUserName string





func init() {
  UsersAuthKeysGenerateCmd.Flags().StringVar(&UsersAuthKeysGenerateCmdOperatorId, "operator-id", "", "operator_id")

  UsersAuthKeysGenerateCmd.Flags().StringVar(&UsersAuthKeysGenerateCmdUserName, "user-name", "", "user_name")




  UsersAuthKeysCmd.AddCommand(UsersAuthKeysGenerateCmd)
}

var UsersAuthKeysGenerateCmd = &cobra.Command{
  Use: "generate",
  Short: TR("Generate AuthKey"),
  Long: TR(`SAMユーザーのAuthKeyを生成する。`),
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
    
    param, err := collectUsersAuthKeysGenerateCmdParams()
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

func collectUsersAuthKeysGenerateCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "POST",
    path: buildPathForUsersAuthKeysGenerateCmd("/operators/{operator_id}/users/{user_name}/auth_keys"),
    query: buildQueryForUsersAuthKeysGenerateCmd(),
    
    
  }, nil
}

func buildPathForUsersAuthKeysGenerateCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "operator_id" + "}", UsersAuthKeysGenerateCmdOperatorId, -1)
  
  
  
  path = strings.Replace(path, "{" + "user_name" + "}", UsersAuthKeysGenerateCmdUserName, -1)
  
  
  
  
  return path
}

func buildQueryForUsersAuthKeysGenerateCmd() string {
  result := []string{}
  
  
  
  
  

  

  
  return strings.Join(result, "&")
}

