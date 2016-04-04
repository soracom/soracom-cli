package cmd

import (

  "fmt"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var SubscribersListCmdLastEvaluatedKey string

var SubscribersListCmdSpeedClassFilter string

var SubscribersListCmdStatusFilter string

var SubscribersListCmdTagName string

var SubscribersListCmdTagValue string

var SubscribersListCmdTagValueMatchMode string


var SubscribersListCmdLimit int64




func init() {
  SubscribersListCmd.Flags().StringVar(&SubscribersListCmdLastEvaluatedKey, "last-evaluated-key", "", "現ページで取得した最後のSubscriberのIMSI。このパラメータを指定することで次のSubscriber以降のリストを取得できる。")

  SubscribersListCmd.Flags().StringVar(&SubscribersListCmdSpeedClassFilter, "speed-class-filter", "", "検索対象にする速度クラス。`|`で区切って複数指定することができる。指定可能な値の一覧は以下のとおり: `s1.minimum`, `s1.slow`, `s1.standard`, `s1.fast`")

  SubscribersListCmd.Flags().StringVar(&SubscribersListCmdStatusFilter, "status-filter", "", "検索対象にするstatus。`|`で区切って複数指定することができる。指定可能な値の一覧は以下のとおり: `active`, `inactive`, `ready`, `instock`, `shipped`, `suspended`, `terminated`")

  SubscribersListCmd.Flags().StringVar(&SubscribersListCmdTagName, "tag-name", "", "検索対象にするタグの名前(完全一致)。")

  SubscribersListCmd.Flags().StringVar(&SubscribersListCmdTagValue, "tag-value", "", "検索対象にするタグの検索文字列。`tag_name` を指定した場合は必須。")

  SubscribersListCmd.Flags().StringVar(&SubscribersListCmdTagValueMatchMode, "tag-value-match-mode", "", "タグの検索条件。")

  SubscribersListCmd.Flags().Int64Var(&SubscribersListCmdLimit, "limit", 0, "取得するSubscriberの上限")




  SubscribersCmd.AddCommand(SubscribersListCmd)
}

var SubscribersListCmd = &cobra.Command{
  Use: "list",
  Short: TR("List Subscribers"),
  Long: TR(`条件にマッチするSubscriberのリストを返す。Subscriberの総数が1ページに収まらない場合は、レスポンス中に次のページにアクセスするためのURLを'Link'ヘッダに含めて返す。
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
    
    param, err := collectSubscribersListCmdParams()
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

func collectSubscribersListCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "GET",
    path: buildPathForSubscribersListCmd("/subscribers"),
    query: buildQueryForSubscribersListCmd(),
    
    
  }, nil
}

func buildPathForSubscribersListCmd(path string) string {
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  return path
}

func buildQueryForSubscribersListCmd() string {
  result := []string{}
  
  
  if SubscribersListCmdLastEvaluatedKey != "" {
    result = append(result, fmt.Sprintf("%s=%s", "last_evaluated_key", SubscribersListCmdLastEvaluatedKey))
  }
  
  
  
  if SubscribersListCmdSpeedClassFilter != "" {
    result = append(result, fmt.Sprintf("%s=%s", "speed_class_filter", SubscribersListCmdSpeedClassFilter))
  }
  
  
  
  if SubscribersListCmdStatusFilter != "" {
    result = append(result, fmt.Sprintf("%s=%s", "status_filter", SubscribersListCmdStatusFilter))
  }
  
  
  
  if SubscribersListCmdTagName != "" {
    result = append(result, fmt.Sprintf("%s=%s", "tag_name", SubscribersListCmdTagName))
  }
  
  
  
  if SubscribersListCmdTagValue != "" {
    result = append(result, fmt.Sprintf("%s=%s", "tag_value", SubscribersListCmdTagValue))
  }
  
  
  
  if SubscribersListCmdTagValueMatchMode != "" {
    result = append(result, fmt.Sprintf("%s=%s", "tag_value_match_mode", SubscribersListCmdTagValueMatchMode))
  }
  
  

  
  
  if SubscribersListCmdLimit != 0 {
    result = append(result, fmt.Sprintf("%s=%d", "limit", SubscribersListCmdLimit))
  }
  
  

  
  return strings.Join(result, "&")
}


