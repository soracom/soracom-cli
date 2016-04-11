package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var SubscribersDeactivateCmdImsi string






func init() {
  SubscribersDeactivateCmd.Flags().StringVar(&SubscribersDeactivateCmdImsi, "imsi", "", "対象のSubscriberのIMSI")




  SubscribersCmd.AddCommand(SubscribersDeactivateCmd)
}

var SubscribersDeactivateCmd = &cobra.Command{
  Use: "deactivate",
  Short: TR("Deactivate Subscriber"),
  Long: TR(`指定されたSubscriberを無効化`),
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
    
    param, err := collectSubscribersDeactivateCmdParams()
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

func collectSubscribersDeactivateCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "POST",
    path: buildPathForSubscribersDeactivateCmd("/subscribers/{imsi}/deactivate"),
    query: buildQueryForSubscribersDeactivateCmd(),
    
    
  }, nil
}

func buildPathForSubscribersDeactivateCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "imsi" + "}", SubscribersDeactivateCmdImsi, -1)
  
  
  
  
  
  return path
}

func buildQueryForSubscribersDeactivateCmd() string {
  result := []string{}
  
  
  

  

  

  

  return strings.Join(result, "&")
}


