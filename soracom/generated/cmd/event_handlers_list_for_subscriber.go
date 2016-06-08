package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var EventHandlersListForSubscriberCmdImsi string






func init() {
  EventHandlersListForSubscriberCmd.Flags().StringVar(&EventHandlersListForSubscriberCmdImsi, "imsi", "", TR("event_handlers.list_event_handlers_by_subscriber.get.parameters.imsi.description"))




  EventHandlersCmd.AddCommand(EventHandlersListForSubscriberCmd)
}

var EventHandlersListForSubscriberCmd = &cobra.Command{
  Use: "list-for-subscriber",
  Short: TR("event_handlers.list_event_handlers_by_subscriber.get.summary"),
  Long: TR(`event_handlers.list_event_handlers_by_subscriber.get.description`),
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

    if result != "" {
      return prettyPrintStringAsJSON(result)
    } else {
      return nil
    }
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


