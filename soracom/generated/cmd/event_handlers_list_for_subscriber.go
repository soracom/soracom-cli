package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var EventHandlersListForSubscriberCmdImsi string





func init() {
  EventHandlersListForSubscriberCmd.Flags().StringVar(&EventHandlersListForSubscriberCmdImsi, "imsi", "", "imsi")




  EventHandlersCmd.AddCommand(EventHandlersListForSubscriberCmd)
}

var EventHandlersListForSubscriberCmd = &cobra.Command{
  Use: "list-for-subscriber",
  Short: TR("List Event Handlers related to Subscriber"),
  Long: TR(`対象IMSIにひもづくイベントハンドラのリストを返す。
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
    
    param, err := collectEventHandlersListForSubscriberCmdParams()
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

func collectEventHandlersListForSubscriberCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "GET",
    path: buildPathForEventHandlersListForSubscriberCmd("/event_handlers/subscribers/{imsi}"),
    query: buildQueryForEventHandlersListForSubscriberCmd(),
    
    
  }, nil
}

func buildPathForEventHandlersListForSubscriberCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "imsi" + "}", EventHandlersListForSubscriberCmdImsi, -1)
  
  
  
  
  return path
}

func buildQueryForEventHandlersListForSubscriberCmd() string {
  result := []string{}
  
  
  

  

  
  return strings.Join(result, "&")
}


