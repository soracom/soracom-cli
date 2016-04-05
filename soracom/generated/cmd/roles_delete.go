package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var RolesDeleteCmdOperatorId string

var RolesDeleteCmdRoleId string





func init() {
  RolesDeleteCmd.Flags().StringVar(&RolesDeleteCmdOperatorId, "operator-id", "", "operator_id")

  RolesDeleteCmd.Flags().StringVar(&RolesDeleteCmdRoleId, "role-id", "", "role_id")




  RolesCmd.AddCommand(RolesDeleteCmd)
}

var RolesDeleteCmd = &cobra.Command{
  Use: "delete",
  Short: TR("Delete Role"),
  Long: TR(`Roleを削除する。`),
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
    
    param, err := collectRolesDeleteCmdParams()
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

func collectRolesDeleteCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "DELETE",
    path: buildPathForRolesDeleteCmd("/operators/{operator_id}/roles/{role_id}"),
    query: buildQueryForRolesDeleteCmd(),
    
    
  }, nil
}

func buildPathForRolesDeleteCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "operator_id" + "}", RolesDeleteCmdOperatorId, -1)
  
  
  
  path = strings.Replace(path, "{" + "role_id" + "}", RolesDeleteCmdRoleId, -1)
  
  
  
  
  return path
}

func buildQueryForRolesDeleteCmd() string {
  result := []string{}
  
  
  
  
  

  

  
  return strings.Join(result, "&")
}


