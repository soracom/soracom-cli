package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var SubscribersDisableTerminationCmdImsi string






func init() {
  SubscribersDisableTerminationCmd.Flags().StringVar(&SubscribersDisableTerminationCmdImsi, "imsi", "", "対象のSubscriberのIMSI")




  SubscribersCmd.AddCommand(SubscribersDisableTerminationCmd)
}

var SubscribersDisableTerminationCmd = &cobra.Command{
  Use: "disable-termination",
  Short: TR("Disable Termination of Subscriber"),
  Long: TR(`指定されたSubscriberをTerminate不可能に設定する`),
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
    
    param, err := collectSubscribersDisableTerminationCmdParams()
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

func collectSubscribersDisableTerminationCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "POST",
    path: buildPathForSubscribersDisableTerminationCmd("/subscribers/{imsi}/disable_termination"),
    query: buildQueryForSubscribersDisableTerminationCmd(),
    
    
  }, nil
}

func buildPathForSubscribersDisableTerminationCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "imsi" + "}", SubscribersDisableTerminationCmdImsi, -1)
  
  
  
  
  
  return path
}

func buildQueryForSubscribersDisableTerminationCmd() string {
  result := []string{}
  
  
  

  

  

  

  return strings.Join(result, "&")
}


