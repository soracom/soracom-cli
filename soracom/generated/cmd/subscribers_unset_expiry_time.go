package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var SubscribersUnsetExpiryTimeCmdImsi string






func init() {
  SubscribersUnsetExpiryTimeCmd.Flags().StringVar(&SubscribersUnsetExpiryTimeCmdImsi, "imsi", "", "対象のSubscriberのIMSI")




  SubscribersCmd.AddCommand(SubscribersUnsetExpiryTimeCmd)
}

var SubscribersUnsetExpiryTimeCmd = &cobra.Command{
  Use: "unset-expiry-time",
  Short: TR("Delete Expiry Time of Subscriber"),
  Long: TR(`指定されたSubscriberの有効期限を削除して無期限に変更`),
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
    
    param, err := collectSubscribersUnsetExpiryTimeCmdParams()
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

func collectSubscribersUnsetExpiryTimeCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "POST",
    path: buildPathForSubscribersUnsetExpiryTimeCmd("/subscribers/{imsi}/unset_expiry_time"),
    query: buildQueryForSubscribersUnsetExpiryTimeCmd(),
    
    
  }, nil
}

func buildPathForSubscribersUnsetExpiryTimeCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "imsi" + "}", SubscribersUnsetExpiryTimeCmdImsi, -1)
  
  
  
  
  
  return path
}

func buildQueryForSubscribersUnsetExpiryTimeCmd() string {
  result := []string{}
  
  
  

  

  

  

  return strings.Join(result, "&")
}


