package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var SubscribersTerminateCmdImsi string





func init() {
  SubscribersTerminateCmd.Flags().StringVar(&SubscribersTerminateCmdImsi, "imsi", "", "対象のSubscriberのIMSI")




  SubscribersCmd.AddCommand(SubscribersTerminateCmd)
}

var SubscribersTerminateCmd = &cobra.Command{
  Use: "terminate",
  Short: TR("Terminate Subscriber"),
  Long: TR(`指定されたSubscriberをTerminate`),
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
    
    param, err := collectSubscribersTerminateCmdParams()
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

func collectSubscribersTerminateCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "POST",
    path: buildPathForSubscribersTerminateCmd("/subscribers/{imsi}/terminate"),
    query: buildQueryForSubscribersTerminateCmd(),
    
    
  }, nil
}

func buildPathForSubscribersTerminateCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "imsi" + "}", SubscribersTerminateCmdImsi, -1)
  
  
  
  
  return path
}

func buildQueryForSubscribersTerminateCmd() string {
  result := []string{}
  
  
  

  

  
  return strings.Join(result, "&")
}


