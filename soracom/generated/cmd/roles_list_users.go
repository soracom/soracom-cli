package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var RolesListUsersCmdOperatorId string

var RolesListUsersCmdRoleId string






func init() {
  RolesListUsersCmd.Flags().StringVar(&RolesListUsersCmdOperatorId, "operator-id", "", "operator_id")

  RolesListUsersCmd.Flags().StringVar(&RolesListUsersCmdRoleId, "role-id", "", "role_id")




  RolesCmd.AddCommand(RolesListUsersCmd)
}

var RolesListUsersCmd = &cobra.Command{
  Use: "list-users",
  Short: TR("List Role Attached Users"),
  Long: TR(`Roleに紐づくユーザーの一覧を取得する。`),
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
    
    param, err := collectRolesListUsersCmdParams()
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

func collectRolesListUsersCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "GET",
    path: buildPathForRolesListUsersCmd("/operators/{operator_id}/roles/{role_id}/users"),
    query: buildQueryForRolesListUsersCmd(),
    
    
  }, nil
}

func buildPathForRolesListUsersCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "operator_id" + "}", RolesListUsersCmdOperatorId, -1)
  
  
  
  path = strings.Replace(path, "{" + "role_id" + "}", RolesListUsersCmdRoleId, -1)
  
  
  
  
  
  return path
}

func buildQueryForRolesListUsersCmd() string {
  result := []string{}
  
  
  
  
  

  

  

  

  return strings.Join(result, "&")
}


