package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var GroupsDeleteConfigCmdGroupId string

var GroupsDeleteConfigCmdName string

var GroupsDeleteConfigCmdNamespace string






func init() {
  GroupsDeleteConfigCmd.Flags().StringVar(&GroupsDeleteConfigCmdGroupId, "group-id", "", "対象のGroup")

  GroupsDeleteConfigCmd.Flags().StringVar(&GroupsDeleteConfigCmdName, "name", "", "削除対象のパラメータ名（URL の Path の一部になるので、パーセントエンコーディングを施す。JavaScript なら encodeURIComponent() したものを指定する）")

  GroupsDeleteConfigCmd.Flags().StringVar(&GroupsDeleteConfigCmdNamespace, "namespace", "", "対象のパラメータのネームスペース")




  GroupsCmd.AddCommand(GroupsDeleteConfigCmd)
}

var GroupsDeleteConfigCmd = &cobra.Command{
  Use: "delete-config",
  Short: TR("Delete Group Configuration Parameters"),
  Long: TR(`指定されたGroupのパラメータを削除
`),
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
    
    param, err := collectGroupsDeleteConfigCmdParams()
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

func collectGroupsDeleteConfigCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "DELETE",
    path: buildPathForGroupsDeleteConfigCmd("/groups/{group_id}/configuration/{namespace}/{name}"),
    query: buildQueryForGroupsDeleteConfigCmd(),
    
    
  }, nil
}

func buildPathForGroupsDeleteConfigCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "group_id" + "}", GroupsDeleteConfigCmdGroupId, -1)
  
  
  
  path = strings.Replace(path, "{" + "name" + "}", GroupsDeleteConfigCmdName, -1)
  
  
  
  path = strings.Replace(path, "{" + "namespace" + "}", GroupsDeleteConfigCmdNamespace, -1)
  
  
  
  
  
  return path
}

func buildQueryForGroupsDeleteConfigCmd() string {
  result := []string{}
  
  
  
  
  
  
  

  

  

  

  return strings.Join(result, "&")
}


