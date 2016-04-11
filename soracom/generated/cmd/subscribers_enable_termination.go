package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var SubscribersEnableTerminationCmdImsi string






func init() {
  SubscribersEnableTerminationCmd.Flags().StringVar(&SubscribersEnableTerminationCmdImsi, "imsi", "", "対象のSubscriberのIMSI")




  SubscribersCmd.AddCommand(SubscribersEnableTerminationCmd)
}

var SubscribersEnableTerminationCmd = &cobra.Command{
  Use: "enable-termination",
  Short: TR("Enable Termination of Subscriber"),
  Long: TR(`指定されたSubscriberをTerminate可能に設定する`),
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
    
    param, err := collectSubscribersEnableTerminationCmdParams()
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

func collectSubscribersEnableTerminationCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "POST",
    path: buildPathForSubscribersEnableTerminationCmd("/subscribers/{imsi}/enable_termination"),
    query: buildQueryForSubscribersEnableTerminationCmd(),
    
    
  }, nil
}

func buildPathForSubscribersEnableTerminationCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "imsi" + "}", SubscribersEnableTerminationCmdImsi, -1)
  
  
  
  
  
  return path
}

func buildQueryForSubscribersEnableTerminationCmd() string {
  result := []string{}
  
  
  

  

  

  

  return strings.Join(result, "&")
}


