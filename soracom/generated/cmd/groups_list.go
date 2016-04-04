package cmd

import (

  "fmt"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var GroupsListCmdLastEvaluatedKey string

var GroupsListCmdTagName string

var GroupsListCmdTagValue string

var GroupsListCmdTagValueMatchMode string


var GroupsListCmdLimit int64




func init() {
  GroupsListCmd.Flags().StringVar(&GroupsListCmdLastEvaluatedKey, "last-evaluated-key", "", "現ページで取得した最後のGroupのID。このパラメータを指定することで次のGroupから始まるリストを取得できる。")

  GroupsListCmd.Flags().StringVar(&GroupsListCmdTagName, "tag-name", "", "GroupについたTag名。完全一致するTag名が検索対象となる。tag_nameを指定した場合、tag_valueはが必須。")

  GroupsListCmd.Flags().StringVar(&GroupsListCmdTagValue, "tag-value", "", "GroupについたTagの値。")

  GroupsListCmd.Flags().StringVar(&GroupsListCmdTagValueMatchMode, "tag-value-match-mode", "", "Tagの値のマッチングモードを指定。完全一致 (exact) あるいは前方一致 (prefix)。無指定の場合のデフォルトはexact。")

  GroupsListCmd.Flags().Int64Var(&GroupsListCmdLimit, "limit", 0, "レスポンス1ページあたりの最大数")




  GroupsCmd.AddCommand(GroupsListCmd)
}

var GroupsListCmd = &cobra.Command{
  Use: "list",
  Short: TR("List Groups"),
  Long: TR(`Groupの一覧を返す
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
    
    param, err := collectGroupsListCmdParams()
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

func collectGroupsListCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "GET",
    path: buildPathForGroupsListCmd("/groups"),
    query: buildQueryForGroupsListCmd(),
    
    
  }, nil
}

func buildPathForGroupsListCmd(path string) string {
  
  
  
  
  
  
  
  
  
  
  
  
  
  return path
}

func buildQueryForGroupsListCmd() string {
  result := []string{}
  
  
  if GroupsListCmdLastEvaluatedKey != "" {
    result = append(result, fmt.Sprintf("%s=%s", "last_evaluated_key", GroupsListCmdLastEvaluatedKey))
  }
  
  
  
  if GroupsListCmdTagName != "" {
    result = append(result, fmt.Sprintf("%s=%s", "tag_name", GroupsListCmdTagName))
  }
  
  
  
  if GroupsListCmdTagValue != "" {
    result = append(result, fmt.Sprintf("%s=%s", "tag_value", GroupsListCmdTagValue))
  }
  
  
  
  if GroupsListCmdTagValueMatchMode != "" {
    result = append(result, fmt.Sprintf("%s=%s", "tag_value_match_mode", GroupsListCmdTagValueMatchMode))
  }
  
  

  
  
  if GroupsListCmdLimit != 0 {
    result = append(result, fmt.Sprintf("%s=%d", "limit", GroupsListCmdLimit))
  }
  
  

  
  return strings.Join(result, "&")
}


