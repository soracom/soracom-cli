package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var GroupsDeleteCmdGroupId string





func init() {
  GroupsDeleteCmd.Flags().StringVar(&GroupsDeleteCmdGroupId, "group-id", "", "対象のGroupのID")




  GroupsCmd.AddCommand(GroupsDeleteCmd)
}

var GroupsDeleteCmd = &cobra.Command{
  Use: "delete",
  Short: TR("Delete Group"),
  Long: TR(`Group IDで指定されたGroupを削除する`),
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
    
    param, err := collectGroupsDeleteCmdParams()
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

func collectGroupsDeleteCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "DELETE",
    path: buildPathForGroupsDeleteCmd("/groups/{group_id}"),
    query: buildQueryForGroupsDeleteCmd(),
    
    
  }, nil
}

func buildPathForGroupsDeleteCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "group_id" + "}", GroupsDeleteCmdGroupId, -1)
  
  
  
  
  return path
}

func buildQueryForGroupsDeleteCmd() string {
  result := []string{}
  
  
  

  

  
  return strings.Join(result, "&")
}


