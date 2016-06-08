package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var UsersDetachRoleCmdOperatorId string

var UsersDetachRoleCmdRoleId string

var UsersDetachRoleCmdUserName string






func init() {
  UsersDetachRoleCmd.Flags().StringVar(&UsersDetachRoleCmdOperatorId, "operator-id", "", TR("operator_id"))

  UsersDetachRoleCmd.Flags().StringVar(&UsersDetachRoleCmdRoleId, "role-id", "", TR("role_id"))

  UsersDetachRoleCmd.Flags().StringVar(&UsersDetachRoleCmdUserName, "user-name", "", TR("user_name"))




  UsersCmd.AddCommand(UsersDetachRoleCmd)
}

var UsersDetachRoleCmd = &cobra.Command{
  Use: "detach-role",
  Short: TR("roles.detach_role.delete.summary"),
  Long: TR(`roles.detach_role.delete.description`),
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
    
    param, err := collectUsersDetachRoleCmdParams()
    if err != nil {
      return err
    }

    result, err := ac.callAPI(param)
    if err != nil {
      cmd.SilenceUsage = true
      return err
    }

    if result != "" {
      return prettyPrintStringAsJSON(result)
    } else {
      return nil
    }
  },
}

func collectUsersDetachRoleCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "DELETE",
    path: buildPathForUsersDetachRoleCmd("/operators/{operator_id}/users/{user_name}/roles/{role_id}"),
    query: buildQueryForUsersDetachRoleCmd(),
    
    
  }, nil
}

func buildPathForUsersDetachRoleCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "operator_id" + "}", UsersDetachRoleCmdOperatorId, -1)
  
  
  
  path = strings.Replace(path, "{" + "role_id" + "}", UsersDetachRoleCmdRoleId, -1)
  
  
  
  path = strings.Replace(path, "{" + "user_name" + "}", UsersDetachRoleCmdUserName, -1)
  
  
  
  
  
  return path
}

func buildQueryForUsersDetachRoleCmd() string {
  result := []string{}
  
  
  
  
  
  
  

  

  

  

  return strings.Join(result, "&")
}


