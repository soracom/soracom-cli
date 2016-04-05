package cmd

import (

  "fmt"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var RolesGetCmdOperatorId string

var RolesGetCmdRoleId string





func init() {
  RolesGetCmd.Flags().StringVar(&RolesGetCmdOperatorId, "operator-id", "", "operator_id")

  RolesGetCmd.Flags().StringVar(&RolesGetCmdRoleId, "role-id", "", "role_id")




  RolesCmd.AddCommand(RolesGetCmd)
}

var RolesGetCmd = &cobra.Command{
  Use: "get",
  Short: TR("Get Role"),
  Long: TR(`Roleを取得する。`),
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
    
    param, err := collectRolesGetCmdParams()
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

func collectRolesGetCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "GET",
    path: buildPathForRolesGetCmd("/operators/{operator_id}/roles/{role_id}"),
    query: buildQueryForRolesGetCmd(),
    
    
  }, nil
}

func buildPathForRolesGetCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "operator_id" + "}", RolesGetCmdOperatorId, -1)
  
  
  
  path = strings.Replace(path, "{" + "role_id" + "}", RolesGetCmdRoleId, -1)
  
  
  
  
  return path
}

func buildQueryForRolesGetCmd() string {
  result := []string{}
  
  
  
  
  

  

  
  return strings.Join(result, "&")
}

