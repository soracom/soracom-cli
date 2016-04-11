package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var SubscribersSessionEventsCmdImsi string

var SubscribersSessionEventsCmdLastEvaluatedKey string


var SubscribersSessionEventsCmdFrom int64

var SubscribersSessionEventsCmdLimit int64

var SubscribersSessionEventsCmdTo int64





func init() {
  SubscribersSessionEventsCmd.Flags().StringVar(&SubscribersSessionEventsCmdImsi, "imsi", "", "対象のSubscriberのIMSI")

  SubscribersSessionEventsCmd.Flags().StringVar(&SubscribersSessionEventsCmdLastEvaluatedKey, "last-evaluated-key", "", "現ページで取得した最後のイベントのタイムスタンプ。このパラメータを指定することで次のイベント以降のリストを取得できる。")

  SubscribersSessionEventsCmd.Flags().Int64Var(&SubscribersSessionEventsCmdFrom, "from", 0, "イベントの検索範囲時刻の始まり")

  SubscribersSessionEventsCmd.Flags().Int64Var(&SubscribersSessionEventsCmdLimit, "limit", 0, "取得するイベント数の上限")

  SubscribersSessionEventsCmd.Flags().Int64Var(&SubscribersSessionEventsCmdTo, "to", 0, "イベントの検索範囲時刻の終わり")




  SubscribersCmd.AddCommand(SubscribersSessionEventsCmd)
}

var SubscribersSessionEventsCmd = &cobra.Command{
  Use: "session-events",
  Short: TR("List Session Events"),
  Long: TR(`指定されたSubscriberのセッション作成・変更・削除のイベント履歴を返す。イベントの総数が1ページに収まらない場合は、レスポンス中に次のページにアクセスするためのURLを'Link'ヘッダに含めて返す。
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
    
    param, err := collectSubscribersSessionEventsCmdParams()
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

func collectSubscribersSessionEventsCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "GET",
    path: buildPathForSubscribersSessionEventsCmd("/subscribers/{imsi}/events/sessions"),
    query: buildQueryForSubscribersSessionEventsCmd(),
    
    
  }, nil
}

func buildPathForSubscribersSessionEventsCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "imsi" + "}", SubscribersSessionEventsCmdImsi, -1)
  
  
  
  
  
  
  
  
  
  
  
  
  
  return path
}

func buildQueryForSubscribersSessionEventsCmd() string {
  result := []string{}
  
  
  
  
  if SubscribersSessionEventsCmdLastEvaluatedKey != "" {
    result = append(result, sprintf("%s=%s", "last_evaluated_key", SubscribersSessionEventsCmdLastEvaluatedKey))
  }
  
  

  
  
  if SubscribersSessionEventsCmdFrom != 0 {
    result = append(result, sprintf("%s=%d", "from", SubscribersSessionEventsCmdFrom))
  }
  
  
  
  if SubscribersSessionEventsCmdLimit != 0 {
    result = append(result, sprintf("%s=%d", "limit", SubscribersSessionEventsCmdLimit))
  }
  
  
  
  if SubscribersSessionEventsCmdTo != 0 {
    result = append(result, sprintf("%s=%d", "to", SubscribersSessionEventsCmdTo))
  }
  
  

  

  

  return strings.Join(result, "&")
}


