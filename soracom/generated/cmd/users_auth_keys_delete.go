package cmd

import (

  "fmt"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var UsersAuthKeysDeleteCmdAuthKeyId string

var UsersAuthKeysDeleteCmdOperatorId string

var UsersAuthKeysDeleteCmdUserName string





func init() {
  UsersAuthKeysDeleteCmd.Flags().StringVar(&UsersAuthKeysDeleteCmdAuthKeyId, "auth-key-id", "", "auth_key_id")

  UsersAuthKeysDeleteCmd.Flags().StringVar(&UsersAuthKeysDeleteCmdOperatorId, "operator-id", "", "operator_id")

  UsersAuthKeysDeleteCmd.Flags().StringVar(&UsersAuthKeysDeleteCmdUserName, "user-name", "", "user_name")




  UsersAuthKeysCmd.AddCommand(UsersAuthKeysDeleteCmd)
}

var UsersAuthKeysDeleteCmd = &cobra.Command{
  Use: "delete",
  Short: TR("Delete User AuthKey"),
  Long: TR(`SAMユーザーのAuthKeyを削除する。`),
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
    
    param, err := collectUsersAuthKeysDeleteCmdParams()
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

func collectUsersAuthKeysDeleteCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "DELETE",
    path: buildPathForUsersAuthKeysDeleteCmd("/operators/{operator_id}/users/{user_name}/auth_keys/{auth_key_id}"),
    query: buildQueryForUsersAuthKeysDeleteCmd(),
    
    
  }, nil
}

func buildPathForUsersAuthKeysDeleteCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "auth_key_id" + "}", UsersAuthKeysDeleteCmdAuthKeyId, -1)
  
  
  
  path = strings.Replace(path, "{" + "operator_id" + "}", UsersAuthKeysDeleteCmdOperatorId, -1)
  
  
  
  path = strings.Replace(path, "{" + "user_name" + "}", UsersAuthKeysDeleteCmdUserName, -1)
  
  
  
  
  return path
}

func buildQueryForUsersAuthKeysDeleteCmd() string {
  result := []string{}
  
  
  
  
  
  
  

  

  
  return strings.Join(result, "&")
}


