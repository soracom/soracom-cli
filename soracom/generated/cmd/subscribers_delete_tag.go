package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var SubscribersDeleteTagCmdImsi string

var SubscribersDeleteTagCmdTagName string






func init() {
  SubscribersDeleteTagCmd.Flags().StringVar(&SubscribersDeleteTagCmdImsi, "imsi", "", "対象のSubscriberのIMSI")

  SubscribersDeleteTagCmd.Flags().StringVar(&SubscribersDeleteTagCmdTagName, "tag-name", "", "削除対象のタグ名（URL の Path の一部になるので、パーセントエンコーディングを施す。JavaScript なら encodeURIComponent() したものを指定する）")




  SubscribersCmd.AddCommand(SubscribersDeleteTagCmd)
}

var SubscribersDeleteTagCmd = &cobra.Command{
  Use: "delete-tag",
  Short: TR("Delete Subscriber Tag"),
  Long: TR(`指定されたSubscriberのタグを削除`),
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
    
    param, err := collectSubscribersDeleteTagCmdParams()
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

func collectSubscribersDeleteTagCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "DELETE",
    path: buildPathForSubscribersDeleteTagCmd("/subscribers/{imsi}/tags/{tag_name}"),
    query: buildQueryForSubscribersDeleteTagCmd(),
    
    
  }, nil
}

func buildPathForSubscribersDeleteTagCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "imsi" + "}", SubscribersDeleteTagCmdImsi, -1)
  
  
  
  path = strings.Replace(path, "{" + "tag_name" + "}", SubscribersDeleteTagCmdTagName, -1)
  
  
  
  
  
  return path
}

func buildQueryForSubscribersDeleteTagCmd() string {
  result := []string{}
  
  
  
  
  

  

  

  

  return strings.Join(result, "&")
}


