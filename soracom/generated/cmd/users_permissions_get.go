package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var UsersPermissionsGetCmdOperatorId string

var UsersPermissionsGetCmdUserName string





func init() {
  UsersPermissionsGetCmd.Flags().StringVar(&UsersPermissionsGetCmdOperatorId, "operator-id", "", "operator_id")

  UsersPermissionsGetCmd.Flags().StringVar(&UsersPermissionsGetCmdUserName, "user-name", "", "user_name")




  UsersPermissionsCmd.AddCommand(UsersPermissionsGetCmd)
}

var UsersPermissionsGetCmd = &cobra.Command{
  Use: "get",
  Short: TR("Get User Permission"),
  Long: TR(`SAMユーザーの権限設定を取得する。`),
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
    
    param, err := collectUsersPermissionsGetCmdParams()
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

func collectUsersPermissionsGetCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "GET",
    path: buildPathForUsersPermissionsGetCmd("/operators/{operator_id}/users/{user_name}/permission"),
    query: buildQueryForUsersPermissionsGetCmd(),
    
    
  }, nil
}

func buildPathForUsersPermissionsGetCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "operator_id" + "}", UsersPermissionsGetCmdOperatorId, -1)
  
  
  
  path = strings.Replace(path, "{" + "user_name" + "}", UsersPermissionsGetCmdUserName, -1)
  
  
  
  
  return path
}

func buildQueryForUsersPermissionsGetCmd() string {
  result := []string{}
  
  
  
  
  

  

  
  return strings.Join(result, "&")
}


