package cmd

import (

  "fmt"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var GroupsListSubscribersCmdGroupId string

var GroupsListSubscribersCmdLastEvaluatedKey string


var GroupsListSubscribersCmdLimit int64




func init() {
  GroupsListSubscribersCmd.Flags().StringVar(&GroupsListSubscribersCmdGroupId, "group-id", "", "対象のGroupのID")

  GroupsListSubscribersCmd.Flags().StringVar(&GroupsListSubscribersCmdLastEvaluatedKey, "last-evaluated-key", "", "現ページで取得した最後のSubscriberのIMSI。このパラメータを指定することで次のSubscriber以降のリストを取得できる。")

  GroupsListSubscribersCmd.Flags().Int64Var(&GroupsListSubscribersCmdLimit, "limit", 0, "レスポンス1ページあたりの最大数")




  GroupsCmd.AddCommand(GroupsListSubscribersCmd)
}

var GroupsListSubscribersCmd = &cobra.Command{
  Use: "list-subscribers",
  Short: TR("List Subscribers in a group"),
  Long: TR(`Group IDで指定されたGroupに属するSubscriberの一覧を返す
`),
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
    
    param, err := collectGroupsListSubscribersCmdParams()
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

func collectGroupsListSubscribersCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "GET",
    path: buildPathForGroupsListSubscribersCmd("/groups/{group_id}/subscribers"),
    query: buildQueryForGroupsListSubscribersCmd(),
    
    
  }, nil
}

func buildPathForGroupsListSubscribersCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "group_id" + "}", GroupsListSubscribersCmdGroupId, -1)
  
  
  
  
  
  
  
  
  return path
}

func buildQueryForGroupsListSubscribersCmd() string {
  result := []string{}
  
  
  
  
  if GroupsListSubscribersCmdLastEvaluatedKey != "" {
    result = append(result, fmt.Sprintf("%s=%s", "last_evaluated_key", GroupsListSubscribersCmdLastEvaluatedKey))
  }
  
  

  
  
  if GroupsListSubscribersCmdLimit != 0 {
    result = append(result, fmt.Sprintf("%s=%d", "limit", GroupsListSubscribersCmdLimit))
  }
  
  

  
  return strings.Join(result, "&")
}


