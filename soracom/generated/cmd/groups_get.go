package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var GroupsGetCmdGroupId string






func init() {
  GroupsGetCmd.Flags().StringVar(&GroupsGetCmdGroupId, "group-id", "", "対象のGroupのID")




  GroupsCmd.AddCommand(GroupsGetCmd)
}

var GroupsGetCmd = &cobra.Command{
  Use: "get",
  Short: TR("Get Group"),
  Long: TR(`Group IDで指定されたGroupを返す
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
    
    param, err := collectGroupsGetCmdParams()
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

func collectGroupsGetCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "GET",
    path: buildPathForGroupsGetCmd("/groups/{group_id}"),
    query: buildQueryForGroupsGetCmd(),
    
    
  }, nil
}

func buildPathForGroupsGetCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "group_id" + "}", GroupsGetCmdGroupId, -1)
  
  
  
  
  
  return path
}

func buildQueryForGroupsGetCmd() string {
  result := []string{}
  
  
  

  

  

  

  return strings.Join(result, "&")
}


