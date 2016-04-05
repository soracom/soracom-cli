package cmd

import (

  "fmt"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var UsersAuthKeysListCmdOperatorId string

var UsersAuthKeysListCmdUserName string





func init() {
  UsersAuthKeysListCmd.Flags().StringVar(&UsersAuthKeysListCmdOperatorId, "operator-id", "", "operator_id")

  UsersAuthKeysListCmd.Flags().StringVar(&UsersAuthKeysListCmdUserName, "user-name", "", "user_name")




  UsersAuthKeysCmd.AddCommand(UsersAuthKeysListCmd)
}

var UsersAuthKeysListCmd = &cobra.Command{
  Use: "list",
  Short: TR("List User AuthKeys"),
  Long: TR(`SAMユーザーのAuthKey一覧を返す。`),
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
    
    param, err := collectUsersAuthKeysListCmdParams()
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

func collectUsersAuthKeysListCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "GET",
    path: buildPathForUsersAuthKeysListCmd("/operators/{operator_id}/users/{user_name}/auth_keys"),
    query: buildQueryForUsersAuthKeysListCmd(),
    
    
  }, nil
}

func buildPathForUsersAuthKeysListCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "operator_id" + "}", UsersAuthKeysListCmdOperatorId, -1)
  
  
  
  path = strings.Replace(path, "{" + "user_name" + "}", UsersAuthKeysListCmdUserName, -1)
  
  
  
  
  return path
}

func buildQueryForUsersAuthKeysListCmd() string {
  result := []string{}
  
  
  
  
  

  

  
  return strings.Join(result, "&")
}

