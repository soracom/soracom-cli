package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var GroupsDeleteTagCmdGroupId string

var GroupsDeleteTagCmdTagName string





func init() {
  GroupsDeleteTagCmd.Flags().StringVar(&GroupsDeleteTagCmdGroupId, "group-id", "", "対象のGroupのID")

  GroupsDeleteTagCmd.Flags().StringVar(&GroupsDeleteTagCmdTagName, "tag-name", "", "削除対象のタグ名（URL の Path の一部になるので、パーセントエンコーディングを施す。JavaScript なら encodeURIComponent() したものを指定する）")




  GroupsCmd.AddCommand(GroupsDeleteTagCmd)
}

var GroupsDeleteTagCmd = &cobra.Command{
  Use: "delete-tag",
  Short: TR("Delete Group Tag"),
  Long: TR(`指定されたGroupのタグを削除
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
    
    param, err := collectGroupsDeleteTagCmdParams()
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

func collectGroupsDeleteTagCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "DELETE",
    path: buildPathForGroupsDeleteTagCmd("/groups/{group_id}/tags/{tag_name}"),
    query: buildQueryForGroupsDeleteTagCmd(),
    
    
  }, nil
}

func buildPathForGroupsDeleteTagCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "group_id" + "}", GroupsDeleteTagCmdGroupId, -1)
  
  
  
  path = strings.Replace(path, "{" + "tag_name" + "}", GroupsDeleteTagCmdTagName, -1)
  
  
  
  
  return path
}

func buildQueryForGroupsDeleteTagCmd() string {
  result := []string{}
  
  
  
  
  

  

  
  return strings.Join(result, "&")
}


